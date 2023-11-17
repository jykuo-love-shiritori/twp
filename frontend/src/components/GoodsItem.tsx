import '@components/style.css';
import '@style/global.css';

import TButton from '@components/TButton';

interface Props {
  id: number;
  name: string;
  imgUrl: string;
  isIndex: boolean;
}

const GoodsItem = ({ id, name, imgUrl, isIndex }: Props) => {
  return (
    <div className='goods_item'>
      <img src={imgUrl} style={{ borderRadius: '0 0 30px 0', width: '100%' }} />
      <div style={{ padding: '2% 7% 2% 7% ' }}>
        <p>{name.substring(0, 13)} ...</p>
      </div>

      <TButton text='more' url={isIndex ? `discover/${id}` : `${id}`} />
    </div>
  );
};

export default GoodsItem;
