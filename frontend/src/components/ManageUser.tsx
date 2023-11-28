import UserTableRow from '@components/UserTableRow';
import { Col, Row } from 'react-bootstrap';
import datas from '@pages/user/admin/UserData.json';
import Pagination from '@components/Pagination';

const ManageUser = () => {
  return (
    <div style={{ display: 'flex', flexDirection: 'column', height: '100%' }}>
      <div style={{ flexGrow: '9' }}>
        {/* title */}
        <Row>
          <Col md={12} xs={12} className='title'>
            Manage Users
          </Col>
        </Row>

        {/* table header */}
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

        {/* table body */}
        {datas.map((data) => (
          <UserTableRow data={data} />
        ))}
      </div>

      {/* pagination and comfirm buttom*/}
      <div style={{ display: 'flex', flexGrow: '1', flexDirection: 'row' }}>
        <div className='manager_user_bottem'>
          <Pagination currentPageInit={1} totalPage={10} />
        </div>
        <div className='manager_user_bottem'>
          <div className='manage_user_confirm_button center center_vertical'>Confirm</div>
        </div>
      </div>
    </div>
  );
};
export default ManageUser;
