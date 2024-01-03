import { useForm, SubmitHandler } from 'react-hook-form';
import { Col, Row } from 'react-bootstrap';
import TButton from '@components/TButton';
import { useEffect, useRef, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { RouteOnNotOK } from '@lib/Status';
import { useQuery } from '@tanstack/react-query';
import { useAuth } from '@lib/Auth';
import defaultImageUrl from '@assets/images/defaultUserIcon.gif';

interface IShopInfo {
  name: string;
  description: string;
  enabled: boolean;
  image_url: string;
  image: File | null;
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
  boxShadow: '2px 4px 10px 2px rgba(0, 0, 0, 0.25)',
};

const labelStyle = {
  fontSize: '16px',
  fontWeight: '500',
  color: 'rgba(255, 255, 255, 0.67)',
};

const SellerInfo = () => {
  const navigate = useNavigate();
  const token = useAuth();
  const { register, handleSubmit, watch, setValue, getValues, reset } = useForm<IShopInfo>({
    defaultValues: {
      name: 'shop name',
      description: 'shop description',
      enabled: false,
      image_url: defaultImageUrl,
      image: null,
    },
  });
  const OnFormOutput: SubmitHandler<IShopInfo> = async (data) => {
    const formData = new FormData();
    formData.append('name', data.name);
    formData.append('description', data.description);
    formData.append('enabled', data.enabled ? 'true' : 'false');
    if (data.image) {
      formData.append('image', data.image, data.image.name);
    }
    const resp = await fetch('/api/seller/info', {
      method: 'PATCH',
      headers: {
        Authorization: `Bearer ${token}`,
      },
      body: formData,
    });
    if (!resp.ok) {
      RouteOnNotOK(resp, navigate);
    } else {
      navigate(0);
    }
  };

  const { data: fetchedData, isLoading } = useQuery({
    queryKey: ['sellerGetShopInfo'],
    queryFn: async () => {
      const resp = await fetch('/api/seller/info', {
        method: 'GET',
        headers: {
          Authorization: `Bearer ${token}`,
          'Content-Type': 'application/json',
        },
      });
      if (!resp) {
        RouteOnNotOK(resp, navigate);
      }
      return resp.json();
    },
    select: (data) => data as IShopInfo,
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
                      src={image ? image : getValues('image_url')}
                      style={userImgStyle}
                      onClick={handleIconClick}
                    />
                  </div>
                  <div className='center' style={{ paddingTop: '10px' }}>
                    <h5>{watch('name')}</h5>
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
                <input type='text' {...register('name', { required: true })} />
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
                    <input type='checkbox' {...register('enabled')} />
                  </Col>
                  <Col className='left'>
                    {watch('enabled')
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
