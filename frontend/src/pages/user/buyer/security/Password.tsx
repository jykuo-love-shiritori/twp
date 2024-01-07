import { useForm } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';
import { useState } from 'react';
import { RouteOnNotOK } from '@lib/Status';
import { useAuth } from '@lib/Auth';
import WarningModal from '@components/WarningModal';
import FormItem from '@components/FormItem';

interface IEditPassword {
  current_password: string;
  new_password: string;
  confirm_password: string;
}

const Password = () => {
  const navigate = useNavigate();
  const token = useAuth();
  const [show, setShow] = useState<boolean>(false);
  const [warningText, setWarningText] = useState<string>('');
  const { register, handleSubmit } = useForm<IEditPassword>();
  const onSubmit = async (data: IEditPassword) => {
    if (
      !data.new_password.match(
        /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&[\],.(){}":;'+\-=_~])[A-Za-z\d@$!%*?&[\],.(){}":;'+\-=_~]{8,72}$/,
      )
    ) {
      setWarningText(
        'Password should contain at least one of each: uppercase letter, lowercase letter, number and special character. And the length should be between 8 and 72.',
      );
      setShow(true);
      return;
    }
    if (data.new_password !== data.confirm_password) {
      setWarningText('New password and confirm password do not match.');
      setShow(true);
      return;
    }
    if (data.new_password === data.current_password) {
      setWarningText('New password and old password are the same.');
      setShow(true);
      return;
    }
    const resp = await fetch('/api/user/security/password', {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        current_password: data.current_password,
        new_password: data.new_password,
      }),
    });
    if (!resp.ok) {
      RouteOnNotOK(resp, navigate);
      if (resp.status === 401) {
        const response = await resp.json();
        alert(response.message);
      }
    } else {
      navigate('/user/security');
    }
  };

  return (
    <>
      <form onSubmit={handleSubmit(onSubmit)}>
        <div className='title'>Security - Password</div>
        <hr className='hr' />
        <FormItem label='Old Password'>
          <input
            type='password'
            placeholder='Old Password'
            {...register('current_password', { required: true })}
          />
        </FormItem>
        <FormItem label='New Password'>
          <input
            type='password'
            placeholder='New Password'
            {...register('new_password', { required: true })}
          />
        </FormItem>
        <FormItem label='Confirm Password'>
          <input
            type='password'
            placeholder='Confirm Password'
            {...register('confirm_password', { required: true })}
          />
        </FormItem>
        <div className='form_item_wrapper'>
          <input type='submit' value='Save' />
        </div>
      </form>
      <WarningModal show={show} text={warningText} onHide={() => setShow(false)} />
    </>
  );
};

export default Password;
