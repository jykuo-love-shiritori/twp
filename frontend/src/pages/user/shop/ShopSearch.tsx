import '@components/style.css';
import '@style/global.css';

import { Col, Form, Row, Offcanvas } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faFilter } from '@fortawesome/free-solid-svg-icons';
import { Navigate, useNavigate, useParams, useSearchParams } from 'react-router-dom';
import { SubmitHandler, useForm } from 'react-hook-form';
import { useState } from 'react';
import { useQuery } from '@tanstack/react-query';

import GoodsItem from '@components/GoodsItem';
import TButton from '@components/TButton';

import { PhoneOffCanvasStyle, FilterProps, ProductProps, isDataValid } from '@pages/search';
import { RouteOnNotOK } from '@lib/Status';

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

const ShopSearch = () => {
  const navigate = useNavigate();
  const [searchParams, setSearchParams] = useSearchParams();

  const [showPhoneFilter, setShowPhoneFilter] = useState(false);
  const phoneFilterOnClick = () => setShowPhoneFilter(!showPhoneFilter);

  const { sellerName } = useParams();
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
    queryKey: ['shop_search', searchParams.get('q'), sellerName],
    queryFn: async () => {
      const response = await fetch(`/api/shop/${sellerName}/search?` + searchParams.toString(), {
        headers: {
          Accept: 'application/json',
        },
      });
      if (!response.ok) {
        RouteOnNotOK(response, navigate);
      }
      return (await response.json()) as ProductProps[];
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
          sm={3}
          md={3}
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

        <Col sm={12} md={9} className='flex_wrapper' style={{ padding: '5% 6% 5% 6%' }}>
          <div className='title'>Products : </div>
          <Row>
            {data !== undefined && data.length !== 0 ? (
              data.map((d, index: number) => (
                <Col key={index} xs={6} md={4}>
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

export default ShopSearch;
