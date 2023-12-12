import { Button, Col, Row } from 'react-bootstrap';
import { Link, createSearchParams, useNavigate } from 'react-router-dom';

import LoginImgUrl from '@assets/images/login.jpg';

import Footer from '@components/Footer';
import { useQuery } from '@tanstack/react-query';
import { TryRefresh } from '@lib/Auth';

const randomString = (length: number) => {
  const array = new Uint32Array(length);
  window.crypto.getRandomValues(array);
  return btoa(array.join('')) //
    .replace(/\+/g, '-')
    .replace(/\//g, '_')
    .replace(/=+$/, '');
};

const generateVerifier = () => {
  return randomString(4);
};

const generateChallenge = async (verifier: string) => {
  const encoder = new TextEncoder();
  const digest = await window.crypto.subtle.digest('SHA-256', encoder.encode(verifier));
  const array = Array.from(new Uint8Array(digest));
  return btoa(String.fromCharCode.apply(null, array)) //
    .replace(/\+/g, '-')
    .replace(/\//g, '_')
    .replace(/=+$/, '');
};

const Login = () => {
  const navigate = useNavigate();
  const authUrl = import.meta.env.VITE_AUTHORIZE_URL;

  const { refetch } = useQuery({
    queryKey: ['refresh'],
    queryFn: TryRefresh,
    retry: false,
    enabled: false,
  });

  const login = async () => {
    const { isSuccess } = await refetch();

    if (isSuccess) {
      console.log('success');
      navigate('/');
      return;
    }

    const state = randomString(8);
    const verifier = generateVerifier();
    const challenge = await generateChallenge(verifier);

    localStorage.setItem('state', state);
    localStorage.setItem('verifier', verifier);

    const searchParams = createSearchParams({
      client_id: 'twp',
      code_challenge: challenge,
      code_challenge_method: 'S256',
      redirect_uri: `${location.origin}/callback`,
      response_type: 'code',
      state: state,
    });

    const url = new URL(authUrl);
    url.search = searchParams.toString();

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
                <div className='title center'>Welcome Back!</div>
                <div style={{ padding: '10% 0 20% 0' }} className='white_word'>
                  <p>
                    We're thrilled to have you back with us. It's always a pleasure to see familiar
                    faces, and we're grateful for your continued support.
                  </p>
                </div>
              </Col>

              <Col xs={12}>
                <Button onClick={login} className='before_button white'>
                  <div className='center white_word pointer'>Log in</div>
                </Button>

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

export default Login;
