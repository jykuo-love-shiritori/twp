import { Col, Row } from 'react-bootstrap';
import { Outlet } from 'react-router-dom';

import BuyerButtons from '@components/BuyerButtons';

import userData from '@pages/user/buyer/buyerInfo.json';

const User = () => {
  return (
    <Row>
      <Col xs={12} md={12}>
        <div className='user_bg center' />
      </Col>
      <Col xs={12} md={3}>
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
      <Col xs={12} md={9} style={{ padding: '1% 7% 6% 7%' }}>
        {/* the personal info, security and order history */}
        <Outlet />
      </Col>
    </Row>
  );
};

export default User;
