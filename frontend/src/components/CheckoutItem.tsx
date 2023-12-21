import { Row, Col } from 'react-bootstrap';

interface Props {
  label: string;
  value: string;
  style?: React.CSSProperties;
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

const CheckoutItem = ({ label, value, style = {} }: Props) => {
  return (
    <>
      {/* layout for dektop and tablet */}
      <div className='disappear_phone'>
        <Row style={{ ...ContentStyle, ...style }}>
          <Col xs={7} style={ColStyle}>
            {label}
          </Col>
          <Col xs={5} style={ColStyle} className='right'>
            {value}
          </Col>
        </Row>
      </div>

      {/* layout for phone */}
      <div className='disappear_tablet disappear_desktop'>
        <Row style={{ ...ContentStylePhone, ...style }}>
          <Col xs={7} style={ColStyle}>
            {label}
          </Col>
          <Col xs={5} style={ColStyle} className='right'>
            {value}
          </Col>
        </Row>
      </div>
    </>
  );
};

export default CheckoutItem;
