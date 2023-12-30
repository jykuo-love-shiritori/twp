import { Col, Row } from 'react-bootstrap';
import { useQuery } from '@tanstack/react-query';
import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import { useNavigate } from 'react-router-dom';

import HistoryItem, { BuyerHistoryItemProps } from '@components/HistoryItem';

const History = () => {
  const navigate = useNavigate();

  const { status, data: buyerOrderData } = useQuery({
    queryKey: ['buyerOder'],
    queryFn: async () => {
      const response = await fetch(`/api/buyer/order?offset=${0}&limit=${8}`, {
        headers: {
          Accept: 'application/json',
        },
      });
      if (!response.ok) {
        RouteOnNotOK(response, navigate);
      }
      return (await response.json()) as BuyerHistoryItemProps[];
    },
  });

  if (status != 'success') {
    return <CheckFetchStatus status={status} />;
  }

  return (
    <div>
      <div className='title'>Order history</div>
      <hr className='hr' />
      <br />

      <Row>
        {buyerOrderData.map((item, index) => {
          return (
            <Col xs={12} key={index}>
              <HistoryItem data={item} user='buyer' />
            </Col>
          );
        })}
      </Row>
    </div>
  );
};

export default History;
