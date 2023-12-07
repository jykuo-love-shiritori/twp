import { useForm, SubmitHandler } from 'react-hook-form';
import { Col, Row } from 'react-bootstrap';

interface ShopInfoProps {
  shopName: string;
  shopIconUrl: string;
  visibility: boolean;
  description: string;
}

const SellerInfo = () => {
  //TODO: get the info from existing shop
  const { register, handleSubmit, watch } = useForm<ShopInfoProps>({
    defaultValues: {
      shopName: 'shop name',
      shopIconUrl: '',
      visibility: false,
      description: 'shop description',
    },
  });
  const OnFormOutput: SubmitHandler<ShopInfoProps> = (data) => {
    console.log(data);
    return data;
  };

  return (
    <div>
      <div className='title'>Shop info</div>
      <hr className='hr' />
      <form onSubmit={handleSubmit(OnFormOutput)}>
        <Row>
          {/* left half */}
          <Col xs={12} md={6}>
            {/* <div className='user_icon'>
              <img src='/placeholder/person.png' alt='user' />
            </div> */}
          </Col>

          {/* right half */}
          <Col xs={12} md={6}>
            <Row>
              <Col xs={12}>Visibility</Col>
              <Col xs={12}>
                <Row>
                  <Col xs='auto' className='center' style={{ padding: '0 0 0 24px' }}>
                    <input type='checkbox' {...register('visibility', { required: true })} />
                  </Col>
                  <Col className='left'>
                    {watch('visibility')
                      ? 'Your shop is visible to everyone.🌞'
                      : 'Your shop is hidden from everyone.🌚'}
                  </Col>

                  <Col xs={12} style={{ paddingTop: '24px' }}>
                    Description
                  </Col>
                  <Col xs={12} className='form_item_wrapper'>
                    <textarea {...register('description', { required: true })} />
                  </Col>
                </Row>
              </Col>
            </Row>
          </Col>

          {/* submit button */}
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

export default SellerInfo;
