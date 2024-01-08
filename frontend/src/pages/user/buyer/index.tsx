import { Col, Row } from 'react-bootstrap';
import { Outlet, useNavigate } from 'react-router-dom';
import { RouteOnNotOK } from '@lib/Status';
import { useQuery } from '@tanstack/react-query';
import { useAuth } from '@lib/Auth';
import NotFound from '@components/NotFound';
import BuyerButtons from '@components/BuyerButtons';
import { IBuyerInfo } from '@pages/user/buyer/info';

const User = () => {
  const token = useAuth();
  const navigate = useNavigate();

  const { data: buyerInfo } = useQuery({
    queryKey: ['userGetInfo'],
    queryFn: async () => {
      const resp = await fetch('/api/user/info', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });
      if (!resp.ok) {
        RouteOnNotOK(resp, navigate);
      }
      return resp.json();
    },
    select: (data) => data as IBuyerInfo,
  });

  if (!buyerInfo) {
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
                  <img src={buyerInfo.image_url} className='user_img' />
                </Col>
                <Col xs={12} className='center'>
                  <h4>{buyerInfo.name}</h4>
                </Col>
              </Row>

            <BuyerButtons />
          </Col>
          <Col xs={12} md={9} ld={10} style={{ padding: '1% 5% 0px 5%' }}>
            <Outlet />
          </Col>
        </Row>
      </Col>
    </Row>
  );
};

export default User;
