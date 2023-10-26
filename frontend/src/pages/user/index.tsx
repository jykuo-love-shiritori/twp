import { Col, Row } from 'react-bootstrap';
import { faUser, faFile } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { Outlet } from 'react-router-dom';

const User = () => {
  return (
    <div className='user_bg'>
      <Row>
        <Col xs={12} md={12}>
          <img src='/images/user_bg.png' style={{ width: '100%' }} />
        </Col>
        <Col xs={12} md={3}>
          <Row className='user_icon'>
            <Col xs={12} className='center'>
              <img src='/images/head.png' className='user_img' />
            </Col>
            <Col xs={12} className='center'>
              <h4>John Jonathan</h4>
            </Col>
          </Row>

          <a href={'/user/info'} className='none'>
            <div className='user_button'>
              <FontAwesomeIcon icon={faUser} className='white_word' /> &nbsp;{' '}
              <span className='white_word'>Personal info</span>
            </div>
          </a>

          <a href={'/user/buyer/order'} className='none'>
            <div className='user_button'>
              <FontAwesomeIcon icon={faFile} className='white_word' /> &nbsp;{' '}
              <span className='white_word'>Order history</span>
            </div>
          </a>
        </Col>
        <Col xs={12} md={9} style={{ padding: '1% 7% 6% 7%' }}>
          <Outlet />
        </Col>
      </Row>
    </div>
  );
};

export default User;
