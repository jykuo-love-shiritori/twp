import { Col, Row } from 'react-bootstrap';

import SellerGoodsItem from '@components/SellerGoodsItem';
import TButton from '@components/TButton';

import goodsData from '@pages/discover/goodsData.json';

const Products = () => {
  return (
    <div>
      <Row>
        <Col sm={12} md={8}>
          <div className='title'>All products</div>
        </Col>
        <Col sm={12} md={4}>
          <div style={{ padding: '20px 0 0 0' }}>
            <TButton text='Add New Item' action='/user/seller/manageProducts/new' />
          </div>
        </Col>
      </Row>
      <hr className='hr' />
      <Row>
        {goodsData.map((data, index) => {
          return (
            <Col xs={6} md={3} key={index}>
              <SellerGoodsItem
                id={data.id}
                name={data.name}
                image_url={data.image_url}
                isIndex={false}
              />
            </Col>
          );
        })}
      </Row>
    </div>
  );
};

export default Products;
