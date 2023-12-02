import { Col, Row } from 'react-bootstrap';
import { Outlet } from 'react-router-dom';

import TButton from '@components/TButton';

import userData from '@pages/user/seller/sellerInfo.json';
import goodsData from '@pages/discover/goodsData.json';

const UserViewShop = () => {
  return (
    <Row style={{ width: '100%', padding: '0', margin: '0' }}>
      <Col xs={12} md={12} style={{ width: '100%', padding: '0' }}>
        <div className='user_bg center'>
          <div style={{ padding: '6% 10% 6% 10%' }}>{userData.introduction}</div>
        </div>
      </Col>
      <Col xs={12} md={3} lg={2}>
        <Row className='user_icon' style={{ padding: '0 7% 0 7%' }}>
          <Col xs={12} className='center'>
            <img src={userData.imgUrl} className='user_img' />
          </Col>
          <Col xs={12}>
            <div className='center'>
              <h4 className='title_color' style={{ padding: '10% 2% 0% 2%' }}>
                <b>{userData.name}</b>
              </h4>
            </div>
            <hr className='hr' />
            <div className='center'> Products : {goodsData.length} items</div>
            <TButton text='Explore Shop' url='/sellerID/shop' />
            <TButton text='Check Coupons' url='/sellerID/coupons' />
          </Col>
        </Row>
      </Col>
      <Col xs={12} md={9} ld={10} style={{ padding: '1% 5% 6% 5%' }}>
        <Outlet />
      </Col>
    </Row>
  );
};

export default UserViewShop;
