import { Col, Row } from 'react-bootstrap';
import { useForm, SubmitHandler } from 'react-hook-form';
import FormItem from '@components/FormItem';

interface BuyerInfoProps {
  name: string;
  email: string;
  address: string;
}

const Info = () => {
  //TODO: get the info from existing user
  const { register, handleSubmit } = useForm<BuyerInfoProps>({
    defaultValues: {
      name: 'name',
      email: 'email',
      address: 'address',
    },
  });
  const OnFormOutput: SubmitHandler<BuyerInfoProps> = (data) => {
    console.log(data);
    return data;
  };

  return (
    <div>
      <div className='title'>Personal info</div>
      <hr className='hr' />
      <form onSubmit={handleSubmit(OnFormOutput)}>
        <FormItem label='Name'>
          <input {...register('name', { required: true })} />
        </FormItem>
        <FormItem label='Email'>
          <input {...register('email', { required: true })} />
        </FormItem>
        <FormItem label='Address'>
          <input {...register('address', { required: true })} />
        </FormItem>

        <Row>
          <Col className='disappear_phone' />
          <Col xs={12} md={4} className='form_item_wrapper' style={{ paddingTop: '24px' }}>
            <input type='submit' value='Save' />
          </Col>
          <Col className='disappear_phone' />
        </Row>
      </form>
    </div>
  );
};

export default Info;
