import { Col, Row } from 'react-bootstrap';
import { useQuery } from '@tanstack/react-query';
import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import { useNavigate } from 'react-router-dom';

import { useAuth } from '@lib/Auth';

import HistoryItem, { BuyerHistoryItemProps } from '@components/HistoryItem';

const History = () => {
  const token = useAuth();
  const navigate = useNavigate();

  const { status, data: buyerOrderData } = useQuery({
    queryKey: ['buyerOder'],
    queryFn: async () => {
      const response = await fetch(`/api/buyer/order?offset=0&limit=20`, {
        headers: {
          Accept: 'application/json',
          Authorization: `Bearer ${token}`,
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
          const data: BuyerHistoryItemProps = item;
          data.user = 'buyer';
          return (
            <Col xs={12} key={index}>
              <HistoryItem data={data} />
            </Col>
          );
        })}
      </Row>
    </div>
  );
};

export default History;
