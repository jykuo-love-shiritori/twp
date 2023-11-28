import { useState } from 'react';
import { Col, Row } from 'react-bootstrap';

const UserTableHeader = () => {
  //resize
  const [winSize, setWinSize] = useState(window.innerWidth);
  window.addEventListener('resize', () => {
    setWinSize(window.innerWidth);
  });

  if (winSize >= 1024) {
    return (
      <>
        <Row className='user_table_header'>
          <Col md={1} xs={2} style={{ textAlign: 'center' }}>
            <h4>Icon</h4>
          </Col>
          <Col md={2} xs={10} style={{ textAlign: 'left' }}>
            <h4>Name</h4>
          </Col>
          <Col md={4} xs={12} style={{ textAlign: 'left' }}>
            <h4>Email</h4>
          </Col>
          <Col md={3} xs={12} style={{ textAlign: 'left' }}>
            <h4>Create Date</h4>
          </Col>
          <Col md={1} xs={6} style={{ textAlign: 'center' }}>
            <h4>Admin</h4>
          </Col>
          <Col md={1} xs={6} style={{ textAlign: 'center' }}>
            <h4>Delete</h4>
          </Col>
        </Row>
        <hr />
      </>
    );
  } else {
    return null;
  }
};

export default UserTableHeader;
