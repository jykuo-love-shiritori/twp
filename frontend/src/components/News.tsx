import 'bootstrap/dist/css/bootstrap.min.css';
import '@style/global.css';

import { CSSProperties } from 'react';

import TButton from '@components/TButton';

export interface NewsProps {
  id: number;
  image_id: string;
  title: string;
}

const News = ({ id, image_id, title }: NewsProps) => {
  const NewsComponentStyle: CSSProperties = {
    borderRadius: '52px',
    boxShadow: '6px 6px 15px 5px rgba(0, 0, 0, 0.15)',
    marginBottom: '20px',
    width: '100%',
    border: 'var(--border) solid 1px',
    height: '250px',
    objectFit: 'cover',
  };

  return (
    <div>
      <div style={{ overflow: ' hidden' }}>
        <img src={image_id} style={NewsComponentStyle} />
      </div>

      <div style={{ padding: '1% 10% 1% 10%' }} className='center'>
        <span>
          {title.substring(0, 25)} {title.length > 25 ? '...' : ''}
        </span>
      </div>

      <TButton text='more' action={`/news/${id}`} />
      <br />
    </div>
  );
};

export default News;
