import GoodsItem from '@components/GoodsItem';
import SellerItem from '@components/SellerItem';
import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import { useQuery } from '@tanstack/react-query';
import { Col, Row } from 'react-bootstrap';
import { useForm } from 'react-hook-form';
import { useNavigate, useSearchParams } from 'react-router-dom';

interface IProduct {
  id: number;
  image_url: string;
  name: string;
  price: number;
  stock: number;
}

interface IShop {
  id: number;
  image_url: string;
  name: string;
  price: number;
  stock: number;
}

interface ISearchResult {
  products: IProduct[];
  shops: IShop[];
}

// empty string means no input (from react form)
interface IFrom {
  minPrice: number | '';
  maxPrice: number | '';
  minStock: number | '';
  maxStock: number | '';
  haveCoupon: boolean;
  sortBy: 'price' | 'stock' | 'sales' | 'relevancy';
  order: 'asc' | 'desc';
}

const defaultFrom: IFrom = {
  minPrice: '',
  maxPrice: '',
  minStock: '',
  maxStock: '',
  haveCoupon: false,
  sortBy: 'relevancy',
  order: 'desc',
};

const Search = () => {
  const [searchParams, setSearchParams] = useSearchParams();
  const navigate = useNavigate();

  const itemLimit = 10;
  const { data, status } = useQuery({
    queryKey: ['search', searchParams.toString()],
    queryFn: async () => {
      // if first time, reset search params
      if (!searchParams.has('limit') || searchParams.get('limit') !== (itemLimit + 1).toString()) {
        const newSearchParams = new URLSearchParams({
          q: searchParams.get('q') ?? '',
          offset: '0',
          limit: (itemLimit + 1).toString(),
        });
        setSearchParams(newSearchParams, { replace: true });
        reset(defaultFrom);
      }

      const resp = await fetch('/api/search?' + searchParams.toString(), {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      RouteOnNotOK(resp, navigate);
      const response = (await resp.json()) as ISearchResult;
      if (response.products.length === itemLimit + 1) {
        response.products.pop();
        // TODO
        console.log('isMore');
      } else {
        // TODO
        console.log('noMore');
      }
      return response;
    },
    enabled: true,
    refetchOnWindowFocus: false,
    retry: false,
  });

  const { register, handleSubmit, reset } = useForm<IFrom>({ defaultValues: defaultFrom });

  const onSubmit = (data: IFrom) => {
    console.log(data);

    // fill other side of range if one is given
    if (data.minPrice !== '' && data.maxPrice === '') data.maxPrice = 999999999;
    if (data.minPrice === '' && data.maxPrice !== '') data.minPrice = 0;
    if (data.minStock !== '' && data.maxStock === '') data.maxStock = 999999999;
    if (data.minStock === '' && data.maxStock !== '') data.minStock = 0;

    // check valid if both exist
    if (data.minPrice !== '' && data.maxPrice !== '') {
      if (data.minPrice > data.maxPrice || data.minPrice < 0 || data.maxPrice < 0) {
        alert('invalid price range');
        return;
      }
    }
    if (data.minStock !== '' && data.maxStock !== '') {
      if (data.minStock > data.maxStock || data.minStock < 0 || data.maxStock < 0) {
        alert('invalid stock range');
        return;
      }
    }

    // set search params
    if (data.minPrice === '') searchParams.delete('minPrice');
    else searchParams.set('minPrice', data.minPrice.toString());

    if (data.maxPrice === '') searchParams.delete('maxPrice');
    else searchParams.set('maxPrice', data.maxPrice.toString());

    if (data.minStock === '') searchParams.delete('minStock');
    else searchParams.set('minStock', data.minStock.toString());

    if (data.maxStock === '') searchParams.delete('maxStock');
    else searchParams.set('maxStock', data.maxStock.toString());

    searchParams.set('haveCoupon', data.haveCoupon.toString());
    searchParams.set('sortBy', data.sortBy);
    searchParams.set('order', data.order);

    setSearchParams(searchParams, { replace: true });
  };

  if (status !== 'success') {
    return <CheckFetchStatus status={status} />;
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
              {/* Price Min */}
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

              {/* Price Max */}
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
              {/* Stock Min */}
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

              {/* Stock Max */}
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
              {/* Has Coupon */}
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

            {/* Sort By */}
            <h6 className='title_color'>Sort by</h6>
            <div style={{ padding: '0px 0px 10px 15px' }}>
              <select className='selectBox' {...register('sortBy')}>
                <option value='price'>Price</option>
                <option value='stock'>Stock</option>
                <option value='sales'>Sales</option>
                <option value='relevancy'>Relevancy</option>
              </select>
            </div>
            <hr />

            {/* Order By */}
            <h6 className='title_color'>Order by</h6>
            <div style={{ padding: '0px 0px 10px 15px' }}>
              <select className='selectBox' {...register('order')}>
                <option value='asc'>asc</option>
                <option value='desc'>desc</option>
              </select>
            </div>
            <hr />

            <input type='submit' value='Submit' />
          </form>
        </Col>

        <Col sm={12} md={10} className='flex_wrapper' style={{ padding: '5% 8% 5% 8%' }}>
          {/* Shop Result */}
          <div className='title'>Shops : </div>
          <Row>
            {data.shops.length !== 0 ? (
              data.shops.map((shop, index: number) => (
                <Col key={index} xs={6} md={3}>
                  <SellerItem
                    data={{ name: shop.name, image_url: shop.image_url, seller_name: shop.name }}
                  />
                </Col>
              ))
            ) : (
              <div>Not found</div>
            )}
          </Row>

          {/* Product Result */}
          <div className='title'>Products : </div>
          <Row>
            {data.products.length !== 0 ? (
              data.products.map((product, index: number) => (
                <Col key={index} xs={6} md={3}>
                  <GoodsItem {...product} />
                </Col>
              ))
            ) : (
              <div>Not found</div>
            )}
          </Row>
        </Col>
      </Row>
    </div>
  );
};

export default Search;
