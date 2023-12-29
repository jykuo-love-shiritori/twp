import { Row, Col } from 'react-bootstrap';
import { useState } from 'react';
import { useForm, SubmitHandler, useFieldArray } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTrash } from '@fortawesome/free-solid-svg-icons';
import { formatDate } from '@lib/Functions';
import TButton from '@components/TButton';
import FormItem from '@components/FormItem';
import CouponItemTemplate from '@components/CouponItemTemplate';

interface IShopCouponDetail {
  coupon_info: {
    description: string;
    discount: number;
    expire_date: string;
    name: string;
    scope: 'global' | 'shop';
    start_date: string;
    type: 'percentage' | 'fixed' | 'shipping';
  };
  tags: [ITag];
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

const NewSellerCoupon = () => {
  const [tag, setTag] = useState('');
  const [suggestTags, setSuggestTags] = useState<ITag[]>([]);
  const navigate = useNavigate();

  const { register, control, handleSubmit, watch } = useForm<IShopCouponDetail>({
    defaultValues: {
      coupon_info: {
        description: '',
        discount: 0,
        expire_date: formatDate(new Date().toISOString()),
        name: '',
        scope: 'shop',
        start_date: formatDate(new Date().toISOString()),
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

  const addNewTag = async (event: React.KeyboardEvent<HTMLInputElement>) => {
    // this addressed the magic number: https://github.com/facebook/react/issues/14512
    if (event.keyCode === 229) return;
    if (event.key === 'Enter') {
      const newTagName = event.currentTarget.value.trim();
      if (newTagName !== '') {
        const foundOldTag = suggestTags.find((tag) => tag.name === newTagName);
        if (foundOldTag) {
          // using old tag
          if (fields.find((tag) => tag.name === newTagName)) {
            alert('tag already exist');
            setTag('');
            return;
          }
          append(foundOldTag);
          setTag('');
        } else {
          // creating a new tag
          const resp = await fetch('/api/seller/tag', {
            method: 'POST',
            headers: {
              accept: 'application/json',
              'Content-Type': 'application/json',
            },
            // TODO: change seller_name to real user name
            body: JSON.stringify({ name: newTagName, seller_name: 'user1' }),
          });
          if (!resp.ok) {
            alert('error when creating new tag');
            return;
          } else {
            const response = await resp.json();
            append({ name: newTagName, tag_id: response.id });
            setTag('');
          }
        }
      }
    }
  };
  const onChangeNewTag = async (event: React.ChangeEvent<HTMLInputElement>) => {
    setTag(event.target.value);
    const resp = await fetch(`/api/seller/tag?name=${event.target.value}`, {
      method: 'GET',
      headers: {
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
  const deleteTag = (index: number) => {
    remove(index);
  };

  const OnConfirm: SubmitHandler<IShopCouponDetail> = async (data) => {
    const startDate = new Date(data.coupon_info.start_date);
    const expDate = new Date(data.coupon_info.expire_date);
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
    if (data.coupon_info.type === 'percentage' && data.coupon_info.discount >= 100) {
      alert('Discount should be less than 100%');
      return;
    }
    if (data.coupon_info.discount < 0) {
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
    const resp = await fetch(`/api/seller/coupon`, {
      method: 'POST',
      headers: {
        accept: 'application/json',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(newCoupon),
    });
    if (!resp.ok) {
      const response = await resp.json();
      alert(response.message);
    } else {
      navigate('/user/seller/manageCoupons');
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
              <TButton text='Cancel' action={() => navigate('/user/seller/manageCoupons')} />
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

              <FormItem label='Discount'>
                <input type='number' {...register('coupon_info.discount', { required: true })} />
              </FormItem>

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

export default NewSellerCoupon;
