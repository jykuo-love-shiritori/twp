import AdminButtons from '@components/AdminButtons';
import { Col, Row } from 'react-bootstrap';
import { Outlet } from 'react-router-dom';

const Admin = () => {
  return (
    <Row style={{ flexGrow: '1' }}>
      <Col xs={12} md={3} lg={2}>
        <div className='user_bg center' style={{ width: '100%' }} />
        <Row className='user_icon'>
          <Col xs={12} className='center'>
            <img src='../../images/person.png' className='user_img' />
          </Col>
          <Col xs={12} className='center'>
            <h4>Name</h4>
          </Col>
        </Row>
        <AdminButtons />
      </Col>

      <Col xs={12} md={9} lg={10} style={{ padding: '3% 5% 0% 5%' }}>
        <Outlet />
      </Col>
    </Row>
  );
};
export default Admin;
