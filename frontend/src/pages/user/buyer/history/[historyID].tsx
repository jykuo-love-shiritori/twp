import {
  faMoneyBill,
  faTruck,
  faBox,
  faCircleXmark,
  faHandshake,
} from '@fortawesome/free-solid-svg-icons';
import { Col, Row, Modal } from 'react-bootstrap';
import { useParams } from 'react-router-dom';
import { useMutation, useQuery } from '@tanstack/react-query';
import { useNavigate } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { IconProp } from '@fortawesome/fontawesome-svg-core';
import { useState } from 'react';

import HistoryProduct from '@components/HistoryProduct';
import NotFound from '@components/NotFound';
import RecordStatus, { StatusProps } from '@components/RecordStatus';
import UserItem from '@components/UserItem';
import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import { useAuth } from '@lib/Auth';
import TButton from '@components/TButton';

interface BuyerOrderProps {
  info: {
    id: number;
    shop_name: string;
    shop_image_url: string;
    shipment: number;
    total_price: number;
    status: 'paid' | 'shipped' | 'delivered' | 'finished';
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

  const [show, setShow] = useState(false);
  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  let currentStatus = 0;

  let statusContainer: StatusProps[] = [
    { text: 'Order paid', icon: faMoneyBill, status: false },
    { text: 'Shipped out', icon: faTruck, status: false },
    { text: 'Order placed', icon: faBox, status: false },
    { text: 'Order finished', icon: faHandshake, status: false },
  ];

  const ChangeStatusContainer = () => {
    return statusContainer.map((statusObj, index) => {
      return index <= currentStatus ? { ...statusObj, status: true } : statusObj;
    });
  };

  const params = useParams<{ history_id?: string }>();
  let order_id: number | undefined;

  if (params.history_id) {
    order_id = parseInt(params.history_id);
  }

  const updateStatus = useMutation({
    mutationFn: async (id: number) => {
      const data = { status: 'finished' };
      const response = await fetch(`/api/buyer/order/${id}`, {
        method: 'PATCH',
        headers: {
          Accept: 'application/json',
          'Content-Type': 'application/json',
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(data),
        redirect: 'follow',
      });
      if (!response.ok) {
        throw new Error('change status failed');
      }
      return await response.json();
    },
    onSuccess: () => {
      refetch();
    },
  });

  const {
    status,
    data: buyerOrderData,
    refetch,
  } = useQuery({
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
        currentStatus = 0;
        break;
      case 'shipped':
        currentStatus = 1;
        break;
      case 'delivered':
        currentStatus = 2;
        break;
      case 'finished':
        currentStatus = 3;
        break;
    }
    statusContainer = ChangeStatusContainer();
  }

  if (!buyerOrderData) {
    return <NotFound />;
  }
  return (
    <div style={{ padding: '7% 10% 10% 10%' }}>
      <div className='title'>Record ID : {buyerOrderData.info.id} </div>
      <Row>
        <Col xs={6} md={3}>
          <RecordStatus
            icon={statusContainer[0].icon}
            text={statusContainer[0].text}
            status={statusContainer[0].status}
          />
        </Col>
        <Col xs={6} md={3}>
          <RecordStatus
            icon={statusContainer[1].icon}
            text={statusContainer[1].text}
            status={statusContainer[1].status}
          />
        </Col>
        <Col xs={6} md={3}>
          <RecordStatus
            icon={statusContainer[2].icon}
            text={statusContainer[2].text}
            status={statusContainer[2].status}
          />
        </Col>
        <Col xs={6} md={3}>
          <RecordStatus
            icon={statusContainer[3].icon}
            text={statusContainer[3].text}
            status={statusContainer[3].status}
          />
        </Col>
      </Row>

      <hr className='hr' />

      <Row>
        <Col xs={6} md={9}>
          <UserItem
            img_path={buyerOrderData.info.shop_image_url}
            name={buyerOrderData.info.shop_name}
          />
        </Col>
        <Col xs={6} md={3}>
          <TButton text='Update status' action={handleShow} />
        </Col>
      </Row>

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

      <Modal show={show} onHide={handleClose} className='coupon_modal'>
        <Modal.Header style={{ paddingTop: '30px' }}>
          <Modal.Title>Update status</Modal.Title>
          <div className='right' style={{ cursor: 'pointer' }}>
            <FontAwesomeIcon icon={faCircleXmark as IconProp} size='2x' onClick={handleClose} />
          </div>
        </Modal.Header>
        <Modal.Body>
          {currentStatus === 3 ? (
            <div>
              <div className='center' style={{ padding: '20px 0 20px 0' }}>
                <FontAwesomeIcon icon={faHandshake} size='4x' />
              </div>
              <div className='center'>
                <h4 className='title_color'>
                  <b>Order finished 🎉</b>
                </h4>
              </div>
            </div>
          ) : (
            <div>
              {currentStatus !== 2 ? (
                <div>
                  <RecordStatus
                    icon={statusContainer[currentStatus].icon}
                    text={statusContainer[currentStatus].text}
                    status={statusContainer[currentStatus].status}
                  />
                  <br />
                  <div className='center'>Your order is not arrived yet.</div>
                </div>
              ) : (
                <div>
                  <div className='center'>
                    Your order is arrived, you can confirm it as finished!
                  </div>

                  <div className='center' style={{ paddingTop: '20px' }}>
                    <div style={{ width: '50%' }}>
                      <TButton
                        text='Confirm order'
                        action={() => updateStatus.mutate(buyerOrderData.info.id)}
                      />
                    </div>
                  </div>
                </div>
              )}
            </div>
          )}
        </Modal.Body>
      </Modal>
    </div>
  );
};

export default BuyerHistoryEach;
