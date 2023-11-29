import { Col, Row } from 'react-bootstrap';
import { useState } from 'react';
import UserTableRow from '@components/UserTableRow';
import datas from '@pages/user/admin/UserData.json';
import Pagination from '@components/Pagination';
import UserTableHeader from './UserTableHeader';

const ManageUser = () => {
  //resize
  const [winSize, setWinSize] = useState(window.innerWidth);
  window.addEventListener('resize', () => {
    setWinSize(window.innerWidth);
  });

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
        <UserTableHeader isBigScreen={winSize >= 1024} />

        {/* table body */}
        {datas.map((data) => (
          <UserTableRow data={data} isBigScreen={winSize >= 1024} />
        ))}
      </div>

      {/* pagination and comfirm buttom*/}
      <div style={{ display: 'flex', flexGrow: '1', flexDirection: 'row' }}>
        <Row style={{ width: '100%' }}>
          <Col className='center' md={winSize >= 1024 ? 6 : 12} xs={12} style={{ margin: '5px 0' }}>
            <Pagination currentPageInit={1} totalPage={10} />
          </Col>
          <Col className='center' md={winSize >= 1024 ? 6 : 12} xs={12} style={{ margin: '5px 0' }}>
            <div className='manage_user_confirm_button center center_vertical'>Confirm</div>
          </Col>
        </Row>
      </div>
    </div>
  );
};
export default ManageUser;
