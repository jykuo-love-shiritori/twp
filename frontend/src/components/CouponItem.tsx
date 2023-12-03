import { Col, Row } from 'react-bootstrap';
import { Link } from 'react-router-dom';

interface CouponItemProps {
  data: {
    id: number;
    name: string;
    policy: string;
    date: string;
  };
}

const couponStyle = {
  backgroundColor: 'var(--button_dark)',
  boxShadow: '3px 5px 10px 0px rgba(0, 0, 0, 0.25)',
  borderRadius: '30px',
  padding: '5%',
};

const CouponItem = ({ data }: CouponItemProps) => {
  return (
    <Link className='none' to={`${window.location.pathname}/${data.id}`}>
      <div style={couponStyle}>
        <Row style={{ height: '4rem' }}>
          <Col xs={8} className='center'>
            <div>
              <div
                className='center'
                style={{ fontSize: '16px', fontWeight: '700', color: 'white' }}
              >
                {data.name}
              </div>
              <div
                className='center'
                style={{ fontSize: '12px', fontWeight: '500', color: 'white' }}
              >
                {data.policy}
              </div>
            </div>
          </Col>
          <Col xs={4} className='center' style={{ borderLeft: '2px dashed #AAAAAA' }}>
            <div>
              <div
                className='center'
                style={{ fontSize: '16px', fontWeight: '500', color: 'white' }}
              >
                exp
              </div>
              <div
                className='center'
                style={{ fontSize: '12px', fontWeight: '500', color: 'white' }}
              >
                {data.date}
              </div>
            </div>
          </Col>
        </Row>
      </div>
    </Link>
  );
};

export default CouponItem;
