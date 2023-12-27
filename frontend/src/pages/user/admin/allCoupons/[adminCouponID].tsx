import { Row, Col } from 'react-bootstrap';
import { useEffect } from 'react';
import { useForm, SubmitHandler } from 'react-hook-form';
import { useQuery } from '@tanstack/react-query';
import { useNavigate, useParams } from 'react-router-dom';
import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import { formatDate } from '@lib/Functions';
import TButton from '@components/TButton';
import FormItem from '@components/FormItem';
import CouponItemTemplate from '@components/CouponItemTemplate';

interface IGlobalCouponDetail {
  description: string;
  discount: number;
  expire_date: string;
  id: number;
  name: string;
  scope: 'global' | 'shop';
  start_date: string;
  type: 'percentage' | 'fixed' | 'shipping';
}

const EachAdminCoupon = () => {
  const navigate = useNavigate();
  const { coupon_id } = useParams();

  const { register, handleSubmit, watch, reset } = useForm<IGlobalCouponDetail>({
    defaultValues: {
      description: '',
      discount: 0,
      expire_date: '',
      id: 0,
      name: '',
      scope: 'global',
      start_date: '',
      type: 'percentage',
    },
  });

  const watchAllFields = watch();

  const { data: initData, status: initStatus } = useQuery({
    queryKey: ['adminGetCouponDetail', coupon_id],
    queryFn: async () => {
      const resp = await fetch(`/api/admin/coupon/${coupon_id}`, {
        method: 'GET',
        headers: {
          accept: 'application/json',
        },
      });
      if (!resp.ok) {
        RouteOnNotOK(resp, navigate);
      } else {
        return await resp.json();
      }
    },
    select: (data) => data as IGlobalCouponDetail,
    enabled: true,
    refetchOnWindowFocus: false,
  });

  const OnConfirm: SubmitHandler<IGlobalCouponDetail> = async (data) => {
    const startDate = new Date(data.start_date);
    const expDate = new Date(data.expire_date);
    const today = new Date();
    if (startDate < today) {
      alert('Start date should be later than today');
      return;
    }
    startDate.setHours(0, 0, 0, 0);
    expDate.setHours(0, 0, 0, 0);
    if (startDate >= expDate) {
      alert('Start date should be earlier than expire date');
      return;
    }
    if (data.type === 'percentage' && data.discount >= 100) {
      alert('Discount should be less than 100%');
      return;
    }
    // nice "end_date" backend team 🥳
    interface INewCoupon {
      description: string;
      discount: number;
      end_date: string;
      name: string;
      start_date: string;
      type: 'percentage' | 'fixed' | 'shipping';
    }
    // I have no idea how discount get turned into string
    const newCoupon: INewCoupon = {
      description: data.description,
      discount: Number(data.discount),
      end_date: new Date(data.expire_date).toISOString(),
      name: data.name,
      start_date: new Date(data.start_date).toISOString(),
      type: data.type,
    };
    const resp = await fetch(`/api/admin/coupon/${coupon_id}`, {
      method: 'PATCH',
      headers: {
        accept: 'application/json',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(newCoupon),
    });
    if (!resp.ok) {
      console.log('error when updating coupon');
    } else {
      navigate('/admin/manageCoupons');
    }
  };
  const OnDelete = async () => {
    const resp = await fetch(`/api/admin/coupon/${coupon_id}`, {
      method: 'DELETE',
      headers: {
        accept: 'application/json',
      },
    });
    if (!resp.ok) {
      console.log('error when deleting coupon');
    } else {
      navigate('/admin/manageCoupons');
    }
  };

  useEffect(() => {
    if (initStatus === 'success') {
      const startDate = new Date(initData.start_date);
      const expDate = new Date(initData.expire_date);
      reset({
        ...initData,
        start_date: formatDate(startDate.toISOString()),
        expire_date: formatDate(expDate.toISOString()),
      });
    }
  }, [initData, initStatus, reset]);

  if (initStatus !== 'success') {
    return <CheckFetchStatus status={initStatus} />;
  }

  return (
    <div style={{ padding: '55px 12% 0 12%' }}>
      <form onSubmit={handleSubmit(OnConfirm)}>
        <Row>
          {/* left half */}
          <Col xs={12} md={5} className='goods_bgW'>
            <div className='flex_wrapper' style={{ padding: '0 8% 10% 8%' }}>
              {/* sample display */}
              <div style={{ padding: '15% 10%' }}>
                <CouponItemTemplate
                  data={{
                    name: watchAllFields.name,
                    type: watchAllFields.type,
                    discount: watchAllFields.discount,
                    expire_date: watchAllFields.expire_date,
                  }}
                />
              </div>

              {/* delete, confirm button */}
              <div style={{ height: '50px' }} />
              <TButton text='Delete Coupon' action={OnDelete} />
              <TButton text='Confirm Changes' action={handleSubmit(OnConfirm)} />
            </div>
          </Col>

          {/* right half */}
          <Col xs={12} md={7}>
            <div style={{ padding: '7% 0% 7% 2%' }}>
              <FormItem label='Coupon Name'>
                <input type='text' {...register('name', { required: true })} />
              </FormItem>

              <FormItem label='Coupon description'>
                <textarea {...register('description', { required: true })} />
              </FormItem>

              <FormItem label='Method'>
                <select {...register('type', { required: true })}>
                  <option value='percentage'>percentage</option>
                  <option value='fixed'>fixed</option>
                  <option value='shipping'>shipping</option>
                </select>
              </FormItem>

              <FormItem label='Discount'>
                <input type='number' {...register('discount', { required: true })} />
              </FormItem>

              <FormItem label='Start Date'>
                <input type='date' {...register('start_date', { required: true })} />
              </FormItem>

              <FormItem label='Expire Date'>
                <input type='date' {...register('expire_date', { required: true })} />
              </FormItem>
            </div>
          </Col>
        </Row>
      </form>
    </div>
  );
};

export default EachAdminCoupon;
