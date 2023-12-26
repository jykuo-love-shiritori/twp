import { Row, Col } from 'react-bootstrap';

interface ICouponItemTemplate {
  name: string;
  type: 'percentage' | 'fixed' | 'shipping';
  discount: number;
  expire_date: string;
}

const couponStyle = {
  backgroundColor: 'var(--button_dark)',
  boxShadow: '3px 5px 10px 0px rgba(0, 0, 0, 0.25)',
  borderRadius: '30px',
  padding: '5%',
  border: 'var(--border) solid 2px',
  width: '100%',
};

const CouponItemTemplate = ({ data }: { data: ICouponItemTemplate }) => {
  const expireDate = new Date(data.expire_date);
  let policy = '';
  switch (data.type) {
    case 'shipping':
      policy = 'Free shipping';
      break;
    case 'percentage':
      policy = `${data.discount}% OFF`;
      break;
    case 'fixed':
      policy = `Save ${data.discount} NTD`;
      break;
  }
  return (
    <div style={{ ...couponStyle }}>
      <Row style={{ height: '100%', padding: '2% 0' }}>
        <Col xs={9} md={9} xl={8} className='center'>
          <div>
            <div className='center' style={{ fontSize: '20px', fontWeight: '700', color: 'white' }}>
              {data.name}
            </div>
            <div className='center' style={{ fontSize: '16px', fontWeight: '500', color: 'white' }}>
              {policy}
            </div>
          </div>
        </Col>
        <Col
          xs={3}
          md={3}
          xl={4}
          className='center'
          style={{ borderLeft: '2px dashed rgba(255, 255, 255, 0.67)' }}
        >
          <div>
            <div className='center' style={{ fontSize: '20px', fontWeight: '500', color: 'white' }}>
              exp
            </div>
            <div className='center' style={{ fontSize: '16px', fontWeight: '500', color: 'white' }}>
              {`${expireDate.getMonth() + 1} / ${expireDate.getDay()}`}
            </div>
          </div>
        </Col>
      </Row>
    </div>
  );
};

export default CouponItemTemplate;
