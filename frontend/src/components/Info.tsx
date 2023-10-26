import TButton from './TButton';
import InfoItem from './InfoItem';

const Info = () => {
  return (
    <div>
      <div className='title'>Personal info</div>
      <hr className='hr' />
      <InfoItem text='Name' isMore={false} />
      <InfoItem text='Email Address' isMore={false} />
      <InfoItem text='Address' isMore={true} />
      <TButton text='Save' url='' />
    </div>
  );
};

export default Info;
