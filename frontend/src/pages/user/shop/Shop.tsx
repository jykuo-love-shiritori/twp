import { Col, Row } from 'react-bootstrap';
import { useQuery } from '@tanstack/react-query';

import GoodsItem from '@components/GoodsItem';
import { Props } from '@components/GoodsItem';
import { CheckFetchStatus } from '@lib/Status';

const Shop = () => {
  const { status, data: shopData } = useQuery({
    queryKey: ['shopsView'],
    queryFn: async () => {
      const response = await fetch(`/api/seller/product?offset=${0}&limit=${8}`, {
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
