import { Button, Col, Row } from 'react-bootstrap';
import { Link, useSearchParams } from 'react-router-dom';
import { useState, FormEventHandler } from 'react';

import Footer from '@components/Footer';
import InfoItem from '@components/InfoItem';
import PasswordItem from '@components/PasswordItem';
import LoginImgUrl from '@assets/images/login.jpg';

const Authorize = () => {
  const [searchParams] = useSearchParams();

  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');

  const authUrl = '/api/oauth/authorize';

  const submitForm: FormEventHandler = async (e) => {
    e.preventDefault();

    const body = Object.fromEntries([...searchParams.entries()]);
    body['email'] = email;
    body['password'] = password;

    const resp = await fetch(authUrl, {
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
      method: 'POST',
      body: JSON.stringify(body),
    });
    const result = await resp.json();

    const redirect_uri = searchParams.get('redirect_uri');
    console.log(redirect_uri);
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
                <form onSubmit={submitForm}>
                  <div className='title center'> Log in</div>
                  <div style={{ padding: '10% 0 20% 0' }} className='white_word'>
                    <InfoItem
                      text='Email Address'
                      isMore={false}
                      value={email}
                      setValue={setEmail}
                    />
                    <PasswordItem text='Password' value={password} setValue={setPassword} />
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
    </div>
  );
};

export default Authorize;
