import TButton from '@components/TButton';
import PasswordItem from '@components/PasswordItem';

const Password = () => {
  return (
    <div>
      <div className='title'>Security - Password</div>
      <hr className='hr' />
      <PasswordItem text='Password' />
      <PasswordItem text='ConfirmedPassword' />
      <PasswordItem text='Old Password' />
      <TButton text='Save' url='' />
    </div>
  );
};

export default Password;
