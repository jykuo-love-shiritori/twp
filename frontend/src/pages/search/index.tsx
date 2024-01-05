import '@components/style.css';
import '@style/global.css';

import { Col, Form, Row, Offcanvas } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faFilter } from '@fortawesome/free-solid-svg-icons';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { SubmitHandler, useForm } from 'react-hook-form';
import { useState, useEffect } from 'react';
import { useMutation } from '@tanstack/react-query';

import SellerItem from '@components/SellerItem';
import GoodsItem from '@components/GoodsItem';
import TButton from '@components/TButton';
import { getUrl } from '@components/SearchBar';

const PhoneOffCanvasStyle = {
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

interface ResponseProps {
  products: ProductProps[];
  shops: ShopProps[];
}

interface ProductProps {
  id: number;
  image_url: string;
  name: string;
  price: string;
  stock: number;
}

interface ShopProps {
  id: number;
  image_url: string;
  name: string;
  price: number;
  stock: number;
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

const isDataValid = (data: FilterProps) => {
  if (typeof data.maxPrice !== 'number' && data.maxPrice !== null && data.maxPrice !== '') {
    if (!isValidNumber(data.maxPrice)) {
      alert('Please enter numbers for max price');
      return false;
    }
  }

  if (typeof data.minPrice !== 'number' && data.minPrice !== null && data.minPrice !== '') {
    if (!isValidNumber(data.minPrice)) {
      alert('Please enter numbers for min price');
      return false;
    }
  }

  if (typeof data.maxStock !== 'number' && data.maxStock !== null && data.maxStock !== '') {
    if (!isValidNumber(data.maxStock)) {
      alert('Please enter numbers for max stock');
      return false;
    }
  }

  if (typeof data.minStock !== 'number' && data.minStock !== null && data.minStock !== '') {
    if (!isValidNumber(data.minStock)) {
      alert('Please enter numbers for min stock');
      return false;
    }
  }

  if (
    (!isInteger(data.maxPrice) && data.maxPrice !== null && data.maxPrice.toString() !== '') ||
    (!isInteger(data.minPrice) && data.minPrice !== null && data.minPrice.toString() !== '') ||
    (!isInteger(data.maxStock) && data.maxStock !== null && data.maxStock.toString() !== '') ||
    (!isInteger(data.minStock) && data.minStock !== null && data.minStock.toString() !== '')
  ) {
    alert("can't enter float numbers!");
    return false;
  }

  if (data.minPrice !== null && parseInt(data.minPrice.toString()) < 0) {
    alert("min price can't negative numbers");
    return false;
  }

  if (data.maxPrice !== null && parseInt(data.maxPrice.toString()) < 0) {
    alert("max price can't negative numbers");
    return false;
  }

  if (data.minStock !== null && parseInt(data.minStock.toString()) < 0) {
    alert("min stock can't negative numbers");
    return false;
  }

  if (data.maxStock !== null && parseInt(data.maxStock.toString()) < 0) {
    alert("max stock can't negative numbers");
    return false;
  }

  if (data.minPrice !== null && parseInt(data.minPrice.toString()) < 0) {
    alert("min price can't negative numbers");
    return false;
  }

  if (data.maxPrice !== null && parseInt(data.maxPrice.toString()) < 0) {
    alert("max price can't negative numbers");
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

  if (isEmpty(data.maxPrice) != isEmpty(data.minPrice)) {
    alert('max and min price should both have values or both have no values at the same time');
    return false;
  }

  if (isEmpty(data.maxStock) != isEmpty(data.minStock)) {
    alert('max and min stock should both have values or both have no values at the same time');
    return false;
  }

  return true;
};

const Search = () => {
  const [searchParams] = useSearchParams();
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
    mutationFn: async (requestString: string) => {
      const currentQ = searchParams.get('q') ?? '';
      if (currentQ === '') {
        return;
      }
      const response = await fetch(`/api${requestString}`, {
        headers: {
          Accept: 'application/json',
        },
      });
      if (!response.ok) {
        throw new Error('query search failed');
      }
      return (await response.json()) as ResponseProps;
    },
    onError: () => {
      navigate('/searchNotFound');
    },
  });

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
    const requestUrl = getUrl(request, newQ);

    setQ(request.q);
    setValue('minPrice', request.minPrice);
    setValue('maxPrice', request.maxPrice);
    setValue('minStock', request.minStock);
    setValue('maxStock', request.maxStock);
    setValue('haveCoupon', request.haveCoupon);
    setValue('sortBy', request.sortBy);
    setValue('order', request.order);
    querySearch.mutate(requestUrl);

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
      navigate(getUrl(newData, newQ));
    }
  };

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
            {querySearch.data?.shops.length !== 0 && querySearch.data !== undefined ? (
              querySearch.data?.shops.map((data, index: number) => (
                <Col key={index} xs={6} md={3}>
                  <SellerItem data={{ name: data.name, image_url: data.image_url }} />
                </Col>
              ))
            ) : (
              <div>Not found</div>
            )}
          </Row>

          <div className='title'>Products : </div>
          <Row>
            {querySearch.data?.products.length !== 0 && querySearch.data !== undefined ? (
              querySearch.data?.products.map((data, index: number) => (
                <Col key={index} xs={6} md={3}>
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

export default Search;
