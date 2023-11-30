import { Col, Row } from 'react-bootstrap';

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
    <div style={couponStyle}>
      <Row>
        <Col className='center center_vertical' md={7}>
          <div>
            <div>{data.name}</div>
            <div>{data.policy}</div>
          </div>
        </Col>
        <Col className='center center_vertical' md={5}>
          <div>
            <div>{data.name}</div>
            <div>{data.policy}</div>
          </div>
        </Col>
      </Row>
    </div>
  );
};

export default CouponItem;
