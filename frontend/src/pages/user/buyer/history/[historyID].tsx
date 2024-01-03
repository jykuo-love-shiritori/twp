import { faFile, faMoneyBill, faTruck, faBox } from '@fortawesome/free-solid-svg-icons';
import { Col, Row } from 'react-bootstrap';
import { useParams } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';
import { useNavigate } from 'react-router-dom';

import HistoryProduct from '@components/HistoryProduct';
import NotFound from '@components/NotFound';
import RecordStatus from '@components/RecordStatus';
import UserItem from '@components/UserItem';
import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import { useAuth } from '@lib/Auth';

interface BuyerOrderProps {
  info: {
    id: number;
    shop_name: string;
    shop_image_url: string;
    shipment: number;
    total_price: number;
    status: 'paid' | 'shipped' | 'delivered';
    created_at: string;
    discount: number;
  };
  details: {
    product_id: number;
    name: string;
    description: string;
    price: number;
    image_url: string;
    quantity: number;
  }[];
}

const BuyerHistoryEach = () => {
  const token = useAuth();
  const navigate = useNavigate();

  const params = useParams<{ history_id?: string }>();
  let order_id: number | undefined;
  const recordStatus: boolean[] = new Array(4).fill(false);

  if (params.history_id) {
    order_id = parseInt(params.history_id);
  }

  const { status, data: buyerOrderData } = useQuery({
    queryKey: ['buyerOrder', order_id],
    queryFn: async () => {
      if (order_id === undefined) {
        throw new Error('Invalid order_id');
      }
      const response = await fetch(`/api/buyer/order/${order_id}`, {
        headers: {
          Accept: 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });
      if (!response.ok) {
        RouteOnNotOK(response, navigate);
      }
      return (await response.json()) as BuyerOrderProps;
    },
  });

  if (status !== 'success') {
    return <CheckFetchStatus status={status} />;
  } else {
    switch (buyerOrderData.info.status) {
      case 'paid':
        recordStatus.fill(true, 0, 1);
        break;
      case 'shipped':
        recordStatus.fill(true, 0, 2);
        break;
      case 'delivered':
        recordStatus.fill(true, 0, 3);
        break;
      default:
        recordStatus.fill(true, 0, 0);
    }
  }

  if (!buyerOrderData) {
    return <NotFound />;
  }
  return (
    <div style={{ padding: '7% 10% 10% 10%' }}>
      <div className='title'>Record ID : {buyerOrderData.info.id} </div>
      <Row>
        <Col xs={6} md={3}>
          <RecordStatus icon={faFile} text='Order placed' status={recordStatus[0]} />
        </Col>
        <Col xs={6} md={3}>
          <RecordStatus icon={faMoneyBill} text='Payment confirmed' status={recordStatus[1]} />
        </Col>
        <Col xs={6} md={3}>
          <RecordStatus icon={faTruck} text='Shipped out' status={recordStatus[2]} />
        </Col>
        <Col xs={6} md={3}>
          <RecordStatus icon={faBox} text='Order received' status={recordStatus[3]} />
        </Col>
      </Row>

      <hr className='hr' />

      <UserItem
        img_path={buyerOrderData.info.shop_image_url}
        name={buyerOrderData.info.shop_name}
      />

      {buyerOrderData.details.map((product, index) => {
        return (
          <HistoryProduct
            data={{
              image_id: product.image_url,
              name: product.name,
              price: Math.floor(product.price),
              quantity: product.quantity,
            }}
            key={index}
          />
        );
      })}

      <Row className='light'>
        <Col xs={12} md={7} />
        <Col xs={6} md={3}>
          Original Total :
        </Col>
        <Col xs={6} md={2}>
          ${' '}
          {Math.floor(
            buyerOrderData.info.total_price +
              buyerOrderData.info.discount -
              buyerOrderData.info.shipment,
          )}
        </Col>

        <Col xs={12} md={7} />
        <Col xs={6} md={3}>
          Shipment :
        </Col>
        <Col xs={6} md={2}>
          $ {Math.floor(buyerOrderData.info.shipment)}
        </Col>

        <Col xs={12} md={7} />
        <Col xs={6} md={3}>
          Coupon :
        </Col>
        <Col xs={6} md={2}>
          -$ {Math.floor(buyerOrderData.info.discount)}
        </Col>
      </Row>
      <hr className='hr' />
      <Row className='light'>
        <Col xs={12} md={7} />
        <Col xs={6} md={3}>
          Order Total :
        </Col>
        <Col xs={6} md={2}>
          $ {Math.floor(buyerOrderData.info.total_price)}
        </Col>
      </Row>
    </div>
  );
};

export default BuyerHistoryEach;
