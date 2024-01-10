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
import Pagination from '@components/Pagination';

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
  price: number;
  sales: number;
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
  if (typeof data.maxPrice !== 'number' && data.maxPrice !== null && data.maxPrice !== '') {
    if (!isValidNumber(data.maxPrice)) {
      alert('Please enter max price');
      return false;
    }
  }

  if (typeof data.minPrice !== 'number' && data.minPrice !== null && data.minPrice !== '') {
    if (!isValidNumber(data.minPrice)) {
      alert('Please enter min price');
      return false;
    }
  }

  if (typeof data.maxStock !== 'number' && data.maxStock !== null && data.maxStock !== '') {
    if (!isValidNumber(data.maxStock)) {
      alert('Please enter max stock');
      return false;
    }
  }

  if (typeof data.minStock !== 'number' && data.minStock !== null && data.minStock !== '') {
    if (!isValidNumber(data.minStock)) {
      alert('Please enter min stock');
      return false;
    }
  }

  if (
    (!isInteger(data.maxPrice) && data.maxPrice !== null && data.maxPrice.toString() !== '') ||
    (!isInteger(data.minPrice) && data.minPrice !== null && data.minPrice.toString() !== '') ||
    (!isInteger(data.maxStock) && data.maxStock !== null && data.maxStock.toString() !== '') ||
    (!isInteger(data.minStock) && data.minStock !== null && data.minStock.toString() !== '')
  ) {
    alert('please enter integers');
    return false;
  }

  if (data.minPrice !== null && parseInt(data.minPrice.toString()) < 0) {
    alert("min price can't be negative");
    return false;
  }

  if (data.maxPrice !== null && parseInt(data.maxPrice.toString()) < 0) {
    alert("max price can't be negative");
    return false;
  }

  if (data.minStock !== null && parseInt(data.minStock.toString()) < 0) {
    alert("min stock can't be negative");
    return false;
  }

  if (data.maxStock !== null && parseInt(data.maxStock.toString()) < 0) {
    alert("max stock can't be negative");
    return false;
  }

  if (data.minPrice !== null && parseInt(data.minPrice.toString()) < 0) {
    alert("min price can't be negative");
    return false;
  }

  if (data.maxPrice !== null && parseInt(data.maxPrice.toString()) < 0) {
    alert("max price can't be negative");
    return false;
  }

  if (
    data.maxPrice !== null &&
    data.minPrice !== null &&
    parseInt(data.maxPrice.toString()) < parseInt(data.minPrice.toString())
  ) {
    alert('min price is greater than max price');
    return false;
  }

  if (
    data.maxStock !== null &&
    data.minStock !== null &&
    parseInt(data.maxStock.toString()) < parseInt(data.minStock.toString())
  ) {
    alert('min stock is greater than max stock');
    return false;
  }

  if (isEmpty(data.maxPrice) != isEmpty(data.minPrice)) {
    alert('please enter max/min price when the other one is present');
    return false;
  }

  if (isEmpty(data.maxStock) != isEmpty(data.minStock)) {
    alert('please enter max/min stock when the other one is present');
    return false;
  }

  return true;
};

const Search = () => {
  const navigate = useNavigate();
  const [searchParams, setSearchParams] = useSearchParams();

  const [showPhoneFilter, setShowPhoneFilter] = useState(false);
  const phoneFilterOnClick = () => setShowPhoneFilter(!showPhoneFilter);

  const [isMore, setIsMore] = useState(true);
  const itemLimit = 8;

  if (!searchParams.has('offset') || Number(searchParams.get('limit')) !== itemLimit + 1) {
    const newSearchParams = new URLSearchParams({
      q: searchParams.get('q') ?? '', // q shouldnt be null here
      offset: '0',
      limit: (itemLimit + 1).toString(),
    });
    setSearchParams(newSearchParams);
  }

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
      const data = (await response.json()) as ResponseProps;
      if (data.products.length === itemLimit + 1) {
        setIsMore(true);
        data.products.pop();
      } else {
        setIsMore(false);
      }
      return data;
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
    params.set('offset', '0');
    params.set('limit', (itemLimit + 1).toString());

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
                  <GoodsItem
                    id={d.id}
                    name={d.name}
                    image_url={d.image_url}
                    price={d.price}
                    sales={d.sales}
                  />
                </Col>
              ))
            ) : (
              <div>Not found</div>
            )}
          </Row>
          <div className='center' style={{ padding: '2% 0px' }}>
            <Pagination limit={itemLimit} isMore={isMore} />
          </div>
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
