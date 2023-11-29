import { Col, Row } from 'react-bootstrap';

interface UserTableHeaderProps {
  isBigScreen: boolean;
}

const UserTableHeader = ({ isBigScreen = true }: UserTableHeaderProps) => {
  if (isBigScreen) {
    // layout for big screen
    return (
      <>
        <Row className='user_table_header' style={{ fontSize: '20px' }}>
          <Col md={1} xs={2} style={{ textAlign: 'center' }}>
            Icon
          </Col>
          <Col md={2} xs={10} style={{ textAlign: 'left' }}>
            Name
          </Col>
          <Col md={4} xs={12} style={{ textAlign: 'left' }}>
            Email
          </Col>
          <Col md={3} xs={12} style={{ textAlign: 'left' }}>
            Create Date
          </Col>
          <Col md={1} xs={6} style={{ textAlign: 'center' }}>
            Admin
          </Col>
          <Col md={1} xs={6} style={{ textAlign: 'center' }}>
            Delete
          </Col>
        </Row>
        <hr />
      </>
    );
  } else {
    // dont show header for small screen
    return null;
  }
};

export default UserTableHeader;
