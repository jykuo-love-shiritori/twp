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

interface SellerOrderProps {
  order_info: {
    id: number;
    shipment: number;
    total_price: number;
    status: 'paid' | 'shipped' | 'delivered';
    created_at: string;
    user_id: number;
    user_name: string;
    user_image_url: string;
    discount: number;
  };
  products: {
    _id: number;
    name: string;
    description: string;
    price: number;
    image_url: string;
    quantity: number;
  }[];
}

const SellerHistoryEach = () => {
  const navigate = useNavigate();

  const params = useParams<{ history_id?: string }>();
  let order_id: number | undefined;
  // eslint-disable-next-line prefer-const
  let recordStatus: boolean[] = new Array(4).fill(false);

  if (params.history_id) {
    order_id = parseInt(params.history_id);
  }

  const { status, data: sellerOrderData } = useQuery({
    queryKey: ['sellerOrder', order_id],
    queryFn: async () => {
      if (order_id === undefined) {
        throw new Error('Invalid order_id');
      }
      const response = await fetch(`/api/seller/order/${order_id}`, {
        headers: {
          Accept: 'application/json',
        },
      });
      if (!response.ok) {
        RouteOnNotOK(response, navigate);
      }
      return (await response.json()) as SellerOrderProps;
    },
  });

  if (status != 'success') {
    return <CheckFetchStatus status={status} />;
  } else {
    switch (sellerOrderData.order_info.status) {
      case 'paid':
        recordStatus.fill(true, 0, 1);
        break;
      case 'shipped':
        recordStatus.fill(true, 0, 2);
        break;
      case 'delivered':
        recordStatus.fill(true, 0, 3);
        break;
    }
  }

  if (!sellerOrderData) {
    return <NotFound />;
  }

  const originalTotalPrice = sellerOrderData.products.reduce(
    (sum, product) => sum + product.price * product.quantity,
    0,
  );

  return (
    <div style={{ padding: '7% 10% 10% 10%' }}>
      <div className='title'>Record ID : {sellerOrderData.order_info.id} </div>
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
        img_path={sellerOrderData.order_info.user_image_url}
        name={sellerOrderData.order_info.user_name}
      />

      {sellerOrderData.products.map((product, index) => {
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
          {Math.floor(originalTotalPrice)}
        </Col>

        <Col xs={12} md={7} />
        <Col xs={6} md={3}>
          Shipment :
        </Col>
        <Col xs={6} md={2}>
          ${Math.floor(sellerOrderData.order_info.shipment)}
        </Col>

        <Col xs={12} md={7} />
        <Col xs={6} md={3}>
          Coupon :
        </Col>
        <Col xs={6} md={2}>
          $
          {Math.floor(
            originalTotalPrice -
              sellerOrderData.order_info.total_price -
              sellerOrderData.order_info.shipment,
          )}
        </Col>
      </Row>
      <hr className='hr' />
      <Row className='light'>
        <Col xs={12} md={7} />
        <Col xs={6} md={3}>
          Order Total :
        </Col>
        <Col xs={6} md={2}>
          ${Math.floor(sellerOrderData.order_info.total_price)}
        </Col>
      </Row>
    </div>
  );
};

export default SellerHistoryEach;
