import '@style/global.css';

import { Col, Row } from 'react-bootstrap';
import { useQuery } from '@tanstack/react-query';

import GoodsItem from '@components/GoodsItem';

import { CheckFetchStatus } from '@lib/Status';
import { GoodsItemProps } from '@components/GoodsItem';

const Discover = () => {
  const { status, data: goodsData } = useQuery({
    queryKey: ['discover'],
    queryFn: async () => {
      const response = await fetch(`/api/discover?offset=${0}&limit=${12}`, {
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
    <div style={{ padding: '10%' }}>
      <span className='title'>Discover</span>

      <div style={{ padding: '2% 4% 2% 4%' }}>
        <Row>
          {goodsData.map((data: GoodsItemProps, index: number) => {
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
