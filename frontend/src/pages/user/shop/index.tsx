import { Col, Row } from 'react-bootstrap';
import { Outlet, useNavigate, useParams } from 'react-router-dom';

import TButton from '@components/TButton';
import { useAuth } from '@lib/Auth';
import { useQuery } from '@tanstack/react-query';
import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';

interface IProduct {
  description: string;
  expire_date: string;
  id: number;
  image_url: string;
  name: string;
  price: number;
  sales: number;
  stock: number;
}

interface IShop {
  info: {
    description: string;
    image_url: string;
    name: string;
    seller_name: string;
  };
  products: IProduct[];
}

const UserViewShop = () => {
  const token = useAuth();
  const navigate = useNavigate();
  const { sellerName } = useParams();

  const { status, data } = useQuery({
    queryKey: ['getShopInfoForShopInfo', sellerName],
    queryFn: async () => {
      const resp = await fetch(`/api/shop/${sellerName}?offset=${0}&limit=${8}`, {
        method: 'GET',
        headers: {
          Accept: 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });
      RouteOnNotOK(resp, navigate);
      return (await resp.json()) as IShop;
    },
    enabled: true,
    refetchOnWindowFocus: false,
    retry: false,
  });

  if (status != 'success') {
    return <CheckFetchStatus status={status} />;
  }

  const shop = data.info;
  const products = data.products;

  return (
    <Row style={{ width: '100%', padding: '0', margin: '0' }} className='flex_wrapper'>
      <Col xs={12} md={12} style={{ width: '100%', padding: '0' }}>
        <div className='user_bg center'>
          <div style={{ padding: '6% 10% 6% 10%' }}>{shop.description}</div>
        </div>
      </Col>
      <Col xs={12} md={12} lg={12} style={{ padding: '0' }}>
        <Row style={{ padding: '0', margin: '0', width: '100%' }}>
          <Col xs={12} md={3} lg={2}>
            <Row className='user_icon'>
              <Col xs={12} className='center'>
                <img src={shop.image_url} className='user_img' />
              </Col>
              <Col xs={12}>
                <div className='center'>
                  <h4 className='title_color' style={{ padding: '10% 2% 0% 2%' }}>
                    <b>{shop.name}</b>
                  </h4>
                </div>
                <hr className='hr' />
                <div className='center'> Products : {products.length} items</div>
                <TButton text='Explore Shop' action={`/${sellerName}/shop`} />
                <TButton text='Check Coupons' action={`/${sellerName}/coupons`} />
              </Col>
            </Row>
          </Col>
          <Col xs={12} md={9} ld={10} style={{ padding: '1% 5% 6% 5%' }}>
            <Outlet />
          </Col>
        </Row>
      </Col>
    </Row>
  );
};

export default UserViewShop;
