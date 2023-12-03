import { Col, Row } from 'react-bootstrap';
import CouponItem from '@components/CouponItem';
import couponData from '@pages/coupon/couponData.json';

const ManageAdminCoupons = () => {
  return (
    <div>
      <Row>
        <Col md={12}>
          <div className='title'>All Global Coupon</div>
        </Col>
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

export default ManageAdminCoupons;
