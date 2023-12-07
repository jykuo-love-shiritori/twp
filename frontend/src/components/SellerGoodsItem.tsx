import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPen } from '@fortawesome/free-solid-svg-icons';

import '@components/style.css';
import '@style/global.css';

import { Link } from 'react-router-dom';

interface Props {
  id: number;
  name: string;
  imgUrl: string;
  isIndex: boolean;
}

const SellerGoodsItem = ({ id, name, imgUrl, isIndex }: Props) => {
  const GoodsItemStyle = {
    boxShadow: '3px 5px 10px 0px rgba(0, 0, 0, 0.25)',
    padding: '9% 8% 9% 8%',
    margin: '15px 0 15px 0',
    borderRadius: '10px',
    border: '1px solid var(--button_border, #34977f)',
    background: 'var(--button_dark, #135142)',
  };

  return (
    <div style={GoodsItemStyle}>
      <img src={imgUrl} style={{ borderRadius: '0 0 30px 0', width: '100%' }} />
      <div style={{ padding: '2% 7% 2% 7% ' }}>
        <p>
          {name.substring(0, 11)} {name.length > 13 ? '...' : ''}
        </p>
      </div>

      <Link to={isIndex ? `discover/${id}` : `${id}`}>
        <div className='button pointer center'>
          <FontAwesomeIcon icon={faPen} className='white_word' />
        </div>
      </Link>
    </div>
  );
};

export default SellerGoodsItem;
