import { Button, Col, Row } from 'react-bootstrap';
import { Link } from 'react-router-dom';

import LoginImgUrl from '@assets/images/login.jpg';

import Footer from '@components/Footer';

const Login = () => {
  console.log(LoginImgUrl);
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
                <Link to={'/authorize'}>
                  <Button className='before_button white'>
                    <div className='center white_word pointer'>Log in</div>
                  </Button>
                </Link>

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
