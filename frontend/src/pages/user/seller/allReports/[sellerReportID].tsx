import '@style/global.css';

import { Col, Row } from 'react-bootstrap';
import { useParams } from 'react-router-dom';

import NotFound from '@components/NotFound';
import GoodsItem from '@components/GoodsItem';

import reportData from '@pages/user/seller/sellerReportData.json';

interface Goods {
  id: number;
  name: string;
  imgUrl: string;
  amount: number;
}

interface Props {
  year: number;
  month: number;
  goods: Goods[];
  totalAmount: number;
  id: number;
}

const SellerReportEach = () => {
  const reportPageStyle = {
    borderRadius: '50px 50px 0px 0px',
    border: '1px solid var(--button_border)',
    background: 'var(--bg)',
    boxShadow: '0px 4px 30px 2px var(--title)',
    padding: ' 7% 12% 10% 12%',
  };

  const params = useParams();

  let data: Props | undefined;
  reportData.findIndex((item) => {
    if (item.id.toString() === params.report_id) {
      data = item;
    }
  });

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
            {data.goods.map((item, index) => {
              return (
                <Col xs={12} md={4} key={index}>
                  <GoodsItem id={item.id} name={item.name} imgUrl={item.imgUrl} />
                  <div className='center title_color'>
                    <h4>No. {index + 1}</h4>
                  </div>
                  <div className='center'>TWD $ {item.amount}</div>
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

export default SellerReportEach;
