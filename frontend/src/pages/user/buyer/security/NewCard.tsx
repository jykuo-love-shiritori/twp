import TButton from '@components/TButton';
import PasswordItem from '@components/PasswordItem';
import InfoItem from '@components/InfoItem';

const NewCard = () => {
  return (
    <div>
      <div className='title'>Security - Credit Card</div>
      <hr className='hr' />
      <InfoItem text='CreditCard Number' isMore={false} />
      <PasswordItem text='Security Code (CVV)' />
      <InfoItem text='Expiration Date' isMore={false} />
      <InfoItem text='Cardholder Name' isMore={false} />
      <TButton text='Save' url='' />
    </div>
  );
};

export default NewCard;
