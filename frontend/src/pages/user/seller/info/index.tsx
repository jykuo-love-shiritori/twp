import { useForm, SubmitHandler } from 'react-hook-form';
import { Col, Row } from 'react-bootstrap';
import TButton from '@components/TButton';
import { useRef, useState } from 'react';

interface ShopInfoProps {
  shopName: string;
  shopIcon: File | undefined;
  visibility: boolean;
  description: string;
}

const SellerItemStyle = {
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
};

const labelStyle = {
  fontSize: '16px',
  fontWeight: '500',
  color: '#FFFFFFAA',
};

const SellerInfo = () => {
  //TODO: get the info from existing shop
  const { register, handleSubmit, watch, setValue } = useForm<ShopInfoProps>({
    defaultValues: {
      shopName: 'shop name',
      shopIcon: undefined,
      visibility: false,
      description: 'shop description',
    },
  });
  const OnFormOutput: SubmitHandler<ShopInfoProps> = (data) => {
    console.log(data);
    return data;
  };

  // icon upload thing
  const [image, setImage] = useState<string | undefined>(undefined);
  const iconOnChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      //TODO? check file type
      setValue('shopIcon', e.target.files[0]);
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
      <div className='title'>Shop info</div>
      <hr className='hr' />
      <form onSubmit={handleSubmit(OnFormOutput)}>
        <Row>
          {/* left half */}
          <Col xs={12} md={6}>
            <Row className='center' style={{ height: '100%' }}>
              <Col />
              <Col xs={8} xl={7}>
                <div style={SellerItemStyle}>
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
                    <h5>{watch('shopName')}</h5>
                  </div>
                  <TButton text='View Shop' />
                </div>
              </Col>
              <Col />
            </Row>
          </Col>

          {/* right half */}
          <Col xs={12} md={6}>
            <Row>
              <Col xs={12} style={labelStyle}>
                Shop Name
              </Col>
              <Col xs={12} className='form_item_wrapper'>
                <input type='text' {...register('shopName', { required: true })} />
              </Col>

              <Col xs={12} style={{ ...labelStyle, paddingTop: '24px' }}>
                Visibility
              </Col>
              <Col xs={12}>
                <Row>
                  <Col
                    xs='auto'
                    className='center'
                    style={{ ...labelStyle, padding: '0 0 0 24px' }}
                  >
                    <input type='checkbox' {...register('visibility')} />
                  </Col>
                  <Col className='left'>
                    {watch('visibility')
                      ? 'Your shop is visible to everyone.🌞'
                      : 'Your shop is hidden from everyone.🌚'}
                  </Col>
                </Row>
              </Col>

              <Col xs={12} style={{ ...labelStyle, paddingTop: '24px' }}>
                Description
              </Col>
              <Col xs={12} className='form_item_wrapper'>
                <textarea {...register('description', { required: true })} />
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
