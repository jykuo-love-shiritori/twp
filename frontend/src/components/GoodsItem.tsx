import '@components/style.css';
import '@style/global.css';

import TButton from '@components/TButton';

export interface Props {
  id: number;
  name: string;
  image_url: string;
}

const GoodsItem = ({ id, name, image_url }: Props) => {
  return (
    <div className='goods_item'>
      <img src={image_url} style={{ borderRadius: '0 0 30px 0', width: '100%' }} />
      <div style={{ padding: '2% 7% 2% 7% ' }}>
        <p>
          {name.substring(0, 11)} {name.length > 13 ? '...' : ''}
        </p>
      </div>

      <TButton text='more' action={`/sellerID/shop/${id}`} />
    </div>
  );
};

export default GoodsItem;
