import { Col, Row } from 'react-bootstrap';
import { Outlet, useNavigate } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';

import { IShopInfo } from '@pages/user/seller/info';
import { useAuth } from '@lib/Auth';
import { RouteOnNotOK } from '@lib/Status';
import SellerButtons from '@components/SellerButtons';
import NotFound from '@components/NotFound';

const Seller = () => {
  const token = useAuth();
  const navigate = useNavigate();

  const { data: sellerInfo } = useQuery({
    queryKey: ['sellerGetShopInfo'],
    queryFn: async () => {
      const resp = await fetch('/api/seller/info', {
        method: 'GET',
        headers: {
          Authorization: `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
      });
      if (!resp.ok) {
        RouteOnNotOK(resp, navigate);
      }
      return resp.json();
    },
    select: (data) => data as IShopInfo,
  });

  if (!sellerInfo) {
    return <NotFound />;
  } else {
    return (
      <Row style={{ width: '100%', padding: '0', margin: '0' }} className='flex_wrapper'>
        <Col xs={12} md={12} style={{ width: '100%', padding: '0' }}>
          <div className='user_bg center' />
        </Col>
        <Col xs={12} md={12} lg={12} style={{ padding: '0' }}>
          <Row style={{ padding: '0', margin: '0', width: '100%' }}>
            <Col xs={12} md={3} lg={2} style={{ padding: '0' }}>
              <Row className='user_icon'>
                <Col xs={12} className='center' style={{ overflow: 'hidden' }}>
                  <img src={sellerInfo.image_url} className='user_img' />
                </Col>
                <Col xs={12} className='center'>
                  <h4>{sellerInfo.name}</h4>
                </Col>
              </Row>

              <SellerButtons />
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

export default Seller;
