import UserTableRow from '@components/UserTableRow';
import { Col, Row } from 'react-bootstrap';
import datas from '@pages/user/admin/UserData.json';
import Pagination from '@components/Pagination';

const ManageUser = () => {
  return (
    <div style={{ display: 'flex', flexDirection: 'column', height: '100%' }}>
      <div style={{ flexGrow: '9' }}>
        <Row>
          <Col md={12} xs={12} className='title'>
            Manage Users
          </Col>
        </Row>
        <Row className='userTableHeader'>
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

        {datas.map((data) => (
          <UserTableRow data={data} />
        ))}
      </div>
      <div style={{ flexGrow: '1', alignSelf: 'center' }}>
        <Pagination currentPageInit={1} totalPage={10} />
      </div>
    </div>
  );
};
export default ManageUser;
