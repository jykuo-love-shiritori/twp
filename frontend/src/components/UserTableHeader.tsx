import { Col, Row } from 'react-bootstrap';

interface UserTableHeaderProps {
  isBigScreen: boolean;
}

const UserTableHeader = ({ isBigScreen = true }: UserTableHeaderProps) => {
  if (isBigScreen) {
    // layout for big screen
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
    // dont show header for small screen
    return null;
  }
};

export default UserTableHeader;
