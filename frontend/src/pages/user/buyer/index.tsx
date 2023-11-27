import { Col, Row } from 'react-bootstrap';
import { Outlet } from 'react-router-dom';

import BuyerButtons from '@components/BuyerButtons';

import userData from '@pages/user/buyer/buyerInfo.json';

const User = () => {
  return (
    <Row style={{ width: '100%', padding: '0', margin: '0' }}>
      <Col xs={12} md={12} style={{ width: '100%', padding: '0' }}>
        <div className='user_bg center' />
      </Col>
      <Col xs={12} md={3} lg={2} style={{ padding: '0' }}>
        <Row className='user_icon'>
          <Col xs={12} className='center'>
            <img src={userData.imgUrl} className='user_img' />
          </Col>
          <Col xs={12} className='center'>
            <h4>{userData.name}</h4>
          </Col>
        </Row>

        <BuyerButtons />
      </Col>
      <Col xs={12} md={10} ld={10} style={{ padding: '1% 7% 6% 7%' }}>
        <Outlet />
      </Col>
    </Row>
  );
};

export default User;
