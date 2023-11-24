import { Button, Col, Row } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { useState } from 'react';

import Footer from '@components/Footer';
import InfoItem from '@components/InfoItem';
import PasswordItem from '@components/PasswordItem';

const Authorize = () => {
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  return (
    <div>
      <div style={{ backgroundColor: 'var(--bg)', width: '100%' }}>
        <Row style={{ width: '100%' }}>
          <Col xs={12} md={6}>
            <div
              className='flex-wrapper'
              style={{
                background: 'url("/images/login.jpg") no-repeat center center/cover',
                width: '100%',
              }}
            ></div>
          </Col>
          <Col xs={12} md={6} style={{ padding: '10% 10% 10% 10%' }}>
            <Row>
              <Col xs={12}>
                <div className='title center'> Log in</div>
                <div style={{ padding: '10% 0 20% 0' }} className='white_word'>
                  <InfoItem text='Email Address' isMore={false} value={email} setValue={setEmail} />
                  <PasswordItem text='Password' value={password} setValue={setPassword} />
                </div>
              </Col>

              <Col xs={12}>
                <Button className='before_button white'>
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

export default Authorize;
