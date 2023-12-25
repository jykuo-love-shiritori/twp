import { RouteOnNotOK } from '@lib/Functions';
import { useQuery } from '@tanstack/react-query';
import { useEffect, useRef, useState } from 'react';
import { Col, Row } from 'react-bootstrap';
import { useForm, SubmitHandler } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';
import defaultImageUrl from '@assets/images/person.png';

interface IBuyerInfo {
  name: string;
  email: string;
  address: string;
  image_url: string;
  image: File | null;
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
  const navigate = useNavigate();
  const { register, handleSubmit, watch, setValue, getValues, reset } = useForm<IBuyerInfo>({
    defaultValues: {
      name: 'username',
      image_url: defaultImageUrl,
      address: 'an address',
      email: 'thisIsAnEmail@mail.com',
      image: null,
    },
  });
  const OnFormOutput: SubmitHandler<IBuyerInfo> = async (data) => {
    const formData = new FormData();
    formData.append('name', data.name);
    formData.append('email', data.email);
    formData.append('address', data.address);
    if (data.image) {
      formData.append('image', data.image);
    }
    const resp = await fetch('/api/user/info', {
      method: 'PATCH',
      headers: {},
      body: formData,
    });
    if (!resp.ok) {
      RouteOnNotOK(resp, navigate);
    } else {
      alert('success');
      navigate(0);
    }
  };

  const { data: fetchedData, isLoading } = useQuery({
    queryKey: ['userGetInfo'],
    queryFn: async () => {
      const resp = await fetch('/api/user/info', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });
      if (!resp) {
        RouteOnNotOK(resp, navigate);
      }
      return resp.json();
    },
    select: (data) => data as IBuyerInfo,
    enabled: true,
    refetchOnWindowFocus: false,
  });

  // icon upload thing
  const [image, setImage] = useState<string | undefined>(undefined);
  const iconOnChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      if (!e.target.files[0].name.match(/\.(jpg|jpeg|png|gif)$/i)) {
        alert('not an image');
      } else {
        setValue('image', e.target.files[0]);
        setImage(URL.createObjectURL(e.target.files[0]));
      }
    }
  };
  const hiddenFileInput = useRef<HTMLInputElement>(null);
  const handleIconClick = () => {
    if (hiddenFileInput.current) {
      hiddenFileInput.current.click();
    }
  };

  useEffect(() => {
    reset({ ...fetchedData });
  }, [fetchedData, isLoading, reset]);

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
                      src={image ?? getValues('image_url')}
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
                Username
              </Col>
              <Col xs={12} className='form_item_wrapper'>
                <input type='text' {...register('name', { required: true })} />
              </Col>
              <Col xs={12} style={{ ...labelStyle, paddingTop: '24px' }}>
                Email
              </Col>
              <Col xs={12} className='form_item_wrapper'>
                <input type='email' {...register('email', { required: true })} />
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
