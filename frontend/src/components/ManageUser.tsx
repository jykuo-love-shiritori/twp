import UserTableRow from '@components/UserTableRow';
import { Col, Row } from 'react-bootstrap';
import datas from '@pages/user/admin/UserData.json';
import Pagination from '@components/Pagination';

const tableColStyle = {
  fontSize: '16px',
  padding: '1% 1% 0% 1%',
  fontWeight: 'bold',
};

const ManageUser = () => {
  return (
    <>
      <Row>
        <Col md={12} xs={12} className='title'>
          Manage Users
        </Col>
      </Row>
      <Row>
        <Col md={1} xs={2} style={{ ...tableColStyle, textAlign: 'center' }}>
          <h4>Icon</h4>
        </Col>
        <Col md={2} xs={10} style={{ ...tableColStyle, textAlign: 'left' }}>
          <h4>Name</h4>
        </Col>
        <Col md={4} xs={12} style={{ ...tableColStyle, textAlign: 'left' }}>
          <h4>Email</h4>
        </Col>
        <Col md={3} xs={12} style={{ ...tableColStyle, textAlign: 'left' }}>
          <h4>Create Date</h4>
        </Col>
        <Col md={1} xs={6} style={{ ...tableColStyle, textAlign: 'center' }}>
          <h4>Admin</h4>
        </Col>
        <Col md={1} xs={6} style={{ ...tableColStyle, textAlign: 'center' }}>
          <h4>Delete</h4>
        </Col>

        <hr />
        {datas.map((data) => (
          <UserTableRow data={data} />
        ))}
      </Row>
      <Pagination currentPageInit={1} totalPage={10} />
    </>
  );
};
export default ManageUser;
