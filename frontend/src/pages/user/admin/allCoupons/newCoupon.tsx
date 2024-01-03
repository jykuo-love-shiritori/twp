import { Row, Col } from 'react-bootstrap';
import { useForm, SubmitHandler } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';
import { formatDate } from '@lib/Functions';
import { useAuth } from '@lib/Auth';
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

const NewAdminCoupon = () => {
  const navigate = useNavigate();
  const token = useAuth();

  const { register, handleSubmit, watch } = useForm<IGlobalCouponDetail>({
    defaultValues: {
      description: '',
      discount: 0,
      expire_date: formatDate(new Date().toISOString()),
      id: 0,
      name: '',
      scope: 'global',
      start_date: formatDate(new Date().toISOString()),
      type: 'percentage',
    },
  });

  const watchAllFields = watch();

  const OnConfirm: SubmitHandler<IGlobalCouponDetail> = async (data) => {
    const startDate = new Date(data.start_date);
    const expDate = new Date(data.expire_date);
    const today = new Date();
    startDate.setHours(0, 0, 0, 0);
    expDate.setHours(0, 0, 0, 0);
    today.setHours(0, 0, 0, 0);
    if (startDate <= today) {
      alert('Start date should be later than today');
      return;
    }
    if (startDate >= expDate) {
      alert('Start date should be earlier than expire date');
      return;
    }
    if (data.type === 'percentage' && data.discount >= 100) {
      alert('Discount should be less than 100%');
      return;
    }
    if (data.discount < 0) {
      alert('Discount should be greater than 0');
      return;
    }
    interface INewCoupon {
      description: string;
      discount: number;
      expire_date: string;
      name: string;
      scope: 'global' | 'shop';
      start_date: string;
      type: 'percentage' | 'fixed' | 'shipping';
    }
    // I have no idea how discount get turned into string
    const newCoupon: INewCoupon = {
      description: data.description,
      discount: Number(data.discount),
      expire_date: new Date(data.expire_date).toISOString(),
      name: data.name,
      scope: data.scope,
      start_date: new Date(data.start_date).toISOString(),
      type: data.type,
    };
    const resp = await fetch(`/api/admin/coupon`, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token}`,
        accept: 'application/json',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(newCoupon),
    });
    if (!resp.ok) {
      const response = await resp.json();
      alert(response.message);
    } else {
      navigate('/admin/manageCoupons');
    }
  };

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
              <TButton text='Cancel' action={() => navigate('/admin/manageCoupons')} />
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

export default NewAdminCoupon;
