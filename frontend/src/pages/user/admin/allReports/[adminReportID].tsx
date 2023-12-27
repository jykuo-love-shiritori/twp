import '@style/global.css';

import { Col, Row } from 'react-bootstrap';
import { useParams } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';

import SellerItem, { SellerItemProps } from '@components/SellerItem';
import NotFound from '@components/NotFound';

import { CheckFetchStatus } from '@lib/Status';

interface SellersProps extends SellerItemProps {
  total_sales: number;
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
  const yearString = year ?? '';
  const monthString = month ?? '';

  const rfc3339Date = new Date(
    Date.UTC(parseInt(yearString), parseInt(monthString) - 1, 1),
  ).toISOString();

  const { status, data: adminReport } = useQuery({
    queryKey: ['adminReport'],
    queryFn: async () => {
      const response = await fetch(`/api/admin/report?date=${rfc3339Date}`, {
        headers: {
          Accept: 'application/json',
        },
      });
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      return response.json();
    },
  });

  if (status != 'success') {
    return <CheckFetchStatus status={status} />;
  }

  if (adminReport.sellers.length === 0 && adminReport.total == 0) {
    console.log('no data');
    return <NotFound />;
  }

  return (
    <div style={{ padding: '10% 10% 0% 10%' }}>
      <div className='flex_wrapper' style={reportPageStyle}>
        <h1 className='title_color' style={{ paddingBottom: '30px' }}>
          <b>
            {adminReport.year} / {adminReport.month} Finical Report
          </b>
        </h1>
        <h4>The top three highest monthly sales</h4>
        <hr className='hr' />
        <Row>
          {adminReport.sellers.map((item: SellersProps, index: number) => {
            return (
              <Col xs={12} md={4} key={index}>
                <div className='center title_color' style={{ paddingTop: '30px' }}>
                  <h4>No. {index + 1}</h4>
                </div>
                <div className='center' style={{ paddingBottom: '30px' }}>
                  TWD $ {item.total_sales}
                </div>

                <SellerItem data={item} />
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
            TWD $ {adminReport.total}
          </Col>
        </Row>
      </div>
    </div>
  );
};

export default AdminReportEach;
