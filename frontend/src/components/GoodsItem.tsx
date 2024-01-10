import '@components/style.css';
import '@style/global.css';

import { CSSProperties } from 'react';

import TButton from '@components/TButton';
import { Col, Row } from 'react-bootstrap';

export interface Props {
  id: number;
  name: string;
  image_url: string;
  price: number;
  sales: number;
}

const GoodsItemStyle = {
  boxShadow: '3px 5px 10px 0px rgba(0, 0, 0, 0.25)',
  padding: '9% 8% 9% 8%',
  margin: '15px 0 15px 0',
  borderRadius: '10px',
  border: '1px solid var(--button_border, #34977f)',
  background: 'var(--button_dark, #135142)',
};

const GoodsImgStyle: CSSProperties = {
  borderRadius: '0 0 30px 0',
  width: '100%',
  minHeight: '20vh',
  maxHeight: '20vh',
  objectFit: 'cover',
};

export interface GoodsItemProps {
  description: string;
  id: number;
  image_url: string;
  name: string;
  price: number;
  sales: number;
}

const GoodsItem = ({ id, name, image_url, price, sales }: Props) => {
  return (
    <div style={GoodsItemStyle}>
      <div style={{ overflow: ' hidden' }}>
        <img src={image_url} style={GoodsImgStyle} />
      </div>

      <Row style={{ padding: '7% 8% 0 8%' }}>
        <Col xs={12} md={8} style={{ padding: '0' }}>
          <h5 className='crop_text'>{name}</h5>
        </Col>
        <Col xs={12} md={4} className='right' style={{ padding: '0' }}>
          <h5 className='crop_text'>{price}</h5>
        </Col>
        <Col xs={12} className='right' style={{ color: 'rgb(133,133,133)', padding: '0' }}>
          {`sold ${sales}`}
        </Col>
      </Row>

      <TButton text='more' action={`/product/${id}`} />
    </div>
  );
};

export default GoodsItem;
