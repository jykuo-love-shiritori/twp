import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPen } from '@fortawesome/free-solid-svg-icons';
import { CSSProperties } from 'react';

import '@components/style.css';
import '@style/global.css';

import { Link } from 'react-router-dom';
import { Col, Row } from 'react-bootstrap';

interface Props {
  id: number;
  name: string;
  image_url: string;
  price: number;
  sales: number;
}

const GoodsItemStyle: CSSProperties = {
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

const SellerGoodsItem = ({ id, name, image_url, price, sales }: Props) => {
  return (
    <div style={GoodsItemStyle}>
      <div style={{ overflow: ' hidden' }}>
        <img src={image_url} style={GoodsImgStyle} />
      </div>

      <Row style={{ padding: '7% 8% 0 8%' }}>
        <Col xs={12} md={12} style={{ padding: '0' }}>
          <h5 className='crop_text'>{name}</h5>
        </Col>
        <Col xs={12} md={6} style={{ color: 'var(--title)', padding: '0' }}>
          <div className='crop_text'>${price}</div>
        </Col>
        <Col xs={12} md={6} className='right' style={{ padding: '0' }}>
          <div className='crop_text'>{`sold ${sales}`}</div>
        </Col>
      </Row>

      <Link to={`/user/seller/manageProducts/${id}`}>
        <div className='button pointer center'>
          <FontAwesomeIcon icon={faPen} className='white_word' />
        </div>
      </Link>
    </div>
  );
};

export default SellerGoodsItem;
