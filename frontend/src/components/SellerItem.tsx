import '@style/global.css';

import TButton from '@components/TButton';
import sellersData from '@pages/user/seller/sellersInfo.json';

interface Input {
  id: number;
}

interface Props extends Input {
  name: string;
  imgUrl: string;
}

const SellerItem = ({ id }: Input) => {
  const SellerItemStyle = {
    borderRadius: '10px',
    padding: '10% 10% 5% 10%',
    marginBottom: '30px',
    border: '1px solid var(--button_border, #34977F)',
    background: 'var(--button_dark, #135142)',
  };

  let data: Props | undefined;
  sellersData.findIndex((item) => {
    if (item.id === id) {
      data = item;
    }
  });

  if (data) {
    const userImgStyle = {
      borderRadius: '50%',
      background: `url(${data.imgUrl}) lightgray 50% / cover no-repeat`,
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

        {/* url is empty because each user's user shop page is different and defining data is tough 
            i want to wait for getting data from backend then deal with it */}
        <TButton text='View Shop' url='' />
      </div>
    );
  }
};

export default SellerItem;
