import { Col, Row } from 'react-bootstrap';
import UserTableRow from '@components/UserTableRow';
import datas from '@pages/user/admin/UserData.json';
import Pagination from '@components/Pagination';
import UserTableHeader from '@components/UserTableHeader';

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
        <Row style={{ width: '100%', margin: '2% 0 2% 0 ' }}>
          {/* pagination */}
          <Col className='center' xl={6} xs={12}>
            <Pagination currentPageInit={1} totalPage={10} />
          </Col>
          <Col className='disappear_desktop' xs={12} style={{ height: '1vh' }} />
          {/* comfirm buttom */}
          <Col className='center' xl={6} xs={12}>
            <div className='manage_user_confirm_button center center_vertical'>Confirm</div>
          </Col>
        </Row>
      </div>
    </div>
  );
};
export default ManageUser;
