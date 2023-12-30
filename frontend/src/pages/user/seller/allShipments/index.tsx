import { Col, Row } from 'react-bootstrap';
import { useQuery } from '@tanstack/react-query';
import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import { useNavigate } from 'react-router-dom';

import HistoryItem, { SellerHistoryItemProps } from '@components/HistoryItem';

const SellerShipment = () => {
  const navigate = useNavigate();

  const { status, data: sellerOrderData } = useQuery({
    queryKey: ['sellerOder'],
    queryFn: async () => {
      const response = await fetch(`/api/seller/order?offset=${0}&limit=${8}`, {
        headers: {
          Accept: 'application/json',
        },
      });
      if (!response.ok) {
        RouteOnNotOK(response, navigate);
      }
      return (await response.json()) as SellerHistoryItemProps[];
    },
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
          return (
            <Col xs={12} key={index}>
              <HistoryItem data={item} user='seller' />
            </Col>
          );
        })}
      </Row>
    </div>
  );
};

export default SellerShipment;
