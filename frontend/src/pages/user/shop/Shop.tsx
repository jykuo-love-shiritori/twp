import { Col, Row } from 'react-bootstrap';

import GoodsItem from '@components/GoodsItem';

import goodsData from '@pages/discover/goodsData.json';

const Shop = () => {
  return (
    <div>
      <div className='title'>All products</div>
      <hr className='hr' />
      <Row>
        {goodsData.map((data, index) => {
          return (
            <Col xs={6} md={3} key={index}>
              <GoodsItem id={data.id} name={data.name} imgUrl={data.imgUrl} />
            </Col>
          );
        })}
      </Row>
    </div>
  );
};

export default Shop;
