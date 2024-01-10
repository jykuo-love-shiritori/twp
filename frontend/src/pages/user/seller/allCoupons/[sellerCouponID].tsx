import { Row, Col } from 'react-bootstrap';
import { useEffect, useState } from 'react';
import { useForm, SubmitHandler, useFieldArray } from 'react-hook-form';
import { useQuery } from '@tanstack/react-query';
import { useNavigate, useParams } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTrash } from '@fortawesome/free-solid-svg-icons';
import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import { formatDate } from '@lib/Functions';
import { GetUserName, useAuth } from '@lib/Auth';
import TButton from '@components/TButton';
import FormItem from '@components/FormItem';
import CouponItemTemplate from '@components/CouponItemTemplate';

interface IShopCouponDetail {
  coupon_info: ICouponInfo;
  tags: [ITag];
}

interface ICouponInfo {
  description: string;
  discount: number;
  expire_date: string;
  name: string;
  scope: 'global' | 'shop';
  start_date: string;
  type: 'percentage' | 'fixed' | 'shipping';
}

interface ITag {
  name: string;
  tag_id: number;
}

const tagStyle = {
  borderRadius: '30px',
  background: ' var(--button_light)',
  padding: '1% 1% 1% 3%',
  color: 'white',
  margin: '5px 0 5px 5px',
  width: '100%',
};

const EachSellerCoupon = () => {
  const [tag, setTag] = useState('');
  const [suggestTags, setSuggestTags] = useState<ITag[]>([]);
  const navigate = useNavigate();
  const { coupon_id } = useParams();
  const token = useAuth();
  const username = GetUserName();

  const { register, control, handleSubmit, watch, reset } = useForm<IShopCouponDetail>({
    defaultValues: {
      coupon_info: {
        description: '',
        discount: 0,
        expire_date: '',
        name: '',
        scope: 'shop',
        start_date: '',
        type: 'percentage',
      },
      tags: [],
    },
  });
  const { fields, append, remove } = useFieldArray({
    control,
    name: 'tags',
  });

  const watchAllFields = watch();

  const { data: initData, status: initStatus } = useQuery({
    queryKey: ['sellerGetCouponDetail'],
    queryFn: async () => {
      const resp = await fetch(`/api/seller/coupon/${coupon_id}`, {
        method: 'GET',
        headers: {
          Authorization: `Bearer ${token}`,
          accept: 'application/json',
        },
      });
      if (!resp.ok) {
        RouteOnNotOK(resp, navigate);
      } else {
        return await resp.json();
      }
    },
    select: (data) => data as IShopCouponDetail,
    enabled: true,
    refetchOnWindowFocus: false,
  });

  const addNewTag = async (event: React.KeyboardEvent<HTMLInputElement>) => {
    // this addressed the magic number: https://github.com/facebook/react/issues/14512
    if (event.keyCode === 229) return;
    if (event.key === 'Enter') {
      const newTagName = event.currentTarget.value.trim();
      if (newTagName !== '') {
        let newTag: ITag;
        const foundOldTag = suggestTags.find((tag) => tag.name === newTagName);
        if (foundOldTag) {
          // using old tag
          if (fields.find((tag) => tag.name === newTagName)) {
            alert('tag already exist');
            setTag('');
            return;
          }
          newTag = foundOldTag;
        } else {
          // creating a new tag
          const resp = await fetch('/api/seller/tag', {
            method: 'POST',
            headers: {
              Authorization: `Bearer ${token}`,
              accept: 'application/json',
              'Content-Type': 'application/json',
            },
            body: JSON.stringify({ name: newTagName, seller_name: username }),
          });
          if (!resp.ok) {
            if (resp.status === 500) {
              alert("error on adding tag, please check your shop's status");
            } else {
              alert('error when creating new tag');
            }
            navigate('/user/seller/manageCoupons');
            return;
          } else {
            const response = await resp.json();
            newTag = { name: newTagName, tag_id: response.id };
          }
        }

        const resp = await fetch(`/api/seller/coupon/${coupon_id}/tag`, {
          method: 'POST',
          headers: {
            Authorization: `Bearer ${token}`,
            accept: 'application/json',
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({ tag_id: newTag.tag_id }),
        });
        if (!resp.ok) {
          if (resp.status === 500) {
            alert("error on adding tag, please check your shop's status");
          } else {
            alert('error when adding new tag');
          }
          navigate('/user/seller/manageCoupons');
          return;
        } else {
          append(newTag);
          setTag('');
        }
      }
    }
  };
  const onChangeNewTag = async (event: React.ChangeEvent<HTMLInputElement>) => {
    setTag(event.target.value);
    if (event.target.value === '') {
      setSuggestTags([]);
      return;
    }
    const resp = await fetch(`/api/seller/tag?name=${event.target.value}`, {
      method: 'GET',
      headers: {
        Authorization: `Bearer ${token}`,
        accept: 'application/json',
      },
    });
    if (!resp.ok) {
      alert('error when fetching existing tag');
    } else {
      const response = await resp.json();
      const newSuggestTags: ITag[] = [];
      response.forEach((tag: { name: string; id: number }) => {
        newSuggestTags.push({ name: tag.name, tag_id: tag.id });
      });
      setSuggestTags(newSuggestTags);
    }
  };
  const deleteTag = async (index: number) => {
    const resp = await fetch(`/api/seller/coupon/${coupon_id}/tag`, {
      method: 'DELETE',
      headers: {
        Authorization: `Bearer ${token}`,
        accept: 'application/json',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ tag_id: fields[index].tag_id }),
    });
    console.log(resp);
    if (!resp.ok) {
      if (resp.status === 404) {
        alert("error on deleting tag, please check your shop's status");
      } else {
        alert('error when deleting tag');
      }
      navigate('/user/seller/manageCoupons');
      return;
    } else {
      remove(index);
    }
  };

  const OnConfirm: SubmitHandler<IShopCouponDetail> = async (data) => {
    const startDate = new Date(data.coupon_info.start_date);
    const expDate = new Date(data.coupon_info.expire_date);
    const today = new Date();
    startDate.setHours(0, 0, 0, 0);
    expDate.setHours(0, 0, 0, 0);
    today.setHours(0, 0, 0, 0);
    // TODO: change to form validation
    if (startDate < today) {
      alert('Start date cannot be earlier than today');
      return;
    }
    if (startDate >= expDate) {
      alert('Start date should be earlier than expire date');
      return;
    }
    if (data.coupon_info.type === 'percentage' && data.coupon_info.discount >= 100) {
      alert('Discount should be less than 100%');
      return;
    }
    if (data.coupon_info.type === 'shipping') {
      data.coupon_info.discount = 0;
    }
    if (data.coupon_info.type != 'shipping' && data.coupon_info.discount <= 0) {
      alert('Discount should be greater than 0');
      return;
    }

    interface INewCoupon {
      description: string;
      discount: number;
      expire_date: string;
      name: string;
      start_date: string;
      tags: number[];
      type: 'percentage' | 'fixed' | 'shipping';
    }
    // I have no idea how discount get turned into string
    const newCoupon: INewCoupon = {
      description: data.coupon_info.description,
      discount: Number(data.coupon_info.discount),
      expire_date: new Date(data.coupon_info.expire_date).toISOString(),
      name: data.coupon_info.name,
      start_date: new Date(data.coupon_info.start_date).toISOString(),
      tags: data.tags.map((tag) => tag.tag_id),
      type: data.coupon_info.type,
    };
    const resp = await fetch(`/api/seller/coupon/${coupon_id}`, {
      method: 'PATCH',
      headers: {
        Authorization: `Bearer ${token}`,
        accept: 'application/json',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(newCoupon),
    });
    // 500 means shop is not enabled
    if (!resp.ok) {
      if (resp.status === 500) {
        alert("error on modifying coupon, please check your shop's status");
        navigate('/user/seller/manageCoupons');
      } else {
        const response = await resp.json();
        alert(response.message);
      }
    } else {
      navigate('/user/seller/manageCoupons');
    }
  };

  const OnDelete = async () => {
    const resp = await fetch(`/api/seller/coupon/${coupon_id}`, {
      method: 'DELETE',
      headers: {
        Authorization: `Bearer ${token}`,
        accept: 'application/json',
      },
    });
    // 404 means shop is not enabled
    if (resp.status === 404) {
      alert("error on deleting coupon, please check your shop's status");
    } else if (!resp.ok) {
      alert('error when deleting coupon');
    }
    navigate('/user/seller/manageCoupons');
    return;
  };

  useEffect(() => {
    if (initStatus === 'success') {
      const startDate = new Date(initData.coupon_info.start_date);
      const expDate = new Date(initData.coupon_info.expire_date);
      reset({
        coupon_info: {
          ...initData.coupon_info,
          start_date: formatDate(startDate.toISOString()),
          expire_date: formatDate(expDate.toISOString()),
        },
        tags: initData.tags,
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
                    name: watchAllFields.coupon_info.name,
                    type: watchAllFields.coupon_info.type,
                    discount: watchAllFields.coupon_info.discount,
                    expire_date: watchAllFields.coupon_info.expire_date,
                  }}
                />
              </div>
              <span className='dark'>add more tags</span>

              {/* new tag input */}
              <input
                type='text'
                placeholder='Input tags'
                className='quantity_box'
                value={tag}
                onChange={onChangeNewTag}
                onKeyDown={addNewTag}
                style={{ marginBottom: '10px' }}
                list='suggestion'
              />
              <datalist id='suggestion'>
                {suggestTags.map((tag, index) => (
                  <option key={index} value={tag.name} />
                ))}
              </datalist>
              {/* dynamic tags fields */}
              {fields.map((field, index) => (
                <div key={field.id} style={tagStyle}>
                  <Row style={{ width: '100%' }} className='center'>
                    <Col xs={2} className='right'>
                      <FontAwesomeIcon
                        icon={faTrash}
                        className='white_word pointer'
                        onClick={() => deleteTag(index)}
                      />
                    </Col>
                    <Col>{field.name}</Col>
                  </Row>
                </div>
              ))}

              {/* delete, confirm button */}
              <div style={{ height: '50px' }} />
              <TButton text='Cancel' action='/user/seller/manageCoupons' />
              <TButton text='Delete Coupon' action={OnDelete} />
              <TButton text='Confirm Changes' action={handleSubmit(OnConfirm)} />
            </div>
          </Col>

          {/* right half */}
          <Col xs={12} md={7}>
            <div style={{ padding: '7% 0% 7% 2%' }}>
              <FormItem label='Coupon Name'>
                <input type='text' {...register('coupon_info.name', { required: true })} />
              </FormItem>
              <FormItem label='Coupon description'>
                <textarea {...register('coupon_info.description', { required: true })} />
              </FormItem>
              <FormItem label='Method'>
                <select {...register('coupon_info.type', { required: true })}>
                  <option value='percentage'>percentage</option>
                  <option value='fixed'>fixed</option>
                  <option value='shipping'>shipping</option>
                </select>
              </FormItem>
              {watch('coupon_info.type') != 'shipping' ? (
                <FormItem label='Discount'>
                  <input type='number' {...register('coupon_info.discount')} />
                </FormItem>
              ) : (
                <></>
              )}
              <FormItem label='Start Date'>
                <input type='date' {...register('coupon_info.start_date', { required: true })} />
              </FormItem>
              <FormItem label='Expire Date'>
                <input type='date' {...register('coupon_info.expire_date', { required: true })} />
              </FormItem>
            </div>
          </Col>
        </Row>
      </form>
    </div>
  );
};

export default EachSellerCoupon;
