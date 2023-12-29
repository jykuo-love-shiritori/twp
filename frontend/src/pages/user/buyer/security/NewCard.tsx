import { useForm } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';
import { useState } from 'react';
import { useQuery } from '@tanstack/react-query';
import { RouteOnNotOK } from '@lib/Status';
import { CheckFetchStatus } from '@lib/Status';
import FormItem from '@components/FormItem';
import WarningModal from '@components/WarningModal';

interface ICreditCard {
  CVV: string;
  name: string;
  card_number: string;
  expiry_date: string;
}

const NewCard = () => {
  const navigate = useNavigate();
  const [show, setShow] = useState<boolean>(false);
  const [warningText, setWarningText] = useState<string>('');
  const { register, handleSubmit } = useForm<ICreditCard>();
  const onSubmit = async (data: ICreditCard) => {
    if (!data.card_number.match(/^\d{4} \d{4} \d{4} \d{4}$/)) {
      setWarningText('Card number must be like: "XXXX XXXX XXXX XXXX".');
      setShow(true);
      return;
    }
    if (!data.CVV.match(/^\d{3}$/)) {
      setWarningText('Security code (CVV) must be 3 digits.');
      setShow(true);
      return;
    }
    if (!data.expiry_date.match(/^\d{2}\/\d{2}$/)) {
      setWarningText('Expiration date must be like: "MM/YY".');
      setShow(true);
      return;
    }

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
    <>
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
            {...register('card_number', { required: true })}
          />
        </FormItem>
        <FormItem label='Security Code (CVV)'>
          <input
            type='text'
            placeholder='Security Code (CVV)'
            {...register('CVV', { required: true })}
          />
        </FormItem>
        <FormItem label='Expiration Date'>
          <input
            type='text'
            placeholder='Expiration Date'
            {...register('expiry_date', { required: true })}
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

export default NewCard;
