import { Col, Row } from 'react-bootstrap';
import { useQuery } from '@tanstack/react-query';
import { useNavigate } from 'react-router-dom';

import GoodsItem from '@components/GoodsItem';
import { Props } from '@components/GoodsItem';
import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import { useAuth } from '@lib/Auth';

const Shop = () => {
  const token = useAuth();
  const navigate = useNavigate();

  const { status, data: shopData } = useQuery({
    queryKey: ['shopsView'],
    queryFn: async () => {
      const response = await fetch(`/api/seller/product?offset=${0}&limit=${8}`, {
        headers: {
          Accept: 'application/json',
          Authorization: `Bearer ${token}`,
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
      <div className='title'>All products</div>
      <hr className='hr' />
      <Row>
        {shopData.map((data: Props, index: number) => {
          return (
            <Col xs={6} md={3} key={index}>
              <GoodsItem id={data.id} name={data.name} image_url={data.image_url} />
            </Col>
          );
        })}
      </Row>
    </div>
  );
};

export default Shop;
