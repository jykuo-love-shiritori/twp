import { Button, Col, Row } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { useForm, SubmitHandler } from 'react-hook-form';
import { useQuery } from '@tanstack/react-query';

import Footer from '@components/Footer';
import RegisterImgUrl from '@assets/images/register.jpg';
import FormItem from '@components/FormItem';

interface SignupProps {
  email: string;
  name: string;
  password: string;
  username: string;
}

const Signup = () => {
  const { register, handleSubmit } = useForm<SignupProps>();
  const OnFormOutput: SubmitHandler<SignupProps> = async (data) => {
    console.log(data);

    // these worked
    // const response = await fetch('/api/signup', {
    //   method: 'POST',
    //   headers: {
    //     'Content-Type': 'application/json',
    //   },
    //   body: JSON.stringify(data),
    // });
    // console.log(response);

    //these not worked
    const result = useQuery({
      queryKey: ['signup'],
      queryFn: async () => {
        const response = await fetch('/api/signup', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(data),
        });
        if (!response.ok) {
          throw new Error('singup failed');
        }
        return response.json();
      },
      refetchOnWindowFocus: false,
      enabled: false,
      retry: false,
    });

    // if (!data.username.match(/^[a-zA-Z0-9]{1,32}$/)) {
    //   console.log('username should only contain letters and numbers');
    //   return;
    // }
    // if (
    //   !data.password.match(/^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,72}$/)
    // ) {
    //   console.log(
    //     'password should contain at least one uppercase letter, one lowercase letter, one number, and one special character',
    //   );
    //   return;
    // }

    console.log('fetch...');
    await result.refetch();
    if (result.isSuccess) {
      console.log('signup success');
      console.log(result.data);
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
              <Col xs={12}>
                <div className='title center'>Sign up</div>
                <div style={{ padding: '10% 0 20% 0' }} className='white_word'>
                  <form onSubmit={handleSubmit(OnFormOutput)}>
                    <FormItem label='email'>
                      <input type='email' {...register('email', { required: true })} />
                    </FormItem>
                    <FormItem label='username'>
                      <input type='text' {...register('username', { required: true })} />
                    </FormItem>
                    <FormItem label='name'>
                      <input type='text' {...register('name', { required: true })} />
                    </FormItem>
                    <FormItem label='Password'>
                      <input type='password' {...register('password', { required: true })} />
                    </FormItem>
                  </form>
                </div>
              </Col>

              <Col xs={12}>
                <Button className='before_button white' onClick={handleSubmit(OnFormOutput)}>
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
