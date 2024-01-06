import { faBan, faPlus } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { Row, Col } from 'react-bootstrap';

interface ICheckoutCoupon {
  description: string;
  discount: number;
  discount_value: number;
  id: number;
  name: string;
  scope: 'global' | 'shop';
  type: 'percentage' | 'fixed' | 'shipping';
}

// optional is because the row "Add More"
interface Props {
  coupon?: ICheckoutCoupon;
  onClick: () => void;
  isAddMore?: boolean;
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

const CheckoutItemCoupon = ({ coupon = undefined, onClick, isAddMore = false }: Props) => {
  return (
    <>
      {/* layout for desktop and tablet */}
      <div className='disappear_phone'>
        <Row style={ContentStyle}>
          <Col xs={7} style={ColStyle}>
            <Row>
              <Col xs='auto'>
                <FontAwesomeIcon
                  className='checkout_button'
                  icon={isAddMore ? faPlus : faBan}
                  size='sm'
                  onClick={onClick}
                />
              </Col>
              <Col style={{ color: isAddMore ? 'rgb(133, 133, 133)' : 'black' }}>
                {isAddMore ? 'Add More' : `${coupon?.name}`}
              </Col>
            </Row>
          </Col>
          <Col xs={5} style={ColStyle} className='right'>
            {isAddMore ? '' : `${coupon?.discount_value} NTD`}
          </Col>
        </Row>
      </div>

      {/* layout for phone */}
      <div className='disappear_tablet disappear_desktop'>
        <Row style={ContentStylePhone}>
          <Col xs={7} style={ColStyle}>
            <Row>
              <Col xs='auto'>
                <FontAwesomeIcon
                  className='checkout_button'
                  icon={isAddMore ? faPlus : faBan}
                  size='sm'
                  onClick={onClick}
                />
              </Col>
              <Col style={{ color: isAddMore ? 'rgb(133, 133, 133)' : 'black' }}>
                {isAddMore ? 'Add More' : `${coupon?.name}`}
              </Col>
            </Row>
          </Col>
          <Col xs={5} style={ColStyle} className='right'>
            {isAddMore ? '' : `${coupon?.discount_value} NTD`}
          </Col>
        </Row>
      </div>
    </>
  );
};

export default CheckoutItemCoupon;
