import { Row, Col } from 'react-bootstrap';

interface CouponItemTemplateProps {
  data: {
    id: number | null;
    name: string;
    policy: string;
    date: string;
    introduction: string;
    tags: { name: string }[];
  };
}

const couponStyle = {
  backgroundColor: 'var(--button_dark)',
  boxShadow: '3px 5px 10px 0px rgba(0, 0, 0, 0.25)',
  borderRadius: '30px',
  padding: '5%',
  border: 'var(--border) solid 2px',
};

const CouponItemTemplate = ({ data }: CouponItemTemplateProps) => {
  return (
    <div style={{ ...couponStyle }}>
      <Row style={{ height: '100%', padding: '2% 0' }}>
        <Col xs={9} md={9} xl={8} className='center'>
          <div>
            <div className='center' style={{ fontSize: '20px', fontWeight: '700', color: 'white' }}>
              {data.name}
            </div>
            <div className='center' style={{ fontSize: '16px', fontWeight: '500', color: 'white' }}>
              {data.policy}
            </div>
          </div>
        </Col>
        <Col xs={3} md={3} xl={4} className='center' style={{ borderLeft: '2px dashed #AAAAAA' }}>
          <div>
            <div className='center' style={{ fontSize: '20px', fontWeight: '500', color: 'white' }}>
              exp
            </div>
            <div className='center' style={{ fontSize: '16px', fontWeight: '500', color: 'white' }}>
              {data.date}
            </div>
          </div>
        </Col>
      </Row>
    </div>
  );
};

export default CouponItemTemplate;
