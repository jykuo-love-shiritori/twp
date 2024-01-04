import '@components/style.css';
import '@style/global.css';

import { Col, Form, Row, Offcanvas } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faFilter } from '@fortawesome/free-solid-svg-icons';

import SellerItem, { SellerItemProps } from '@components/SellerItem';

import goodsData from '@pages/discover/goodsData.json';
import GoodsItem from '@components/GoodsItem';
import { SubmitHandler, useForm } from 'react-hook-form';
import TButton from '@components/TButton';
import { useState } from 'react';

const PhoneOffCanvasStyle = {
  width: '330px',
  border: ' 1px solid var(--button_light, #60998B)',
  background: ' var(--bg, #061E18)',
  boxShadow: '0px 4px 20px 2px #60998B',
  padding: '10%',
  color: 'white',
};

interface FilterProps {
  maxPrice: number | null;
  minPrice: number | null;
  maxStock: number | null;
  minStock: number | null;
  hasCoupon: boolean;
  sort: 'price' | 'stock' | 'sales' | 'relevancy' | null;
  order: 'asc' | 'desc' | null;
}

const defaultFilterValues: FilterProps = {
  maxPrice: null,
  minPrice: null,
  maxStock: null,
  minStock: null,
  hasCoupon: false,
  sort: null,
  order: null,
};

const dataArray: SellerItemProps[] = [
  { seller_name: 'test1', name: 'test1', image_url: '' },
  { seller_name: 'test2', name: 'test2', image_url: '' },
  { seller_name: 'test3', name: 'test3', image_url: '' },
  { seller_name: 'test4', name: 'test4', image_url: '' },
];

const isValidNumber = (value: string | number | null | undefined): boolean => {
  return typeof value === 'number' || (typeof value === 'string' && !isNaN(parseInt(value)));
};

const isInteger = (value: string | number | null | undefined): boolean => {
  return typeof value === 'number' || (typeof value === 'string' && /^-?\d+$/.test(value));
};

const isDataValid = (data: FilterProps) => {
  if (typeof data.maxPrice !== 'number' && data.maxPrice !== null) {
    if (!isValidNumber(data.maxPrice)) {
      alert('Please enter numbers for max price');
      return false;
    }
  }

  if (typeof data.minPrice !== 'number' && data.minPrice !== null) {
    if (!isValidNumber(data.minPrice)) {
      alert('Please enter numbers for min price');
      return false;
    }
  }

  if (typeof data.maxStock !== 'number' && data.maxStock !== null) {
    if (!isValidNumber(data.maxStock)) {
      alert('Please enter numbers for max stock');
      return false;
    }
  }

  if (typeof data.minStock !== 'number' && data.minStock !== null) {
    if (!isValidNumber(data.minStock)) {
      alert('Please enter numbers for min stock');
      return false;
    }
  }

  if (
    (!isInteger(data.maxPrice) && data.maxPrice !== null) ||
    (!isInteger(data.minPrice) && data.minPrice !== null) ||
    (!isInteger(data.maxStock) && data.maxStock !== null) ||
    (!isInteger(data.minStock) && data.minStock !== null)
  ) {
    alert("can't enter float numbers!");
    return false;
  }

  if (
    data.maxPrice !== null &&
    data.minPrice !== null &&
    parseInt(data.maxPrice.toString()) < parseInt(data.minPrice.toString())
  ) {
    alert("min price can't bigger than max value");
    return false;
  }

  if (
    data.maxStock !== null &&
    data.minStock !== null &&
    parseInt(data.maxStock.toString()) < parseInt(data.minStock.toString())
  ) {
    alert("min stock can't bigger than max value");
    return false;
  }

  return true;
};

const Search = () => {
  const [showPhoneFilter, setShowPhoneFilter] = useState(false);
  const phoneFilterOnClick = () => setShowPhoneFilter(!showPhoneFilter);

  const FILTER_OPTIONS: string[] = ['price', 'stock', 'sales', 'relevancy'];
  const ORDER_OPTIONS: string[] = ['asc', 'desc'];

  const { register, handleSubmit, reset } = useForm<FilterProps>({
    defaultValues: defaultFilterValues,
  });

  const onSubmit: SubmitHandler<FilterProps> = async (data: FilterProps) => {
    if (!isDataValid(data)) {
      return;
    } else {
      const newData: FilterProps = data;
      if (newData.minPrice) {
        newData.minPrice = parseInt(newData.minPrice.toString());
      }
      if (newData.maxPrice) {
        newData.maxPrice = parseInt(newData.maxPrice.toString());
      }
      if (newData.minStock) {
        newData.minStock = parseInt(newData.minStock.toString());
      }
      if (newData.maxStock) {
        newData.maxStock = parseInt(newData.maxStock.toString());
      }
      console.log('success');
    }
  };

  return (
    <div style={{ width: '100%' }}>
      <Row style={{ width: '100%' }}>
        <Col
          sm={2}
          md={2}
          className='filter_bar disappear_phone'
          style={{ padding: '8% 2% 5% 2%' }}
        >
          <form onSubmit={handleSubmit(onSubmit)}>
            <h4 className='title_color' style={{ paddingBottom: '20px' }}>
              <b>Product Filter</b>
            </h4>
            <h6 className='title_color'>Price Range</h6>
            <Row style={{ padding: '0px 0px 10px 23px' }}>
              <Col xs={2} className='center'>
                min
              </Col>
              <Col xs={10} className='center' style={{ paddingBottom: '5px' }}>
                <input
                  className='input_box no_spin_button'
                  type='number'
                  {...register('minPrice', { required: false })}
                />
              </Col>
              <Col xs={2} className='center'>
                max
              </Col>
              <Col xs={10} className='center'>
                <input
                  className='input_box no_spin_button'
                  type='number'
                  {...register('maxPrice', { required: false })}
                />
              </Col>
            </Row>
            <hr />
            <h6 className='title_color'>Stock Range</h6>
            <Row style={{ padding: '0px 0px 10px 23px' }}>
              <Col xs={2} className='center'>
                min
              </Col>
              <Col xs={10} className='center' style={{ paddingBottom: '5px' }}>
                <input
                  className='input_box no_spin_button'
                  type='number'
                  {...register('minStock', { required: false })}
                />
              </Col>
              <Col xs={2} className='center'>
                max
              </Col>
              <Col xs={10} className='center'>
                <input
                  className='input_box no_spin_button'
                  type='number'
                  {...register('maxStock', { required: false })}
                />
              </Col>
            </Row>
            <hr />
            <h6 className='title_color'>Coupon</h6>
            <Row style={{ padding: '0px 0px 10px 10px' }}>
              <Col xs={3} style={{ paddingBottom: '5px' }}>
                <input
                  className='input_box'
                  type='checkbox'
                  {...register('hasCoupon', { required: false })}
                />
              </Col>
              <Col xs={9}>has coupon</Col>
            </Row>
            <hr />
            <h6 className='title_color'>Sort by</h6>
            <div style={{ padding: '0px 0px 10px 15px' }}>
              <Form.Select
                style={{ backgroundColor: 'var(--button_dark)', color: 'white' }}
                className='selectBox'
                {...register('sort')}
                aria-label='Sort selection'
              >
                <option value='' disabled selected>
                  select
                </option>
                {FILTER_OPTIONS.map((s) => (
                  <option key={s} value={s}>
                    {s}
                  </option>
                ))}
              </Form.Select>
            </div>
            <hr />
            <h6 className='title_color'>Order by</h6>
            <div style={{ padding: '0px 0px 10px 15px' }}>
              <Form.Select
                style={{ backgroundColor: 'var(--button_dark)', color: 'white' }}
                className='selectBox'
                {...register('order')}
                aria-label='Order selection'
              >
                <option value='' disabled selected>
                  select
                </option>
                {ORDER_OPTIONS.map((o) => (
                  <option key={o} value={o}>
                    {o}
                  </option>
                ))}
              </Form.Select>
            </div>

            <hr />
            <TButton text='Reset' action={reset} />
            <TButton text='Submit' action={handleSubmit(onSubmit)} />
          </form>
        </Col>

        <Col sm={12} md={10} className='flex_wrapper' style={{ padding: '5% 8% 5% 8%' }}>
          <div className='title'>Shops : </div>
          <Row>
            {dataArray.map((data, index) => (
              <Col key={index} xs={6} md={3}>
                <SellerItem data={data} />
              </Col>
            ))}
          </Row>

          <div className='title'>Products : </div>
          <Row>
            {goodsData.map((data, index) => (
              <Col key={index} xs={6} md={3}>
                <GoodsItem {...data} />
              </Col>
            ))}
          </Row>
        </Col>
      </Row>
      <div className='disappear_desktop' onClick={phoneFilterOnClick}>
        <div className=' pointer center filter_icon '>
          <FontAwesomeIcon icon={faFilter} />
        </div>
      </div>

      <Offcanvas show={showPhoneFilter} onHide={phoneFilterOnClick} style={PhoneOffCanvasStyle}>
        <h4 className='title_color' style={{ paddingBottom: '20px' }}>
          <b>Product Filter</b>
        </h4>
        <h6 className='title_color'>Price Range</h6>
        <Row style={{ padding: '0px 0px 10px 23px' }}>
          <Col xs={2} className='center'>
            min
          </Col>
          <Col xs={10} className='center' style={{ paddingBottom: '5px' }}>
            <input
              className='input_box no_spin_button'
              type='number'
              {...register('minPrice', { required: false })}
            />
          </Col>
          <Col xs={2} className='center'>
            max
          </Col>
          <Col xs={10} className='center'>
            <input
              className='input_box no_spin_button'
              type='number'
              {...register('maxPrice', { required: false })}
            />
          </Col>
        </Row>
        <hr />
        <h6 className='title_color'>Stock Range</h6>
        <Row style={{ padding: '0px 0px 10px 23px' }}>
          <Col xs={2} className='center'>
            min
          </Col>
          <Col xs={10} className='center' style={{ paddingBottom: '5px' }}>
            <input
              className='input_box no_spin_button'
              type='number'
              {...register('minStock', { required: false })}
            />
          </Col>
          <Col xs={2} className='center'>
            max
          </Col>
          <Col xs={10} className='center'>
            <input
              className='input_box no_spin_button'
              type='number'
              {...register('maxStock', { required: false })}
            />
          </Col>
        </Row>
        <h6 className='title_color'>Coupon</h6>
        <Row style={{ padding: '0px 0px 10px 10px' }}>
          <Col xs={3} style={{ paddingBottom: '5px' }}>
            <input
              className='input_box'
              type='checkbox'
              {...register('hasCoupon', { required: false })}
            />
          </Col>
          <Col xs={9}>has coupon</Col>
        </Row>
        <hr />
        <h6 className='title_color'>Sort by</h6>
        <div style={{ padding: '0px 0px 10px 15px' }}>
          <Form.Select
            style={{ backgroundColor: 'var(--button_dark)', color: 'white' }}
            className='selectBox'
            {...register('sort')}
            aria-label='Sort selection'
          >
            <option value='' disabled selected>
              select
            </option>
            {FILTER_OPTIONS.map((s) => (
              <option key={s} value={s}>
                {s}
              </option>
            ))}
          </Form.Select>
        </div>
        <hr />
        <h6 className='title_color'>Order by</h6>
        <div style={{ padding: '0px 0px 10px 15px' }}>
          <Form.Select
            style={{ backgroundColor: 'var(--button_dark)', color: 'white' }}
            className='selectBox'
            {...register('order')}
            aria-label='Order selection'
          >
            <option value='' disabled selected>
              select
            </option>
            {ORDER_OPTIONS.map((o) => (
              <option key={o} value={o}>
                {o}
              </option>
            ))}
          </Form.Select>
        </div>
        <hr />
        <TButton text='Reset' action={reset} />
        <TButton text='Submit' action={handleSubmit(onSubmit)} />
      </Offcanvas>
    </div>
  );
};

export default Search;
