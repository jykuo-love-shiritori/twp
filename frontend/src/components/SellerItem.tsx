import '@style/global.css';

import TButton from '@components/TButton';

export interface SellerItemProps {
  name: string;
  image_url: string;
  seller_name: string;
}

interface Props {
  data: SellerItemProps;
}

const SellerItem = ({ data }: Props) => {
  const SellerItemStyle = {
    borderRadius: '10px',
    padding: '10% 10% 5% 10%',
    marginBottom: '30px',
    border: '1px solid var(--button_border, #34977F)',
    background: 'var(--button_dark, #135142)',
  };

  const userImgStyle = {
    borderRadius: '50%',
    background: `url(${data.image_url}) lightgray 50% / cover no-repeat`,
    width: '100px',
    height: '100px',
  };

  return (
    <div style={SellerItemStyle}>
      <div className='center'>
        <div style={userImgStyle} />
      </div>
      <div className='center' style={{ paddingTop: '10px' }}>
        <h5>{data.name}</h5>
      </div>

      <TButton text='View Shop' action={`/shop/${data.seller_name}`} />
    </div>
  );
};

export default SellerItem;
