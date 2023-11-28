import { useState } from 'react';

import TButton from '@components/TButton';
import PasswordItem from '@components/PasswordItem';

const Password = () => {
  const [password, setPassword] = useState<string>('');
  const [confirmedPassword, setConfirmedPassword] = useState<string>('');
  const [oldPassword, setOldPassword] = useState<string>('');

  return (
    <div>
      <div className='title'>Security - Password</div>
      <hr className='hr' />
      <PasswordItem text='Password' value={password} setValue={setPassword} />
      <PasswordItem
        text='ConfirmedPassword'
        value={confirmedPassword}
        setValue={setConfirmedPassword}
      />
      <PasswordItem text='Old Password' value={oldPassword} setValue={setOldPassword} />
      <TButton text='Save' url='' />
    </div>
  );
};

export default Password;
