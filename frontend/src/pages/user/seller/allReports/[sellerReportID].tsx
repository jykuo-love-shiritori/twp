import '@style/global.css';

import { Col, Row } from 'react-bootstrap';
import { useParams } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';
import { useNavigate } from 'react-router-dom';

import NotFound from '@components/NotFound';
import GoodsItem from '@components/GoodsItem';

import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';

interface GoodsProps {
  product_id: number;
  name: string;
  price: number;
  image_url: string;
  total_quantity: number;
  total_sell: number;
  order_count: number;
}

interface SellerReportProps {
  products: GoodsProps[];
  report: {
    total_income: number;
    order_count: number;
  };
}

const reportPageStyle = {
  borderRadius: '50px 50px 0px 0px',
  border: '1px solid var(--button_border)',
  background: 'var(--bg)',
  boxShadow: '0px 4px 30px 2px var(--title)',
  padding: ' 7% 12% 10% 12%',
};

const SellerReportEach = () => {
  const navigate = useNavigate();

  const { year, month } = useParams();
  const yearString = year ?? '';
  const monthString = month ?? '';

  const rfc3339Date = new Date(
    Date.UTC(parseInt(yearString), parseInt(monthString) - 1, 1),
  ).toISOString();

  const { status, data: sellerReport } = useQuery({
    queryKey: ['sellerReport', year, month],
    queryFn: async () => {
      const response = await fetch(`/api/seller/report?time=${rfc3339Date}`, {
        headers: {
          Accept: 'application/json',
        },
      });
      if (!response.ok) {
        RouteOnNotOK(response, navigate);
      }
      return (await response.json()) as SellerReportProps;
    },
  });

  if (status != 'success') {
    return <CheckFetchStatus status={status} />;
  }

  if (sellerReport.products.length === 0 && sellerReport.report.order_count == 0) {
    console.log('no data');
    return <NotFound />;
  }

  return (
    <div style={{ padding: '10% 10% 0% 10%' }}>
      <div className='flex_wrapper' style={reportPageStyle}>
        <h1 className='title_color' style={{ paddingBottom: '30px' }}>
          <b>
            {year} / {month} Finical Report
          </b>
        </h1>
        <h4>The top three highest monthly sales</h4>
        <hr className='hr' />
        <Row>
          {sellerReport.products.map((item, index) => {
            return (
              <Col xs={12} md={4} key={index}>
                <div className='center title_color' style={{ paddingTop: '30px' }}>
                  <h4>No. {index + 1}</h4>
                </div>
                <div className='center' style={{ paddingBottom: '30px' }}>
                  TWD $ {item.total_sell}
                </div>

                <GoodsItem id={item.product_id} name={item.name} image_url={item.image_url} />
              </Col>
            );
          })}
        </Row>

        <hr className='hr' />

        <Row>
          <Col xs={12} className='center_vertical'>
            <h4 style={{ paddingBottom: '10px' }}>
              <b>Total sales revenue :</b>
            </h4>
          </Col>
          <Col xs={12} className='center_vertical'>
            TWD ${sellerReport.report.total_income}
          </Col>
        </Row>
      </div>
    </div>
  );
};

export default SellerReportEach;
