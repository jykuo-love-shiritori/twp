import { Col, Row } from 'react-bootstrap';
import { useQuery } from '@tanstack/react-query';
import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import { useNavigate } from 'react-router-dom';

import { useAuth } from '@lib/Auth';

import HistoryItem, { SellerHistoryItemProps } from '@components/HistoryItem';

const SellerShipment = () => {
  const token = useAuth();
  const navigate = useNavigate();

  const { status, data: sellerOrderData } = useQuery({
    queryKey: ['sellerOder'],
    queryFn: async () => {
      const response = await fetch(`/api/seller/order?offset=0&limit=8`, {
        headers: {
          Accept: 'application/json',
          Authorization: `Bearer ${token}`,
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
          const data: SellerHistoryItemProps = item;
          data.user = 'seller';
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

export default SellerShipment;
