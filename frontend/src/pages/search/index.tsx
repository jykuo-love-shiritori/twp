import '@components/style.css';
import '@style/global.css';

import { Col, Form, Row, Offcanvas } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faFilter } from '@fortawesome/free-solid-svg-icons';
import { Navigate, useNavigate, useSearchParams } from 'react-router-dom';
import { SubmitHandler, useForm } from 'react-hook-form';
import { useState } from 'react';
import { useQuery } from '@tanstack/react-query';

import SellerItem from '@components/SellerItem';
import GoodsItem from '@components/GoodsItem';
import TButton from '@components/TButton';
import { RouteOnNotOK } from '@lib/Status';

export const PhoneOffCanvasStyle = {
  width: '330px',
  border: ' 1px solid var(--button_light, #60998B)',
  background: ' var(--bg, #061E18)',
  boxShadow: '0px 4px 20px 2px #60998B',
  padding: '10%',
  color: 'white',
};

export interface FilterProps {
  maxPrice: number | null;
  minPrice: number | null;
  maxStock: number | null;
  minStock: number | null;
  haveCoupon: boolean | null;
  sortBy: 'price' | 'stock' | 'sales' | 'relevancy' | null;
  order: 'asc' | 'desc' | null;
}

export interface SearchProps extends FilterProps {
  q: string;
}

export interface ResponseProps {
  products: ProductProps[];
  shops: ShopProps[];
}

export interface ProductProps {
  id: number;
  image_url: string;
  name: string;
  price: string;
  stock: number;
}

export interface ShopProps {
  id: number;
  image_url: string;
  name: string;
  seller_name: string;
}

// eslint-disable-next-line react-refresh/only-export-components
export const defaultFilterValues: FilterProps = {
  maxPrice: null,
  minPrice: null,
  maxStock: null,
  minStock: null,
  haveCoupon: null,
  sortBy: null,
  order: null,
};

const isValidNumber = (value: string | number | null | undefined): boolean => {
  return typeof value === 'number' || (typeof value === 'string' && !isNaN(parseInt(value)));
};

const isInteger = (value: string | number | null | undefined): boolean => {
  return typeof value === 'number' || (typeof value === 'string' && /^-?\d+$/.test(value));
};

const isEmpty = (value: string | number | null | undefined): boolean => {
  return value === '' || value === null ? true : false;
};

const toNumber = (input: string | null) => {
  if (!input) return null;
  const output = Number(input);
  return isNaN(output) ? null : output;
};

const toSortBy = (input: string | null) => {
  switch (input) {
    case 'price':
      return 'price';
    case 'stock':
      return 'stock';
    case 'sales':
      return 'sales';
    case 'relevancy':
      return 'relevancy';
    default:
      return null;
  }
};

const toOrderBy = (input: string | null) => {
  switch (input) {
    case 'asc':
      return 'asc';
    case 'desc':
      return 'desc';
    default:
      return null;
  }
};

// eslint-disable-next-line react-refresh/only-export-components
export const isDataValid = (data: FilterProps) => {
  const AlertWithReturn = (message: string) => {
    alert(message);
    return false;
  };

  const isNumber = (value: number | string | null) => typeof value === 'number';
  const isNull = (value: number | string | null) => value === null;
  const isEmptyString = (value: number | string | null) => value === '';
  const isEmpty = (value: number | string | null) => isNull(value) || isEmptyString(value);
  const isEntered = (value: number | string | null) => !isNumber(value) && !isEmpty(value);

  const isMinBiggerThanMax = (min: number | string | null, max: number | string | null) =>
    isEntered(min) &&
    isEntered(max) &&
    min !== null &&
    max !== null &&
    Number(min.toString()) > Number(max.toString());

  const isNegative = (value: number | string | null) =>
    value !== null && parseInt(value.toString()) < 0;

  if (isEntered(data.maxPrice)) {
    if (!isValidNumber(data.maxPrice))
      return AlertWithReturn('Please enter numbers for the max price.');
    if (isNegative(data.maxPrice)) return AlertWithReturn("Max price can't be negative numbers.");
  }

  if (isEntered(data.minPrice)) {
    if (!isValidNumber(data.minPrice))
      return AlertWithReturn('Please enter numbers for the min price.');
    if (isNegative(data.minPrice)) return AlertWithReturn("Min price can't be negative numbers.");
  }

  if (isEntered(data.maxStock)) {
    if (!isValidNumber(data.maxStock))
      return AlertWithReturn('Please enter numbers for the max stock.');
    if (isNegative(data.maxStock)) return AlertWithReturn("Max stock can't be negative numbers.");
  }

  if (isEntered(data.minStock)) {
    if (!isValidNumber(data.minStock))
      return AlertWithReturn('Please enter numbers for the min stock.');
    if (isNegative(data.minStock)) return AlertWithReturn("Min stock can't be negative numbers.");
  }

  const isFloat = (value: number | string | null) =>
    !isInteger(value) && value !== null && value.toString() !== '';

  if ([data.maxPrice, data.minPrice, data.maxStock, data.minStock].some(isFloat)) {
    alert("Can't enter float numbers!");
    return false;
  }

  if (isMinBiggerThanMax(data.minPrice, data.maxPrice)) {
    return AlertWithReturn("Min price can't be bigger than max values.");
  }

  if (isMinBiggerThanMax(data.minStock, data.maxStock)) {
    return AlertWithReturn("Min stock can't be bigger than max value.");
  }

  if (isEntered(data.maxPrice) !== isEntered(data.minPrice)) {
    return AlertWithReturn(
      'Max and min price should both have valid values or both have no values at the same time.',
    );
  }

  if (isEntered(data.maxStock) !== isEntered(data.minStock)) {
    return AlertWithReturn(
      'Max and min stock should both have values or both have no values at the same time.',
    );
  }

  return true;
};

const Search = () => {
  const navigate = useNavigate();
  const [searchParams, setSearchParams] = useSearchParams();

  const [showPhoneFilter, setShowPhoneFilter] = useState(false);
  const phoneFilterOnClick = () => setShowPhoneFilter(!showPhoneFilter);

  const FILTER_OPTIONS: string[] = ['price', 'stock', 'sales', 'relevancy'];
  const ORDER_OPTIONS: string[] = ['asc', 'desc'];

  const { register, handleSubmit, reset } = useForm<FilterProps>({
    defaultValues: {
      minPrice: toNumber(searchParams.get('minPrice')),
      maxPrice: toNumber(searchParams.get('maxPrice')),
      minStock: toNumber(searchParams.get('minStock')),
      maxStock: toNumber(searchParams.get('maxStock')),
      haveCoupon: Boolean(searchParams.get('haveCoupon')),
      sortBy: toSortBy(searchParams.get('sortBy')),
      order: toOrderBy(searchParams.get('order')),
    },
  });

  const { data, isError } = useQuery({
    queryKey: ['search', searchParams.toString()],
    queryFn: async () => {
      const response = await fetch('/api/search?' + searchParams.toString(), {
        headers: {
          Accept: 'application/json',
        },
      });
      if (!response.ok) {
        RouteOnNotOK(response, navigate);
      }
      return (await response.json()) as ResponseProps;
    },
  });

  const onSubmit: SubmitHandler<FilterProps> = async (data, e) => {
    e?.preventDefault();

    if (!isDataValid(data)) {
      return;
    }

    const params = new URLSearchParams();

    const q = searchParams.get('q') ?? '';
    if (!q) {
      return;
    }

    params.set('q', q);
    if (data.minPrice) params.set('minPrice', data.minPrice.toString());
    if (data.maxPrice) params.set('maxPrice', data.maxPrice.toString());
    if (data.minStock) params.set('minStock', data.minStock.toString());
    if (data.maxStock) params.set('maxStock', data.maxStock.toString());
    if (data.haveCoupon) params.set('haveCoupon', data.haveCoupon.toString());
    if (data.sortBy) params.set('sortBy', data.sortBy);
    if (data.order) params.set('order', data.order);

    setSearchParams(params);
  };

  if (isError) {
    return <Navigate to='/notFound' />;
  }

  return (
    <div style={{ width: '100%', minHeight: '100vh' }}>
      <Row style={{ width: '100%', margin: '0' }}>
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
                  {...register('haveCoupon', { required: false })}
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
                {...register('sortBy')}
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
            {data && data.shops.length !== 0 ? (
              data.shops.map((d, index: number) => (
                <Col key={index} xs={6} md={3}>
                  <SellerItem
                    data={{ name: d.name, image_url: d.image_url, seller_name: d.seller_name }}
                  />
                </Col>
              ))
            ) : (
              <div>Not found</div>
            )}
          </Row>

          <div className='title'>Products : </div>
          <Row>
            {data && data.products.length !== 0 ? (
              data.products.map((d, index: number) => (
                <Col key={index} xs={6} md={3}>
                  <GoodsItem {...d} />
                </Col>
              ))
            ) : (
              <div>Not found</div>
            )}
          </Row>
        </Col>
      </Row>
      <div className='disappear_desktop' onClick={phoneFilterOnClick}>
        <div className=' pointer center filter_icon '>
          <FontAwesomeIcon icon={faFilter} />
        </div>
      </div>

      <Offcanvas show={showPhoneFilter} onHide={phoneFilterOnClick} style={PhoneOffCanvasStyle}>
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
          <h6 className='title_color'>Coupon</h6>
          <Row style={{ padding: '0px 0px 10px 10px' }}>
            <Col xs={3} style={{ paddingBottom: '5px' }}>
              <input
                className='input_box'
                type='checkbox'
                {...register('haveCoupon', { required: false })}
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
              {...register('sortBy')}
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
          <div onClick={phoneFilterOnClick}>
            <TButton text='Submit' action={handleSubmit(onSubmit)} />
          </div>
        </form>
      </Offcanvas>
    </div>
  );
};

export default Search;
