import CouponItem from '@components/CouponItem';
import { Col, Row } from 'react-bootstrap';
import couponData from '@pages/coupon/couponData.json';
import TButton from '@components/TButton';

const ManageSellerCoupons = () => {
  return (
    <div>
      <Row>
        {/* display title for desktop */}
        <div className='disappear_phone disappear_tablet'>
          <Row>
            <Col xl={6} className='left'>
              <div className='title'>All Global Coupon</div>
            </Col>
            <Col xl={3} />
            <Col xl={3} className='right'>
              <TButton text='New Coupon' url='new' />
            </Col>
          </Row>
        </div>
        {/* display title for tablet */}
        <div className='disappear_desktop disappear_phone'>
          <Row>
            <Col md={8} className='left'>
              <div className='title'>All Global Coupon</div>
            </Col>
            <Col md={4} className='right'>
              <TButton text='New Coupon' url='new' />
            </Col>
          </Row>
        </div>
        {/* display title for phone */}
        <div className='disappear_desktop disappear_tablet'>
          <Row>
            <Col xs={12} className='center' style={{ padding: '0' }}>
              <div className='title'>All Global Coupon</div>
            </Col>
            <Col />
            <Col xs={6} className='center' style={{ padding: '0 0 2% 0' }}>
              <TButton text='New Coupon' url='new' />
            </Col>
            <Col />
          </Row>
        </div>

        <hr />
        <Row>
          {couponData.map((data, index) => {
            return (
              <Col xs={6} md={4} key={index} style={{ padding: '2%' }}>
                <CouponItem data={data} />
              </Col>
            );
          })}
        </Row>
      </Row>
    </div>
  );
};

export default ManageSellerCoupons;
