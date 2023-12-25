// import CouponItem from '@components/CouponItem';
import { Col, Row } from 'react-bootstrap';
import couponData from '@pages/coupon/couponData.json';

// interface CouponProps {
//   id: number;
//   type: 'percentage' | 'fixed' | 'shipping';
//   name: string;
//   description: string;
//   discount: number;
//   start_date: string;
//   expire_date: string;
//   tags: {
//     name: string;
//   }[];
// }

const SellerCoupons = () => {
  return (
    <div>
      <Row>
        <Col md={12}>
          <div className='title'>All Coupon</div>
        </Col>
        <hr />
        <Row>
          <div className='disappear_phone'>
            <Row>
              {couponData.map((data, index) => {
                return (
                  <Col md={4} xl={3} key={index} style={{ padding: '2%' }}>
                    {/* TODO */}
                    {/* <CouponItem data={data as CouponProps} /> */}
                    {data.id}
                  </Col>
                );
              })}
            </Row>
          </div>
          <div className='disappear_desktop disappear_tablet'>
            <Row>
              {couponData.map((data, index) => {
                return (
                  <Col xs={12} key={index} style={{ padding: '2% 10%' }}>
                    {/* TODO */}
                    {/* <CouponItem data={data as CouponProps} /> */}
                    {data.id}
                  </Col>
                );
              })}
            </Row>
          </div>
        </Row>
      </Row>
    </div>
  );
};

export default SellerCoupons;
