import { Button, Col, Row } from 'react-bootstrap';
import { Link, useNavigate } from 'react-router-dom';
import { useForm, SubmitHandler } from 'react-hook-form';

import Footer from '@components/Footer';
import RegisterImgUrl from '@assets/images/register.jpg';
import FormItem from '@components/FormItem';
import WarningModal from '@components/WarningModal';
import { useState } from 'react';

interface SignupProps {
  email: string;
  name: string;
  password: string;
  username: string;
}

const Signup = () => {
  const navigate = useNavigate();
  const [show, setShow] = useState<boolean>(false);
  const [warningText, setWarningText] = useState<string>('');

  const { register, handleSubmit } = useForm<SignupProps>();
  const OnFormOutput: SubmitHandler<SignupProps> = async (data) => {
    if (!data.username.match(/^[a-zA-Z0-9]{1,32}$/)) {
      setWarningText('username should only contain letters and numbers\n');
      setShow(true);
      return;
    } else if (
      !data.password.match(/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,72}$/)
    ) {
      setWarningText(
        'password should contain at least one of each: uppercase letter, lowercase letter, number and special character\n',
      );
      setShow(true);
      return;
    }

    const response = await fetch('/api/signup', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });
    if (!response.ok) {
      const error = await response.json();
      setWarningText(error.message);
      setShow(true);
      return;
    } else {
      navigate('/login');
    }
  };

  return (
    <div>
      <div style={{ backgroundColor: 'var(--bg)' }}>
        <Row style={{ width: '100%', padding: '0', margin: '0' }}>
          <Col xs={12} md={6} style={{ padding: '0' }}>
            <div
              className='flex-wrapper'
              style={{
                background: `url(${RegisterImgUrl}) no-repeat center center/cover`,
                width: '100%',
              }}
            ></div>
          </Col>
          <Col xs={12} md={6} style={{ padding: '10% 10% 10% 10%' }}>
            <Row>
              <form onSubmit={handleSubmit(OnFormOutput)}>
                <Col xs={12}>
                  <div className='title center'>Sign up</div>
                  <div style={{ padding: '10% 0 20% 0' }} className='white_word'>
                    <FormItem label='email'>
                      <input type='email' {...register('email', { required: true })} />
                    </FormItem>
                    <FormItem label='username'>
                      <input type='text' {...register('username', { required: true })} />
                    </FormItem>
                    <FormItem label='name'>
                      <input type='text' {...register('name', { required: true })} />
                    </FormItem>
                    <FormItem label='password'>
                      <input type='password' {...register('password', { required: true })} />
                    </FormItem>
                  </div>
                </Col>

                <Col xs={12}>
                  <Button className='before_button white' type='submit'>
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
              </form>
            </Row>
          </Col>
        </Row>
      </div>
      <Footer />

      <WarningModal show={show} onHide={() => setShow(false)} text={warningText} />
    </div>
  );
};

export default Signup;
