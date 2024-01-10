import { Col, Row } from 'react-bootstrap';
import { useQuery } from '@tanstack/react-query';
import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { useState } from 'react';

import { useAuth } from '@lib/Auth';

import HistoryItem, { SellerHistoryItemProps } from '@components/HistoryItem';
import Pagination from '@components/Pagination';

const SellerShipment = () => {
  const token = useAuth();
  const navigate = useNavigate();
  const [searchParams, setSearchParams] = useSearchParams();
  const [isMore, setIsMore] = useState(true);

  const itemLimit = 4;

  if (!searchParams.has('offset') || Number(searchParams.get('limit')) !== itemLimit + 1) {
    const newSearchParams = new URLSearchParams({
      offset: '0',
      limit: (itemLimit + 1).toString(),
    });
    setSearchParams(newSearchParams, { replace: true });
  }

  const { status, data: sellerOrderData } = useQuery({
    queryKey: ['sellerOder', searchParams.toString()],
    queryFn: async () => {
      const resp = await fetch(`/api/seller/order?` + searchParams.toString(), {
        headers: {
          Accept: 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });
      if (!resp.ok) {
        RouteOnNotOK(resp, navigate);
        return [];
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
    select: (data) => data as SellerHistoryItemProps[],
    enabled: true,
    refetchOnWindowFocus: false,
  });

  if (status != 'success') {
    return <CheckFetchStatus status={status} />;
  }

  return (
    <div>
      <div className='title'>All Shipments</div>
      <hr className='hr' />
      <br />

      <Row>
        {sellerOrderData.map((item, index) => {
          const data: SellerHistoryItemProps = item;
          data.user = 'seller';
          return (
            <Col xs={12} key={index}>
              <HistoryItem data={data} />
            </Col>
          );
        })}
      </Row>
      <div className='center'>
        <Pagination limit={itemLimit} isMore={isMore} />
      </div>
    </div>
  );
};

export default SellerShipment;
