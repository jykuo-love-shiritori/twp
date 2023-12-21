import '@style/global.css';

import { Col, Row } from 'react-bootstrap';
import { useParams } from 'react-router-dom';

import SellerItem from '@components/SellerItem';
import NotFound from '@components/NotFound';

import adminReportData from '@pages/user/admin/adminReportData.json';

interface Sellers {
  id: number;
  name: string;
  imgUrl: string;
  amount: number;
}

interface Props {
  year: number;
  month: number;
  sellers: Sellers[];
  totalAmount: number;
  id: number;
}

const AdminReportEach = () => {
  const reportPageStyle = {
    borderRadius: '50px 50px 0px 0px',
    border: '1px solid var(--button_border)',
    background: 'var(--bg)',
    boxShadow: '0px 4px 30px 2px var(--title)',
    padding: ' 7% 12% 10% 12%',
  };

  const { year, month } = useParams();

  const data: Props | undefined = adminReportData.find(
    (item) => item.year.toString() === year && item.month.toString() === month,
  );

  if (data) {
    return (
      <div style={{ padding: '10% 10% 0% 10%' }}>
        <div className='flex_wrapper' style={reportPageStyle}>
          <h1 className='title_color' style={{ paddingBottom: '30px' }}>
            <b>
              {data.year} / {data.month} Finical Report
            </b>
          </h1>
          <h4>The top three highest monthly sales</h4>
          <hr className='hr' />
          <Row>
            {data.sellers.map((item, index) => {
              return (
                <Col xs={12} md={4} key={index}>
                  <div className='center title_color' style={{ paddingTop: '30px' }}>
                    <h4>No. {index + 1}</h4>
                  </div>
                  <div className='center' style={{ paddingBottom: '30px' }}>
                    TWD $ {item.amount}
                  </div>

                  <SellerItem id={item.id} />
                </Col>
              );
            })}
          </Row>

          <hr className='hr' />

          <Row>
            <Col xs={12} className='center_vertical'>
              <h4>
                <b>Total sales revenue :</b>
              </h4>
            </Col>
            <Col xs={12} className='center_vertical'>
              TWD $ {data.totalAmount}
            </Col>
          </Row>
        </div>
      </div>
    );
  } else {
    return <NotFound />;
  }
};

export default AdminReportEach;