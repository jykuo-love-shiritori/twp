import { Col, Row } from 'react-bootstrap';
import { useQuery } from '@tanstack/react-query';

import SellerGoodsItem from '@components/SellerGoodsItem';
import TButton from '@components/TButton';
import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import { GoodsItemProps } from '@components/GoodsItem';
import { useNavigate } from 'react-router-dom';

const Products = () => {
  const navigate = useNavigate();

  const { status, data: sellerShopData } = useQuery({
    queryKey: ['sellerShopView'],
    queryFn: async () => {
      const response = await fetch(`/api/seller/product?offset=${0}&limit=${8}`, {
        headers: {
          Accept: 'application/json',
        },
      });
      if (!response.ok) {
        RouteOnNotOK(response, navigate);
      }
      return await response.json();
    },
  });

  if (status != 'success') {
    return <CheckFetchStatus status={status} />;
  }

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
        {sellerShopData.map((data: GoodsItemProps, index: number) => {
          return (
            <Col xs={6} md={3} key={index}>
              <SellerGoodsItem id={data.id} name={data.name} image_url={data.image_url} />
            </Col>
          );
        })}
      </Row>
    </div>
  );
};

export default Products;
