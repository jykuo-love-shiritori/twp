import CouponItem from '@components/CouponItem';
import { Col, Row } from 'react-bootstrap';

const couponsData = [
  { id: 0, name: 'Coupon 1', policy: 'Save 20%', date: '10/10' },
  { id: 1, name: 'Coupon 2', policy: 'Save 20%', date: '10/10' },
  { id: 2, name: 'Coupon 3', policy: 'Save 20%', date: '10/10' },
  { id: 3, name: 'Coupon 4', policy: 'Save 20%', date: '10/10' },
];

const tableStyle = {};

const ManageSellerCoupons = () => {
  return (
    <div>
      <Row>
        <Col md={12}>
          <div className='title'>All Coupon</div>
        </Col>
        <hr />
        <Row style={tableStyle}>
          {couponsData.map((data, index) => {
            return (
              <Col xs={6} md={4} key={index} style={{ padding: '10px' }}>
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
