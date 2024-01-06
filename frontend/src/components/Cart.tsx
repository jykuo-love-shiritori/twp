import { Row, Col, Offcanvas, Modal } from 'react-bootstrap';
import { useState } from 'react';
import { useQuery } from '@tanstack/react-query';
import { useForm } from 'react-hook-form';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCircleXmark } from '@fortawesome/free-regular-svg-icons';
import { IconProp } from '@fortawesome/fontawesome-svg-core';
import { useNavigate } from 'react-router-dom';
import { useAuth } from '@lib/Auth';
import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import { formatDate } from '@lib/Functions';
import CartProduct from '@components/CartProduct';
import UserItem from '@components/UserItem';
import TButton from '@components/TButton';
import CouponItemTemplate from '@components/CouponItemTemplate';
import CheckoutItem from '@components/CheckoutItem';
import CheckoutItemCoupon from '@components/CheckoutItemCoupon';

interface IProduct {
  enabled: true;
  image_url: string;
  name: string;
  price: number;
  product_id: number;
  quantity: number;
  stock: number;
}

interface Props {
  cartInfo: {
    id: number;
    seller_name: string;
    shop_image_url: string;
    shop_name: string;
  };
  products: IProduct[];
  refresh: () => void;
}

interface ICheckoutCoupon {
  description: string;
  discount: number;
  discount_value: number;
  id: number;
  name: string;
  scope: 'global' | 'shop';
  type: 'percentage' | 'fixed' | 'shipping';
}

interface ICreditCard {
  CVV: string;
  name: string;
  card_number: string;
  expiry_date: string;
}

interface ICheckout {
  coupons: ICheckoutCoupon[];
  payments: ICreditCard[];
  shipment: number;
  subtotal: number;
  total: number;
  total_discount: number;
}

interface IUsableCoupon {
  description: string;
  discount: number;
  expire_date: string;
  id: number;
  name: string;
  scope: 'global' | 'shop';
  type: 'percentage' | 'fixed' | 'shipping';
}

const LabelStyle = {
  fontSize: '24px',
  fontWeight: '700',
  margin: '4% 0',
};

const ContentStyle = {
  fontSize: '18px',
  fontWeight: '500',
  margin: '1% 4%',
};

const Cart = ({ products, cartInfo, refresh }: Props) => {
  const navigate = useNavigate();
  const token = useAuth();
  const [canvaShow, setCanvaShow] = useState(false);
  const [modalShow, setModalShow] = useState(false);
  console.log(cartInfo.seller_name);
  // get the checkout detail
  const {
    data: checkoutData,
    status: checkoutStatus,
    refetch: refetchCheckout,
  } = useQuery({
    queryKey: ['buyerGetCheckout', cartInfo.id],
    queryFn: async () => {
      const resp = await fetch(`/api/buyer/cart/${cartInfo.id}/checkout`, {
        method: 'GET',
        headers: {
          Authorization: `Bearer ${token}`,
          Accept: 'application/json',
        },
      });
      RouteOnNotOK(resp, navigate);
      return (await resp.json()) as ICheckout;
    },
    refetchOnWindowFocus: false,
    enabled: true,
  });

  // get the usable coupons
  const {
    data: usableCouponData,
    status: usableCouponStatus,
    refetch: refetchUsableCoupon,
  } = useQuery({
    queryKey: ['BuyerGetUsableCoupon', cartInfo.id],
    queryFn: async () => {
      const response = await fetch(`/api/buyer/cart/${cartInfo.id}/coupon`, {
        method: 'GET',
        headers: {
          Authorization: `Bearer ${token}`,
          Accept: 'application/json',
        },
      });
      RouteOnNotOK(response, navigate);
      return (await response.json()) as IUsableCoupon[];
    },
    refetchOnWindowFocus: false,
    enabled: true,
  });

  const onViewCheckout = () => {
    console.log('refetch checkout');
    refetchCheckout();
    setCanvaShow(true);
  };
  const onChooseCoupon = () => {
    console.log('refetch usable coupon');
    refetchUsableCoupon();
    setModalShow(true);
  };
  const onPay = () => {
    // TODO: Buyer checkout
    // POST /buyer/cart/:cart_id/checkout
    if (watch('card_id') === null) {
      // TODO: implement real card selection
      console.log('please select a card');
      return;
    }
    console.log('pay');
    refresh();
    setCanvaShow(false);
  };
  const onApplyCoupon = async (coupon_id: number) => {
    console.log(`apply coupon ${coupon_id}`);
    const resp = await fetch(`/api/buyer/cart/${cartInfo.id}/coupon/${coupon_id}`, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token}`,
        Accept: 'application/json',
      },
    });
    if (!resp.ok) {
      // TODO: remove navigate
      RouteOnNotOK(resp, navigate);
    } else {
      // refresh();
      refetchCheckout();
      setModalShow(false);
    }
  };
  const onRemoveCoupon = async (coupon_id: number) => {
    console.log(`remove coupon ${coupon_id}`);
    const resp = await fetch(`/api/buyer/cart/${cartInfo.id}/coupon/${coupon_id}`, {
      method: 'DELETE',
      headers: {
        Authorization: `Bearer ${token}`,
        Accept: 'application/json',
      },
    });
    if (!resp.ok) {
      // TODO: remove navigate
      RouteOnNotOK(resp, navigate);
    } else {
      // refresh();
      refetchCheckout();
    }
  };

  interface FormProps {
    card_id: number;
  }
  const { register, watch } = useForm<FormProps>();

  if (checkoutStatus !== 'success') {
    return <CheckFetchStatus status={checkoutStatus} />;
  }
  if (usableCouponStatus !== 'success') {
    return <CheckFetchStatus status={usableCouponStatus} />;
  }

  return (
    <>
      {/* single cart */}
      <div className='cart_group'>
        <div className='disappear_phone' style={{ fontSize: '20px' }}>
          <Row className='center_vertical' style={{ width: '100%', padding: '0 2%' }}>
            <Col md={6}>
              <UserItem img_path={cartInfo.shop_image_url} name={cartInfo.shop_name} />
            </Col>
            <Col md={3} className='center'>
              Subtotal: ${products.reduce((acc, cur) => acc + cur.price * cur.quantity, 0)}
            </Col>
            <Col md={3} className='center'>
              <TButton text='Checkout' action={onViewCheckout} />
            </Col>
          </Row>
        </div>

        <div className='disappear_tablet disappear_desktop'>
          <Row className='center_vertical' style={{ width: '100%', padding: '0 3%', margin: '0' }}>
            <Col xs={6} style={{ padding: '0 0 0 5%' }}>
              <UserItem img_path={cartInfo.shop_image_url} name={cartInfo.shop_name} />
            </Col>
            <Col xs={6} className='right'>
              Subtotal: ${products.reduce((acc, cur) => acc + cur.price * cur.quantity, 0)}
            </Col>
            <Col xs={12} className='center'>
              <TButton text='Checkout' action={onViewCheckout} />
            </Col>
          </Row>
        </div>

        {products.map((product, index) => (
          <CartProduct data={product} cart_id={cartInfo.id} key={index} refresh={refresh} />
        ))}
      </div>

      {/* checkout canvas */}
      <div>
        <Offcanvas
          show={canvaShow}
          onHide={() => setCanvaShow(false)}
          placement='end'
          style={{
            width: 'min(600px, 90vw)',
            backgroundColor: 'rgba(217, 217, 217, 0.8)',
            backdropFilter: 'blur(15px)',
            border: 'none',
            padding: '2% 4%',
          }}
        >
          <Offcanvas.Header closeButton>
            <Offcanvas.Title>
              <div className='title' style={{ color: 'var(--layout)' }}>
                Checkout
              </div>
            </Offcanvas.Title>
          </Offcanvas.Header>
          <hr style={{ border: '2px solid black', borderRadius: '10px' }} />
          <Offcanvas.Body style={{ padding: '3% 6%' }}>
            <Row style={{ margin: '0' }}>
              <Col xs={12} style={LabelStyle}>
                Cart
                {products.map((productData, index) => (
                  <CheckoutItem
                    label={`${productData.name} x ${productData.quantity}`}
                    value={`${productData.price * productData.quantity} NTD`}
                    key={index}
                  />
                ))}
                <CheckoutItem
                  label={'Subtotal'}
                  value={`${checkoutData.subtotal} NTD`}
                  style={{ fontWeight: '700' }}
                />
              </Col>

              <Col xs={12} style={LabelStyle}>
                Delivery
                <CheckoutItem label={'Shipment'} value={`${checkoutData.shipment} NTD`} />{' '}
              </Col>

              <Col xs={12} style={LabelStyle}>
                <Row>
                  <Col xs='auto' style={{ paddingRight: '0' }}>
                    Discount
                  </Col>
                </Row>
                {checkoutData.coupons.map((couponData, index) => (
                  <CheckoutItemCoupon
                    coupon={couponData}
                    onClick={() => onRemoveCoupon(couponData.id)}
                    key={index}
                  />
                ))}
                <CheckoutItemCoupon onClick={onChooseCoupon} isAddMore={true} />
              </Col>

              <Col xs={12} style={LabelStyle}>
                Summary
                <CheckoutItem label={'Subtotal'} value={`${checkoutData.subtotal} NTD`} />
                <CheckoutItem label={'Shipment'} value={`${checkoutData.shipment} NTD`} />
                <CheckoutItem label={'Discount'} value={`-${checkoutData.total_discount} NTD`} />
                <CheckoutItem
                  label={'Total'}
                  value={`${checkoutData?.total} NTD`}
                  style={{ fontWeight: '900', textDecoration: 'underline' }}
                />
              </Col>

              <Col xs={12} style={LabelStyle}>
                Payment Method
                <form>
                  {checkoutData.payments.map((paymentData, index) => (
                    <Row style={ContentStyle}>
                      <Col xs={2} className='center' key={index}>
                        <input
                          type='radio'
                          value={index}
                          {...register('card_id')}
                          style={{ cursor: 'pointer' }}
                        />
                      </Col>
                      <Col xs={10}>{paymentData.name}</Col>
                    </Row>
                  ))}
                </form>
              </Col>

              <Col className='disappear_phone' />
              <Col xs={12} md={6} className='center'>
                <TButton text='Pay' action={onPay} />
              </Col>
              <Col className='disappear_phone' />
            </Row>
          </Offcanvas.Body>
        </Offcanvas>
      </div>

      {/* select coupons modal */}
      <Modal show={modalShow} onHide={() => setModalShow(false)} centered className='coupon_modal'>
        <Modal.Header style={{ border: 'none' }}>
          <Row className='center_vertical' style={{ width: '100%' }}>
            <Col xs={8} md={11} className='title'>
              Select Coupons
            </Col>
            <Col xs={4} md={1} className='right' style={{ padding: '0' }}>
              <FontAwesomeIcon
                icon={faCircleXmark as IconProp}
                size='2x'
                onClick={() => setModalShow(false)}
                style={{ cursor: 'pointer' }}
              />
            </Col>
          </Row>
        </Modal.Header>
        <Modal.Body>
          <Row style={{ width: '100%' }}>
            {usableCouponData.map((couponData, index) => (
              <Col xs={12} md={6} key={index} style={{ padding: '3%' }}>
                <div onClick={() => onApplyCoupon(couponData.id)} style={{ cursor: 'pointer' }}>
                  <CouponItemTemplate
                    data={{
                      name: couponData.name,
                      discount: couponData.discount,
                      type: couponData.type,
                      expire_date: formatDate(couponData.expire_date),
                    }}
                  />
                </div>
              </Col>
            ))}
          </Row>
        </Modal.Body>
      </Modal>
    </>
  );
};

export default Cart;
