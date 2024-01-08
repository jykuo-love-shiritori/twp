import AdminButtons from '@components/AdminButtons';
import { Col, Row } from 'react-bootstrap';
import { Outlet } from 'react-router-dom';

const Admin = () => {
  return (
    <Row style={{ width: '100%', padding: '0', margin: '0' }} className='flex_wrapper'>
      <Col xs={12} style={{ width: '100%', padding: '0' }}>
        <div className='user_bg center' />
      </Col>

      <Col xs={12} style={{ padding: '0' }}>
        <Row style={{ padding: '0', margin: '0', width: '100%' }}>
          <Col xs={12} md={3} lg={2} style={{ padding: '0' }}>
            <Row className='user_icon'>
              <Col xs={12} className='center'>
                <img src='/placeholder/person.png' className='user_img' />
              </Col>
              <Col xs={12} className='center'>
                <h4>Name</h4>
              </Col>
            </Row>

            <AdminButtons />
          </Col>
          <Col xs={12} md={9} lg={10} style={{ padding: '1% 5% 6% 5%' }}>
            <Outlet />
          </Col>
        </Row>
      </Col>
    </Row>
  );
};
export default Admin;
