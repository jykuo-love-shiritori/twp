import { Col, Row } from 'react-bootstrap';

import TButton from '@components/TButton';
import SellerReportItem from '@components/SellerReportItem';

import sellerReportData from '@pages/user/seller/sellerReportData.json';

const SellerReport = () => {
  return (
    <div>
      <Row>
        <Col sm={12} md={8}>
          <div className='title'>All Reports</div>
        </Col>
        <Col sm={12} md={4}>
          <div style={{ padding: '20px 0 0 0' }}>
            <TButton text='Generate Report' url='/user/seller/reports/new' />
          </div>
        </Col>
      </Row>
      <hr className='hr' />
      <Row>
        {sellerReportData.map((data, index) => {
          return (
            <Col xs={6} md={3} key={index}>
              <SellerReportItem year={data.year} month={data.month} />
            </Col>
          );
        })}
      </Row>
    </div>
  );
};

export default SellerReport;
