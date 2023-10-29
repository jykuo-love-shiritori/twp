import '@style/global.css';

import { Col, Row } from 'react-bootstrap';

import GoodsItem from '@components/GoodsItem';

import goodsData from '@pages/discover/goodsData.json';

const Discover = () => {
  return (
    <div style={{ padding: '10%' }}>
      <span className='title'>Discover</span>

      <div style={{ padding: '2% 4% 2% 4%' }}>
        <Row>
          {goodsData.map((data) => {
            return (
              <Col xs={6} md={3}>
                <GoodsItem id={data.id} name={data.name} imgUrl={data.imgUrl} isIndex={false} />
              </Col>
            );
          })}
        </Row>
      </div>
    </div>
  );
};

export default Discover;
