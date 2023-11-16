import { Col, Row } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faLock, faCreditCard } from '@fortawesome/free-solid-svg-icons';
import { Link } from 'react-router-dom';

const Security = () => {
  const ContainerStyle = {
    borderRadius: '24px',
    border: '1px solid var(--button_border, #34977F)',
    background: ' var(--button_dark, #135142)',
    padding: '10% 5% 5% 5%',
    color: 'white',
    marginBottom: '15px',
  };

  return (
    <div>
      <div className='title'>Security</div>
      <hr className='hr' />
      <Row>
        <Col xs={6} md={3}>
          <Link to={'/user/security/password'} className='none'>
            <div style={ContainerStyle}>
              <div className='center'>
                <FontAwesomeIcon icon={faLock} size='3x' />
              </div>
              <div className='center' style={{ padding: '5%' }}>
                Passsword
              </div>
            </div>
          </Link>
        </Col>
        <Col xs={6} md={3}>
          <Link to={'/user/security/manageCreditCard'} className='none'>
            <div style={ContainerStyle}>
              <div className='center'>
                <FontAwesomeIcon icon={faCreditCard} size='3x' />
              </div>
              <div className='center' style={{ padding: '5%' }}>
                Credit Card
              </div>
            </div>
          </Link>
        </Col>
      </Row>
    </div>
  );
};

export default Security;
