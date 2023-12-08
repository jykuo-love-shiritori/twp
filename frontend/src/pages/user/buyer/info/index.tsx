import { useRef, useState } from 'react';
import { Col, Row } from 'react-bootstrap';
import { useForm, SubmitHandler } from 'react-hook-form';

interface BuyerInfoProps {
  name: string;
  email: string;
  address: string;
  image: File | undefined;
}

const BuyerItemStyle = {
  borderRadius: '10px',
  padding: '10% 10% 5% 10%',
  border: '1px solid var(--button_border, #34977F)',
  background: 'var(--button_dark, #135142)',
};

const userImgStyle = {
  borderRadius: '50%',
  width: '100px',
  height: '100px',
  cursor: 'pointer',
  boxShadow: '2px 4px 10px 2px rgba(0, 0, 0, 0.25)',
};

const labelStyle = {
  fontSize: '16px',
  fontWeight: '500',
  color: 'rgba(255, 255, 255, 0.67)',
};

const Info = () => {
  //TODO: get the info from existing user
  const { register, handleSubmit, watch, setValue } = useForm<BuyerInfoProps>({
    defaultValues: {
      name: 'username',
      image: undefined,
      address: 'an address',
      email: 'thisIsAnEmail@mail.com',
    },
  });
  const OnFormOutput: SubmitHandler<BuyerInfoProps> = (data) => {
    console.log(data);
    return data;
  };

  // icon upload thing
  const [image, setImage] = useState<string | undefined>(undefined);
  const iconOnChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      //TODO? check file type
      setValue('image', e.target.files[0]);
      setImage(URL.createObjectURL(e.target.files[0]));
    }
  };
  const hiddenFileInput = useRef<HTMLInputElement>(null);
  const handleIconClick = () => {
    if (hiddenFileInput.current) {
      hiddenFileInput.current.click();
    }
  };

  return (
    <div>
      <div className='title'>Personal info</div>
      <hr className='hr' />
      <form onSubmit={handleSubmit(OnFormOutput)}>
        <Row>
          {/* left half */}
          <Col xs={12} md={6}>
            <Row className='center' style={{ height: '100%' }}>
              <Col />
              <Col xs={8} xl={7}>
                <div style={BuyerItemStyle}>
                  {/* upload img */}
                  <div className='center'>
                    <input
                      type='file'
                      onChange={iconOnChange}
                      style={{ display: 'none' }}
                      ref={hiddenFileInput}
                    />
                    <img
                      src={image ? image : '/placeholder/person.png'}
                      style={userImgStyle}
                      onClick={handleIconClick}
                    />
                  </div>
                  <div className='center' style={{ paddingTop: '10px' }}>
                    <h5>{watch('name')}</h5>
                  </div>
                </div>
              </Col>
              <Col />
            </Row>
          </Col>

          {/* right half */}
          <Col xs={12} md={6}>
            <Row>
              <Col xs={12} style={labelStyle}>
                Userame
              </Col>
              <Col xs={12} className='form_item_wrapper'>
                <input type='text' {...register('name', { required: true })} />
              </Col>

              <Col xs={12} style={{ ...labelStyle, paddingTop: '24px' }}>
                Email
              </Col>
              <Col xs={12} className='form_item_wrapper'>
                <input type='text' {...register('email', { required: true })} />
              </Col>

              <Col xs={12} style={{ ...labelStyle, paddingTop: '24px' }}>
                Address
              </Col>
              <Col xs={12} className='form_item_wrapper'>
                <textarea {...register('address', { required: true })} />
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

export default Info;
