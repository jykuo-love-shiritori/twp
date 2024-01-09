import '@components/style.css';
import '@style/global.css';

import { CSSProperties } from 'react';

import TButton from '@components/TButton';

export interface Props {
  id: number;
  name: string;
  image_url: string;
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

const GoodsItem = ({ id, name, image_url }: Props) => {
  return (
    <div style={GoodsItemStyle}>
      <div style={{ overflow: ' hidden' }}>
        <img src={image_url} style={GoodsImgStyle} />
      </div>

      <div style={{ padding: '2% 7% 2% 7% ' }}>
        <p>
          {name.substring(0, 11)} {name.length > 13 ? '...' : ''}
        </p>
      </div>

      <TButton text='more' action={`/product/${id}`} />
    </div>
  );
};

export default GoodsItem;
