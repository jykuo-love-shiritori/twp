import { useForm } from 'react-hook-form';
import FormItem from '@components/FormItem';
import { RouteOnNotOK } from '@lib/Status';
import { useNavigate } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';
import { CheckFetchStatus } from '@lib/Status';

interface ICreditCard {
  CVV: string;
  name: string;
  card_number: string;
  expiry_date: string;
}

const NewCard = () => {
  const navigate = useNavigate();
  const { register, handleSubmit } = useForm<ICreditCard>();
  const onSubmit = async (data: ICreditCard) => {
    const resp = await fetch('/api/user/security/credit_card', {
      method: 'PATCH',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(oldCards ? oldCards.concat(data) : [data]),
    });
    if (!resp.ok) {
      RouteOnNotOK(resp, navigate);
    } else {
      navigate('/user/security/manageCreditCard');
    }
  };

  const { data: oldCards, status } = useQuery({
    queryKey: ['userGetCreditCard'],
    queryFn: async () => {
      const resp = await fetch('/api/user/security/credit_card');
      RouteOnNotOK(resp, navigate);
      return await resp.json();
    },
    select: (data) => data as [ICreditCard],
    retry: false,
    enabled: true,
    refetchOnWindowFocus: false,
  });
  if (status !== 'success') {
    return <CheckFetchStatus status={status} />;
  }

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <div className='title'>Security - Credit Card</div>
      <hr className='hr' />
      <FormItem label='Card Name'>
        <input type='text' placeholder='Card Name' {...register('name', { required: true })} />
      </FormItem>
      <FormItem label='Card Number'>
        <input
          type='text'
          placeholder='Card Number'
          {...register('card_number', {
            required: true,
            pattern: {
              value: /^\d{4} \d{4} \d{4} \d{4}$/,
              message: 'must be like: "XXXX XXXX XXXX XXXX"',
            },
          })}
        />
      </FormItem>
      <FormItem label='Security Code (CVV)'>
        <input
          type='text'
          placeholder='Security Code (CVV)'
          {...register('CVV', {
            required: true,
            pattern: { value: /^\d{3}$/, message: 'must be 3 digits' },
          })}
        />
      </FormItem>
      <FormItem label='Expiration Date'>
        <input
          type='text'
          placeholder='Expiration Date'
          {...register('expiry_date', {
            required: true,
            pattern: {
              value: /^\d{2}\/\d{2}$/,
              message: 'must be like: "MM/YY"',
            },
          })}
        />
      </FormItem>
      <div className='form_item_wrapper'>
        <input type='submit' value='Save' />
      </div>
    </form>
  );
};

export default NewCard;
