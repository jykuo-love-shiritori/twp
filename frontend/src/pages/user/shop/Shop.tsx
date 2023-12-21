import { Col, Row } from 'react-bootstrap';
import { useQuery } from '@tanstack/react-query';

import GoodsItem from '@components/GoodsItem';
import { Props } from '@components/GoodsItem';

const Shop = () => {
  const { isLoading, isError, data } = useQuery({
    queryKey: ['shopsView'],
    queryFn: async () => {
      const response = await fetch(`/api/seller/product?offset=${0}&limit=${8}`, {
        headers: {
          'Content-Type': 'application/json',
        },
      });
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      return response.json();
    },
  });

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (isError) {
    return <div>Error</div>;
  }

  console.log(data);

  return (
    <div>
      <div className='title'>All products</div>
      <hr className='hr' />
      <Row>
        {data.map((d: Props, index: number) => {
          return (
            <Col xs={6} md={3} key={index}>
              <GoodsItem id={d.id} name={d.name} image_url={d.image_url} />
            </Col>
          );
        })}
      </Row>
    </div>
  );
};

export default Shop;
