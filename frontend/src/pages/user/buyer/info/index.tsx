import { useState } from 'react';

import TButton from '@components/TButton';
import InfoItem from '@components/InfoItem';

const Info = () => {
  const [name, setName] = useState<string>('');
  const [email, setEmail] = useState<string>('');
  const [address, setAddress] = useState<string>('');

  return (
    <div>
      <div className='title'>Personal info</div>
      <hr className='hr' />
      <InfoItem text='Name' isMore={false} value={name} setValue={setName} />
      <InfoItem text='Email Address' isMore={false} value={email} setValue={setEmail} />
      <InfoItem text='Address' isMore={true} value={address} setValue={setAddress} />
      <TButton text='Save' url='' />
    </div>
  );
};

export default Info;
