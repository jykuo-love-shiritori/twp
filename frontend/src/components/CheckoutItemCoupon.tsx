import { faBan } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { Row, Col } from 'react-bootstrap';

interface Props {
  coupon: {
    description: string;
    discount: number;
    discount_value: number;
    id: number;
    name: string;
    scope: string; // 'global' | 'shop'
    type: string; // 'percentage' | 'fixed' | 'shipping'
  };
  onClick: () => void;
}

const ContentStyle = {
  fontSize: '18px',
  fontWeight: '500',
  margin: '1% 6%',
};

const ContentStylePhone = {
  fontSize: '14px',
  fontWeight: '500',
  margin: '1% 6%',
};

const ColStyle = {
  padding: '0',
  margin: '0',
};

const CheckoutItemCoupon = ({ coupon, onClick }: Props) => {
  return (
    <>
      {/* layout for dektop and tablet */}
      <div className='disappear_phone'>
        <Row style={{ ...ContentStyle }}>
          <Col xs={7} style={ColStyle}>
            <Row>
              <Col xs='auto'>
                <FontAwesomeIcon
                  className='checkout_button'
                  icon={faBan}
                  size='sm'
                  onClick={onClick}
                />
              </Col>
              {coupon.name}
            </Row>
          </Col>
          <Col xs={5} style={ColStyle} className='right'>
            {coupon.discount_value} NTD
          </Col>
        </Row>
      </div>

      {/* layout for phone */}
      <div className='disappear_tablet disappear_desktop'>
        <Row style={{ ...ContentStylePhone }}>
          <Col xs={7} style={ColStyle}>
            <Row>
              <Col xs='auto'>
                <FontAwesomeIcon
                  className='checkout_button'
                  icon={faBan}
                  size='sm'
                  onClick={onClick}
                />
              </Col>
              {coupon.name}
            </Row>
          </Col>
          <Col xs={5} style={ColStyle} className='right'>
            {coupon.discount_value} NTD
          </Col>
        </Row>
      </div>
    </>
  );
};

export default CheckoutItemCoupon;