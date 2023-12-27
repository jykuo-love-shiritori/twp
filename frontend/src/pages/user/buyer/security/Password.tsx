import { useForm } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';
import { RouteOnNotOK } from '@lib/Status';
import FormItem from '@components/FormItem';

interface IEditPassword {
  current_password: string;
  new_password: string;
  confirm_password: string;
}

const Password = () => {
  const navigate = useNavigate();
  const { register, handleSubmit } = useForm<IEditPassword>();
  const onSubmit = async (data: IEditPassword) => {
    if (
      !data.new_password.match(
        /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,72}$/,
      )
    ) {
      alert(
        'password should contain at least one of each: uppercase letter, lowercase letter, number and special character',
      );
      return;
    }
    if (data.new_password !== data.confirm_password) {
      alert('passwords do not match');
      return;
    }
    if (data.new_password === data.current_password) {
      alert('new password cannot be the same as old password');
      return;
    }
    const resp = await fetch('/api/user/security/password', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        current_password: data.current_password,
        new_password: data.new_password,
      }),
    });
    if (!resp.ok) {
      RouteOnNotOK(resp, navigate);
      // TODO: change when another pr is merged
      const a = await resp.json();
      alert(a.message);
    } else {
      navigate('/user/security');
    }
  };

  return (
    <div>
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
    </div>
  );
};

export default Password;
