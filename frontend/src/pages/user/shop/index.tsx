import { Col, Row } from 'react-bootstrap';
import { useQuery } from '@tanstack/react-query';
import { Outlet, useParams } from 'react-router-dom';
import { useNavigate } from 'react-router-dom';
import { CSSProperties } from 'react';

import TButton from '@components/TButton';
import { RouteOnNotOK } from '@lib/Status';
import NotFound from '@components/NotFound';

interface UserViewShopProps {
  info: {
    description: string;
    image_url: string;
    name: string;
    seller_name: string;
  };
  products: {
    description: string;
    expire_data: string;
    id: number;
    image_url: string;
    name: string;
    price: number;
    sales: number;
    stock: number;
  };
}

const userImgStyle: CSSProperties = {
  borderRadius: '50%',
  minHeight: '23vh',
  minWidth: '23vh',
  maxHeight: '23vh',
  maxWidth: '23vh',
  objectFit: 'cover',
  cursor: 'pointer',
  boxShadow: '2px 4px 10px 2px rgba(0, 0, 0, 0.25)',
};

const UserViewShop = () => {
  const { sellerName } = useParams();
  const navigate = useNavigate();

  const { data: userViewShop } = useQuery({
    queryKey: ['userViewShop', sellerName],
    queryFn: async () => {
      const resp = await fetch(`/api/shop/${sellerName}?offset=0&limit=12`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
      if (!resp.ok) {
        RouteOnNotOK(resp, navigate);
      }
      return resp.json();
    },
    select: (data) => data as UserViewShopProps,
  });

  if (!userViewShop) {
    return <NotFound />;
  } else {
    return (
      <Row style={{ width: '100%', padding: '0', margin: '0' }} className='flex_wrapper'>
        <Col xs={12} md={12} style={{ width: '100%', padding: '0' }}>
          <div className='user_bg center'>
            <div style={{ padding: '6% 10% 6% 10%' }}>{userViewShop.info.description}</div>
          </div>
        </Col>
        <Col xs={12} md={12} lg={12} style={{ padding: '0' }}>
          <Row style={{ padding: '0', margin: '0', width: '100%' }}>
            <Col xs={12} md={3} lg={2}>
              <Row className='user_icon'>
                <Col xs={12} className='center' style={{ overflow: 'hidden' }}>
                  <img src={userViewShop.info.image_url} style={userImgStyle} />
                </Col>
                <Col xs={12}>
                  <div className='center'>
                    <h4 className='title_color' style={{ padding: '10% 2% 10% 2%' }}>
                      <b>{userViewShop.info.name}</b>
                    </h4>
                  </div>
                  <TButton text='Explore Shop' action='/sellerID/shop' />
                  <TButton text='Check Coupons' action='/sellerID/coupons' />
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
  }
};

export default UserViewShop;
