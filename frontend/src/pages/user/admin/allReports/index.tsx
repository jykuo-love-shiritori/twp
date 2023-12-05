import { Col, Row } from 'react-bootstrap';

import TButton from '@components/TButton';
import AdminReportItem from '@components/AdminReportItem';

import adminReportData from '@pages/user/admin/adminReportData.json';

const AdminReport = () => {
  return (
    <div>
      <Row>
        <Col sm={12} md={8}>
          <div className='title'>All Reports</div>
        </Col>
        <Col sm={12} md={4}>
          <div style={{ padding: '20px 0 0 0' }}>
            <TButton text='Generate Report' url='/admin/reports/new' />
          </div>
        </Col>
      </Row>
      <hr className='hr' />
      <Row>
        {adminReportData.map((data, index) => {
          return (
            <Col xs={6} md={3} key={index}>
              <AdminReportItem year={data.year} month={data.month} />
            </Col>
          );
        })}
      </Row>
    </div>
  );
};

export default AdminReport;
