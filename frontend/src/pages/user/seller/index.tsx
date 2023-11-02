import { Col, Row } from 'react-bootstrap';
import { Outlet } from 'react-router-dom';

import SellerButtons from '@components/SellerButtons';

import sellerData from '@pages/user/seller/sellerInfo.json';

const Seller = () => {
  return (
    <Row>
      <Col xs={12} md={3}>
        <div className='user_bg center' />
        <Row className='user_icon'>
          <Col xs={12} className='center'>
            <img src={sellerData.imgUrl} className='user_img' />
          </Col>
          <Col xs={12} className='center'>
            <h4>{sellerData.name}</h4>
          </Col>
        </Row>

        <SellerButtons />
      </Col>

      <Col xs={12} md={9} style={{ padding: '1% 7% 6% 7%' }}>
        {/* the personal info, security and order history */}
        <Outlet />
      </Col>
    </Row>
  );
};

export default Seller;
