import '@style/global.css';

import { Col, Row } from 'react-bootstrap';
import { useQuery } from '@tanstack/react-query';
import { useNavigate } from 'react-router-dom';

import GoodsItem from '@components/GoodsItem';

import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import { GoodsItemProps } from '@components/GoodsItem';

const Discover = () => {
  const navigate = useNavigate();

  const { status, data: goodsData } = useQuery({
    queryKey: ['discover'],
    queryFn: async () => {
      const response = await fetch(`/api/discover?offset=${0}&limit=${20}`, {
        headers: {
          Accept: 'application/json',
        },
      });
      if (!response.ok) {
        RouteOnNotOK(response, navigate);
      }
      return (await response.json()) as GoodsItemProps[];
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
