import '@components/style.css';
import '@style/global.css';

import TButton from '@components/TButton';

interface Props {
  id: number;
  name: string;
  imgUrl: string;
}

const GoodsItem = ({ id, name, imgUrl }: Props) => {
  return (
    <div className='goods_item'>
      <img src={imgUrl} style={{ borderRadius: '0 0 30px 0', width: '100%' }} />
      <div style={{ padding: '2% 7% 2% 7% ' }}>
        <p>
          {name.substring(0, 11)} {name.length > 13 ? '...' : ''}
        </p>
      </div>

      <TButton text='more' url={`/sellerID/shop/${id}`} />
    </div>
  );
};

export default GoodsItem;
