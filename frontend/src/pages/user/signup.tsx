import { Button, Col, Row } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { useState } from 'react';

import Footer from '@components/Footer';
import InfoItem from '@components/InfoItem';
import PasswordItem from '@components/PasswordItem';
import Registerimage_url from '@assets/images/register.jpg';

const Signup = () => {
  const [name, setName] = useState<string>('');
  const [email, setEmail] = useState<string>('');
  const [password, setPassword] = useState<string>('');
  const [confirmedPassword, setConfirmedPassword] = useState<string>('');

  return (
    <div>
      <div style={{ backgroundColor: 'var(--bg)' }}>
        <Row style={{ width: '100%', padding: '0', margin: '0' }}>
          <Col xs={12} md={6} style={{ padding: '0' }}>
            <div
              className='flex-wrapper'
              style={{
                background: `url(${Registerimage_url}) no-repeat center center/cover`,
                width: '100%',
              }}
            ></div>
          </Col>
          <Col xs={12} md={6} style={{ padding: '10% 10% 10% 10%' }}>
            <Row>
              <Col xs={12}>
                <div className='title center'>Sign up</div>
                <div style={{ padding: '10% 0 20% 0' }} className='white_word'>
                  <InfoItem text='Name' isMore={false} value={name} setValue={setName} />
                  <InfoItem text='Email Address' isMore={false} value={email} setValue={setEmail} />
                  <PasswordItem text='Password' value={password} setValue={setPassword} />
                  <PasswordItem
                    text='Confirm Password'
                    value={confirmedPassword}
                    setValue={setConfirmedPassword}
                  />
                </div>
              </Col>

              <Col xs={12}>
                <Button className='before_button white'>
                  <div className='center white_word pointer'>Sign up</div>
                </Button>

                <div className='center' style={{ fontSize: '12px' }}>
                  Sign up to agree to our Terms of Use and confirm that you've read our Privacy
                  Policy.
                </div>
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
                  <span style={{ color: 'white' }}>Already have an account? &nbsp; </span>
                  <span>
                    <u>
                      <Link to='/login'>Log in</Link>
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

export default Signup;
