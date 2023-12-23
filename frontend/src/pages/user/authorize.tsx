import { Button, Col, Row } from 'react-bootstrap';
import { Link, useNavigate, useSearchParams } from 'react-router-dom';
import { SubmitHandler, useForm } from 'react-hook-form';
import { useState } from 'react';

import WarningModal from '@components/WarningModal';
import Footer from '@components/Footer';
import LoginImgUrl from '@assets/images/login.jpg';
import FormItem from '@components/FormItem';

interface FormProps {
  email: string;
  password: string;
}

const Authorize = () => {
  const [searchParams] = useSearchParams();
  const [show, setShow] = useState<boolean>(false);
  const [warningText, setWarningText] = useState<string>('');
  const navigate = useNavigate();

  const { register, handleSubmit } = useForm<FormProps>();

  const authUrl = '/api/oauth/authorize';
  const body = Object.fromEntries([...searchParams.entries()]);

  const submitForm: SubmitHandler<FormProps> = async (data) => {
    const resp = await fetch(authUrl, {
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
      method: 'POST',
      body: JSON.stringify({ ...body, ...data }),
    });
    const result = await resp.json();

    if (!resp.ok) {
      console.log(resp);
      if (resp.status === 500) {
        navigate('/login');
      } else {
        setWarningText(result.message);
        setShow(true);
      }
      return;
    }

    const redirect_uri = searchParams.get('redirect_uri');
    if (!redirect_uri) {
      alert('No redirect uri set');
      return;
    }

    const url = new URL(redirect_uri);
    url.searchParams.set('state', result.state);
    url.searchParams.set('code', result.code);

    window.location.href = url.toString();
  };

  return (
    <div>
      <div style={{ backgroundColor: 'var(--bg)', width: '100%' }}>
        <Row style={{ width: '100%', padding: '0', margin: '0' }}>
          <Col xs={12} md={6} style={{ padding: '0' }}>
            <div
              className='flex-wrapper'
              style={{
                background: `url(${LoginImgUrl}) no-repeat center center/cover`,
                width: '100%',
              }}
            ></div>
          </Col>
          <Col xs={12} md={6} style={{ padding: '10% 10% 10% 10%' }}>
            <Row>
              <Col xs={12}>
                <form onSubmit={handleSubmit(submitForm)}>
                  <div className='title center'> Log in</div>
                  <div style={{ padding: '10% 0 20% 0' }} className='white_word'>
                    <FormItem label='Email Address or Username'>
                      <input type='text' {...register('email', { required: true })} />
                    </FormItem>
                    <FormItem label='Password'>
                      <input type='password' {...register('password', { required: true })} />
                    </FormItem>
                  </div>
                  <Button className='before_button white' type='submit'>
                    <div className='center white_word pointer'>Log in</div>
                  </Button>
                </form>

                <div className='center' style={{ fontSize: '12px' }}></div>
                <br />

                <Row>
                  <Col xs={4}>
                    <hr style={{ color: 'white' }} />
                  </Col>
                  <Col xs={4} className='center'>
                    <p>Or With</p>
                  </Col>
                  <Col xs={4}>
                    <hr style={{ color: 'white' }} />
                  </Col>
                </Row>

                <div className='center'>
                  <span style={{ color: 'white' }}>Don’t have an account ? &nbsp; </span>
                  <span>
                    <u>
                      <Link to='/signup'>Sign up</Link>
                    </u>
                  </span>
                </div>
              </Col>
            </Row>
          </Col>
        </Row>
      </div>
      <Footer />

      <WarningModal show={show} text={warningText} onHide={() => setShow(false)} />
    </div>
  );
};

export default Authorize;
