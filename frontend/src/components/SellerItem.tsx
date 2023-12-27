import '@style/global.css';

import TButton from '@components/TButton';

export interface SellerItemProps {
  seller_name: string;
  name: string;
  image_url: string;
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

  if (!data) {
    return <div style={SellerItemStyle}>no user</div>;
  }

  return (
    <div style={SellerItemStyle}>
      <div className='center'>
        <div style={userImgStyle} />
      </div>
      <div className='center' style={{ paddingTop: '10px' }}>
        <h5>{data.name}</h5>
      </div>

      {/* url is empty because each user's user shop page is different and defining data is tough 
            i want to wait for getting data from backend then deal with it */}
      <TButton text='View Shop' />
    </div>
  );
};

export default SellerItem;
