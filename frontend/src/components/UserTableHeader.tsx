import { Col, Row } from 'react-bootstrap';

const UserTableHeader = () => {
  return (
    <div className='disappear_tablet disappear_phone'>
      <Row className='user_table_header'>
        <Col md={1} xs={2} style={{ textAlign: 'center' }}>
          Icon
        </Col>
        <Col md={3} xs={10} style={{ textAlign: 'left' }}>
          Name
        </Col>
        <Col md={5} xs={12} style={{ textAlign: 'left' }}>
          Email
        </Col>
        <Col md={2} xs={12} style={{ textAlign: 'left' }}>
          Created
        </Col>
        <Col md={1} xs={6} style={{ textAlign: 'center' }}>
          Delete
        </Col>
      </Row>
      <hr />
    </div>
  );
};

export default UserTableHeader;
