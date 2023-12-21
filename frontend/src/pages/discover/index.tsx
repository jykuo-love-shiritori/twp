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
          {goodsData.map((data, index) => {
            return (
              <Col xs={6} md={3} key={index}>
                <GoodsItem id={data.id} name={data.name} image_url={data.image_url} />
              </Col>
            );
          })}
        </Row>
      </div>
    </div>
  );
};

export default Discover;
