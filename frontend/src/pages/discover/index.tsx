import '@style/global.css';

import { Col, Row } from 'react-bootstrap';
import { useQuery } from '@tanstack/react-query';
import { useNavigate, useSearchParams } from 'react-router-dom';

import GoodsItem from '@components/GoodsItem';

import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import { GoodsItemProps } from '@components/GoodsItem';
import { useState } from 'react';
import Pagination from '@components/Pagination';

const Discover = () => {
  const navigate = useNavigate();
  const [searchParams, setSearchParams] = useSearchParams();
  const [isMore, setIsMore] = useState(true);

  const itemLimit = 12;

  if (!searchParams.has('offset') || Number(searchParams.get('limit')) !== itemLimit + 1) {
    const newSearchParams = new URLSearchParams({
      offset: '0',
      limit: (itemLimit + 1).toString(),
    });
    setSearchParams(newSearchParams, { replace: true });
  }

  const {
    status,
    data: goodsData,
    refetch,
  } = useQuery({
    queryKey: ['discover', searchParams.toString()],
    queryFn: async () => {
      const resp = await fetch(`/api/discover?` + searchParams.toString(), {
        headers: {
          Accept: 'application/json',
        },
      });
      if (!resp.ok) {
        RouteOnNotOK(resp, navigate);
      }
      const response = await resp.json();
      if (response.length === itemLimit + 1) {
        setIsMore(true);
        response.pop();
      } else {
        setIsMore(false);
      }
      return response;
    },
    select: (data) => data as GoodsItemProps[],
    enabled: true,
    refetchOnWindowFocus: false,
  });

  if (status != 'success') {
    return <CheckFetchStatus status={status} />;
  }

  return (
    <div style={{ padding: '10% 10% 2% 10%' }}>
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
      <div className='center'>
        <Pagination
          searchParams={searchParams}
          setSearchParams={setSearchParams}
          refetch={refetch}
          limit={itemLimit}
          isMore={isMore}
        />
      </div>
    </div>
  );
};

export default Discover;
