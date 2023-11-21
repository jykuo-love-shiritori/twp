import { Col, Row } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCreditCard } from '@fortawesome/free-solid-svg-icons';

import creditData from '@pages/user/buyer/security/creditCard.json';
import TButton from '@components/TButton';

const CreditCard = () => {
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
      <div className='title'>Security - Credit Card</div>
      <hr className='hr' />
      <Row>
        <Col sm={12} md={8}></Col>
        <Col sm={12} md={4}>
          <TButton text='Add New Card' url='/user/security/manageCreditCard/newCard' />
        </Col>
      </Row>
      <br />
      <Row>
        {creditData.map((data, index) => {
          return (
            <Col xs={6} md={3} key={index}>
              <div style={ContainerStyle}>
                <div className='title_color' style={{ padding: '0% 5% 5% 10%' }}>
                  <b>{data.company}</b>
                </div>
                <div className='center'>
                  <FontAwesomeIcon icon={faCreditCard} size='3x' />
                </div>
                <div className='center' style={{ padding: '5%' }}>
                  ...{data.last_four_code}
                </div>
              </div>
              <TButton text='delete' url='' />
            </Col>
          );
        })}
      </Row>
    </div>
  );
};

export default CreditCard;
