import { Col, Row } from 'react-bootstrap';
import UserTableRow from '@components/UserTableRow';
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
        {datas.map((data, index) => (
          <UserTableRow data={data} key={index} />
        ))}
      </div>
      <div className='center' style={{ display: 'flex', flexDirection: 'row' }}>
        <Row style={{ width: '100%' }}>
          {/* pagination */}
          <Col className='center' xl={6} md={12} xs={12} style={{ margin: '5px 0 0 0 ' }}>
            <Pagination currentPageInit={1} totalPage={10} />
          </Col>

          {/* comfirm buttom */}
          <Col className='center' xl={6} md={12} xs={12} style={{ margin: '5px 0 5px 0' }}>
            <div className='manage_user_confirm_button center center_vertical'>Confirm</div>
          </Col>
        </Row>
      </div>
    </div>
  );
};
export default ManageUser;
