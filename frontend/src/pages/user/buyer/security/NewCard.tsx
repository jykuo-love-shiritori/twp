import { useState } from 'react';

import TButton from '@components/TButton';
import PasswordItem from '@components/PasswordItem';
import InfoItem from '@components/InfoItem';

const NewCard = () => {
  const [creditCardNumber, setCreditCardNumber] = useState<string>('');
  const [CVV, setCVV] = useState<string>('');
  const [expirationDate, setExpirationDate] = useState<string>('');
  const [cardHolderName, setCardHolderName] = useState<string>('');

  return (
    <div>
      <div className='title'>Security - Credit Card</div>
      <hr className='hr' />
      <InfoItem
        text='CreditCard Number'
        isMore={false}
        value={creditCardNumber}
        setValue={setCreditCardNumber}
      />
      <PasswordItem text='Security Code (CVV)' value={CVV} setValue={setCVV} />
      <InfoItem
        text='Expiration Date'
        isMore={false}
        value={expirationDate}
        setValue={setExpirationDate}
      />
      <InfoItem
        text='Cardholder Name'
        isMore={false}
        value={cardHolderName}
        setValue={setCardHolderName}
      />
      <TButton text='Save' url='' />
    </div>
  );
};

export default NewCard;
