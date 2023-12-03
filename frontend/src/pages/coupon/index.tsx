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
            <hr style={{ border: '1px solid white', borderRadius: '5px', opacity: '.5' }} />
            {data.coupons.map((data, index) => {
              return (
                <Col xs={6} md={4} xl={3} key={index} style={{ padding: '2%' }}>
                  <CouponItem data={data} />
                </Col>
              );
            })}
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
