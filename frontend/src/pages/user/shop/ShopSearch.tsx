import '@components/style.css';
import '@style/global.css';

import { Col, Form, Row, Offcanvas } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faFilter } from '@fortawesome/free-solid-svg-icons';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { SubmitHandler, useForm } from 'react-hook-form';
import { useState, useEffect } from 'react';
import { useMutation } from '@tanstack/react-query';

import GoodsItem from '@components/GoodsItem';
import TButton from '@components/TButton';
import { getShopSearchUrl } from '@components/SearchBar';

import {
  PhoneOffCanvasStyle,
  FilterProps,
  SearchProps,
  ProductProps,
  defaultFilterValues,
  isDataValid,
} from '@pages/search';

const ShopSearch = () => {
  const [searchParams, setSearchParams] = useSearchParams();
  const navigate = useNavigate();
  const [q, setQ] = useState<string>(searchParams.get('q') ?? '');

  const [showPhoneFilter, setShowPhoneFilter] = useState(false);
  const phoneFilterOnClick = () => setShowPhoneFilter(!showPhoneFilter);

  const FILTER_OPTIONS: string[] = ['price', 'stock', 'sales', 'relevancy'];
  const ORDER_OPTIONS: string[] = ['asc', 'desc'];

  const { register, handleSubmit, reset, setValue } = useForm<FilterProps>({
    defaultValues: defaultFilterValues,
  });

  const querySearch = useMutation({
    mutationFn: async () => {
      const currentQ = searchParams.get('q') ?? '';
      if (currentQ === '') {
        return;
      }

      const response = await fetch('/api/shop/s/search?' + getNewParams(), {
        headers: {
          Accept: 'application/json',
        },
      });
      if (!response.ok) {
        throw new Error('query search failed');
      }
      return (await response.json()) as ProductProps[];
    },
    onError: () => {
      navigate('/searchNotFound');
    },
  });

  const getNewParams = () => {
    const minPrice = searchParams.get('minPrice');
    const maxPrice = searchParams.get('maxPrice');
    const minStock = searchParams.get('minStock');
    const maxStock = searchParams.get('maxStock');
    const haveCouponValue = searchParams.get('haveCoupon');
    const sortBy = searchParams.get('sortBy');
    const order = searchParams.get('order');

    if (minPrice) searchParams.set('minPrice', minPrice);
    if (maxPrice) searchParams.set('maxPrice', maxPrice);
    if (minStock) searchParams.set('minStock', minStock);
    if (maxStock) searchParams.set('maxStock', maxStock);
    if (haveCouponValue === 'true' || haveCouponValue === 'false') {
      searchParams.set('haveCoupon', haveCouponValue);
    }
    if (sortBy && ['price', 'stock', 'sales', 'relevancy'].includes(sortBy)) {
      searchParams.set('sortBy', sortBy);
    }
    if (order && ['asc', 'desc'].includes(order)) {
      searchParams.set('order', order);
    }

    setSearchParams(searchParams);

    return searchParams.toString();
  };

  useEffect(() => {
    const newQ = searchParams.get('q') ?? '';
    if (newQ !== q) {
      setQ(newQ);
    }
  }, [searchParams, q]);

  useEffect(() => {
    const newQ = searchParams.get('q') ?? '';
    if (newQ === '') {
      return;
    }

    const request: SearchProps = setData();

    setQ(request.q);
    setValue('minPrice', request.minPrice);
    setValue('maxPrice', request.maxPrice);
    setValue('minStock', request.minStock);
    setValue('maxStock', request.maxStock);
    setValue('haveCoupon', request.haveCoupon);
    setValue('sortBy', request.sortBy);
    setValue('order', request.order);
    querySearch.mutate();

    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [searchParams, setValue]);

  const setData = () => {
    const data: SearchProps = { ...defaultFilterValues, q: q };

    const maxPrice = parseInt(searchParams.get('maxPrice') || '');
    if (!isNaN(maxPrice)) {
      data.maxPrice = maxPrice;
    }

    const minPrice = parseInt(searchParams.get('minPrice') || '');
    if (!isNaN(minPrice)) {
      data.minPrice = minPrice;
    }

    const maxStock = parseInt(searchParams.get('maxStock') || '');
    if (!isNaN(maxStock)) {
      data.maxStock = maxStock;
    }

    const minStock = parseInt(searchParams.get('minStock') || '');
    if (!isNaN(minStock)) {
      data.minStock = minStock;
    }

    const haveCouponValue = searchParams.get('haveCoupon');
    if (haveCouponValue !== null && haveCouponValue !== '') {
      if (haveCouponValue === 'true' || haveCouponValue === 'false') {
        data.haveCoupon = JSON.parse(haveCouponValue);
      }
    }

    const sortBy = searchParams.get('sortBy');
    if (
      sortBy !== null &&
      sortBy !== '' &&
      (sortBy === 'price' || sortBy === 'stock' || sortBy === 'sales' || sortBy === 'relevancy')
    ) {
      data.sortBy = sortBy;
    }

    const order = searchParams.get('order');
    if (order !== null && order !== '' && (order === 'asc' || order === 'desc')) {
      data.order = order;
    }

    return data;
  };

  const onSubmit: SubmitHandler<FilterProps> = async (data: FilterProps) => {
    if (!isDataValid(data)) {
      return;
    } else {
      const newQ = searchParams.get('q') ?? '';
      const newData: SearchProps = { ...data, q: newQ };
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
      // idk why but when you click the checkbox with phone, it gives you haveCoupon: ['on']
      if (newData.haveCoupon !== null && newData.haveCoupon.toString() === 'on') {
        newData.haveCoupon = true;
      }
      navigate(getShopSearchUrl(newData, newQ));
    }
  };

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
            {querySearch.data?.length !== 0 && querySearch.data !== undefined ? (
              querySearch.data?.map((data, index: number) => (
                <Col key={index} xs={6} md={4}>
                  <GoodsItem {...data} />
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
