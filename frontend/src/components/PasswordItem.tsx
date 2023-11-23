import { Col, Row } from 'react-bootstrap';

interface Props {
  text: string;
}

const PasswordItem = ({ text }: Props) => {
  return (
    <Row style={{ margin: '2% 0% 2% 0% ' }}>
      <Col xs={12} md={4} className='center_vertical'>
        {text}
      </Col>
      <Col xs={12} md={8}>
        <input type='password' placeholder={text} className='inputBox' />
      </Col>
    </Row>
  );
};

export default PasswordItem;