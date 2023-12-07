import { Col, Row } from 'react-bootstrap';
import AllCouponData from './AllCouponData.json';
import CouponItem from '@components/CouponItem';

const Coupons = () => {
  return (
    <div style={{ padding: '6% 8% 0 8%' }}>
      {AllCouponData.map((data, index) => {
        return (
          <Row key={index} style={{ paddingBottom: '4%' }}>
            <Col className='title' style={{ padding: '0 0 1% 1%' }}>
              {data.owner}
            </Col>
            <hr className='hr' />
            <div className='disappear_phone'>
              <Row>
                {data.coupons.map((data, index) => {
                  return (
                    <Col md={4} xl={3} key={index} style={{ padding: '2%' }}>
                      <CouponItem data={data} />
                    </Col>
                  );
                })}
              </Row>
            </div>
            <div className='disappear_desktop disappear_tablet'>
              <Row>
                {data.coupons.map((data, index) => {
                  return (
                    <Col xs={12} key={index} style={{ padding: '2% 10%' }}>
                      <CouponItem data={data} />
                    </Col>
                  );
                })}
              </Row>
            </div>
            <Col xs={12} className='right'>
              <div className='more_button'>{'more >'}</div>
            </Col>
          </Row>
        );
      })}
    </div>
  );
};

export default Coupons;
