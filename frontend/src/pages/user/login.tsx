import { Col, Row } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import Carousel from 'react-bootstrap/Carousel';

import Footer from '@components/Footer';
import InfoItem from '@components/InfoItem';
import { useState } from 'react';
import PasswordItem from '@components/PasswordItem';

const Login = () => {
  const [activeIndex, setActiveIndex] = useState(0);

  const handleSelect = (selectedIndex: number) => {
    setActiveIndex(selectedIndex);
  };

  const handleButtonClick = () => {
    if (activeIndex === 0) {
      handleSelect(1);
    }
  };

  return (
    <div>
      <div className='center' style={{ backgroundColor: 'var(--bg)' }}>
        <Row>
          <Col xs={12} md={6}>
            <img
              src={'/images/login.jpg'}
              style={{ height: '100%', width: '100%' }}
              className='flex-wrapper'
            />
          </Col>
          <Col xs={12} md={6} style={{ padding: '10% 10% 10% 10%' }}>
            <Row>
              <Col xs={12}>
                <Carousel
                  controls={false}
                  indicators={false}
                  interval={null}
                  activeIndex={activeIndex}
                  onSelect={handleSelect}
                >
                  <Carousel.Item className='center'>
                    <div className='title center'>Welcome Back!</div>
                    <div style={{ padding: '10% 0 20% 0' }} className='white_word'>
                      <p>
                        We're thrilled to have you back with us. It's always a pleasure to see
                        familiar faces, and we're grateful for your continued support.
                      </p>
                    </div>
                  </Carousel.Item>
                  <Carousel.Item className='center'>
                    <div className='title center'> Log in</div>
                    <div style={{ padding: '10% 0 20% 0' }} className='white_word'>
                      <InfoItem text='Email Address' isMore={false} />
                      <PasswordItem text='Password' />
                    </div>
                  </Carousel.Item>
                </Carousel>
              </Col>

              <Col xs={12}>
                <div className='before_button white'>
                  <div className='center white_word pointer' onClick={handleButtonClick}>
                    Log in
                  </div>
                </div>

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
