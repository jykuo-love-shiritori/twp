import { Col, Row } from 'react-bootstrap';

interface Props {
  data: ProductProps;
}

interface ProductProps {
  enabled?: boolean;
  image_id: string;
  name: string;
  price: number;
  product_id?: number;
  quantity: number;
  stock?: number;
}

const HistoryProduct = ({ data }: Props) => {
  return (
    <div className='cart_item ' style={{ margin: '2% 2%' }}>
      {/* layout on tablet, desktop */}
      <div className='disappear_phone ' style={{ fontSize: '24px' }}>
        <Row className='center_vertical dark' style={{ padding: '0' }}>
          <Col md={1} className='center'>
            <img src={data.image_id} style={{ width: 'max(40px, 90%)', borderRadius: '10px' }} />
          </Col>
          <Col md={4} style={{ padding: '0 2%' }}>
            {data.name}
          </Col>
          <Col md={3} className='center'>
            x{data.quantity}
          </Col>
          <Col md={3} className='right'>
            {data.price * data.quantity} NTD
          </Col>
        </Row>
      </div>

      {/* layout on phone */}
      <div className='disappear_tablet disappear_desktop' style={{ fontSize: '14px' }}>
        <Row className='center_vertical dark'>
          <Col xs={4} className='center'>
            <img src={data.image_id} style={{ width: '90%', borderRadius: '10px' }} />
          </Col>
          <Col xs={8} style={{ padding: '0 2%' }}>
            <Row>
              <Col xs={12}>{data.name}</Col>
              <Col xs={4}>x{data.quantity}</Col>
              <Col xs={7} className='right'>
                {data.price * data.quantity} NTD
              </Col>
            </Row>
          </Col>
        </Row>
      </div>
    </div>
  );
};

export default HistoryProduct;
