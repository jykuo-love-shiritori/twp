import {
  faMoneyBill,
  faTruck,
  faBox,
  faHandshake,
  faCircleXmark,
} from '@fortawesome/free-solid-svg-icons';
import { Col, Row, Modal } from 'react-bootstrap';
import { useParams } from 'react-router-dom';
import { useMutation, useQuery } from '@tanstack/react-query';
import { useNavigate } from 'react-router-dom';
import { useState } from 'react';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { IconProp } from '@fortawesome/fontawesome-svg-core';

import HistoryProduct from '@components/HistoryProduct';
import NotFound from '@components/NotFound';
import RecordStatus, { StatusProps } from '@components/RecordStatus';
import UserItem from '@components/UserItem';
import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import { useAuth } from '@lib/Auth';
import TButton from '@components/TButton';

interface SellerOrderProps {
  order_info: {
    id: number;
    shipment: number;
    total_price: number;
    status: 'paid' | 'shipped' | 'delivered' | 'finished';
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

interface PatchProps {
  id: number;
  current_status: 'paid' | 'shipped' | 'delivered' | 'finished';
}

const SellerHistoryEach = () => {
  const token = useAuth();
  const navigate = useNavigate();

  const [show, setShow] = useState(false);
  const handleClose = () => setShow(false);
  const handleShow = () => setShow(true);

  let currentStatus = 0;

  const statusText: ('paid' | 'shipped' | 'delivered' | 'finished')[] = [
    'paid',
    'shipped',
    'delivered',
    'finished',
  ];

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
    mutationFn: async (data: PatchProps) => {
      const response = await fetch(`/api/seller/order/${data.id}`, {
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
    data: sellerOrderData,
    refetch,
  } = useQuery({
    queryKey: ['sellerOrder', order_id],
    queryFn: async () => {
      if (order_id === undefined) {
        throw new Error('Invalid order_id');
      }
      const response = await fetch(`/api/seller/order/${order_id}`, {
        headers: {
          Accept: 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });
      if (!response.ok) {
        RouteOnNotOK(response, navigate);
      }
      return (await response.json()) as SellerOrderProps;
    },
  });

  if (status !== 'success') {
    return <CheckFetchStatus status={status} />;
  } else {
    switch (sellerOrderData.order_info.status) {
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
            img_path={sellerOrderData.order_info.user_image_url}
            name={sellerOrderData.order_info.user_name}
          />
        </Col>
        <Col xs={6} md={3}>
          <TButton text='Update status' action={handleShow} />
        </Col>
      </Row>

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
          $ {Math.floor(originalTotalPrice)}
        </Col>

        <Col xs={12} md={7} />
        <Col xs={6} md={3}>
          Shipment :
        </Col>
        <Col xs={6} md={2}>
          $ {Math.floor(sellerOrderData.order_info.shipment)}
        </Col>

        <Col xs={12} md={7} />
        <Col xs={6} md={3}>
          Coupon :
        </Col>
        <Col xs={6} md={2}>
          -${' '}
          {Math.floor(
            originalTotalPrice -
              sellerOrderData.order_info.total_price +
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
          $ {Math.floor(sellerOrderData.order_info.total_price)}
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
                  <Row>
                    <Col xs={6}>
                      <RecordStatus
                        icon={statusContainer[currentStatus].icon}
                        text={statusContainer[currentStatus].text}
                        status={statusContainer[currentStatus].status}
                      />
                    </Col>
                    <Col xs={6}>
                      <RecordStatus
                        icon={statusContainer[currentStatus + 1].icon}
                        text={statusContainer[currentStatus + 1].text}
                        status={true}
                      />
                    </Col>
                    <Col xs={12} className='center' style={{ paddingTop: '20px' }}>
                      Ready from
                      <span className='title_color' style={{ padding: '0 10px 0 10px' }}>
                        <b>{statusContainer[currentStatus].text}</b>
                      </span>
                      →
                      <span className='title_color' style={{ padding: '0 10px 0 10px' }}>
                        <b>{statusContainer[currentStatus + 1].text}</b>
                      </span>
                    </Col>
                  </Row>
                  <div className='center' style={{ paddingTop: '20px' }}>
                    <div style={{ width: '30%' }}>
                      <TButton
                        text='Confirm'
                        action={() =>
                          updateStatus.mutate({
                            id: sellerOrderData.order_info.id,
                            current_status: statusText[currentStatus],
                          })
                        }
                      />
                    </div>
                  </div>
                </div>
              ) : (
                <div className='center'>
                  The order is arrived, please wait for buyer to confirm!
                  <div className='center' style={{ paddingTop: '20px' }}>
                    <div style={{ width: '30%' }}></div>
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

export default SellerHistoryEach;
