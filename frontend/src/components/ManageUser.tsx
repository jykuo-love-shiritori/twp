import UserTableRow from '@components/UserTableRow';
import { Col, Row } from 'react-bootstrap';
import datas from '@pages/user/admin/UserData.json';
import Pagination from '@components/Pagination';
import UserTableHeader from './UserTableHeader';

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
        <UserTableHeader />

        {/* table body */}
        {datas.map((data) => (
          <UserTableRow data={data} />
        ))}
      </div>

      {/* pagination and comfirm buttom*/}
      <div style={{ display: 'flex', flexGrow: '1', flexDirection: 'row' }}>
        <Row style={{ width: '100%' }}>
          <Col className='center' md={6} xs={12} style={{ margin: '5px 0' }}>
            <Pagination currentPageInit={1} totalPage={10} />
          </Col>
          <Col className='center' md={6} xs={12} style={{ margin: '5px 0' }}>
            <div className='manage_user_confirm_button center center_vertical'>Confirm</div>
          </Col>
        </Row>
      </div>
    </div>
  );
};
export default ManageUser;
