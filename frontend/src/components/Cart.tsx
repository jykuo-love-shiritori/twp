import { Row, Col, Offcanvas, Modal } from 'react-bootstrap';
import { useState } from 'react';
import { useQuery } from '@tanstack/react-query';
import { useNavigate } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCircleXmark } from '@fortawesome/free-regular-svg-icons';
import { IconProp } from '@fortawesome/fontawesome-svg-core';

import CartProduct from '@components/CartProduct';
import sellerInfo from '@pages/user/seller/sellerInfo.json';
import UserItem from '@components/UserItem';
import TButton from './TButton';
import CouponItemTemplate from './CouponItemTemplate';

interface Props {
  data: CartProps;
  onRefetch: () => void;
}

interface CartProps {
  cartInfo: { id: number; image_id: string; seller_name: string; shop_name: string };
  coupons: CouponProps[];
  products: ProductProps[];
}

interface CouponProps {
  description: string;
  discount: number;
  id: number;
  name: string;
  type: string; // 'percentage' | 'fixed' | 'shipping'
  scope: string; // 'global' | 'shop'
}

interface ProductProps {
  enabled: boolean;
  image_id: string;
  name: string;
  price: number;
  product_id: number;
  quantity: number;
  stock: number;
}

interface CheckoutProps {
  coupons: [
    {
      description: string;
      discount: number;
      discount_value: number;
      id: number;
      name: string;
      scope: string; // 'global' | 'shop'
      type: string; // 'percentage' | 'fixed' | 'shipping'
    },
  ];
  shipment: number;
  subtotal: number;
  total: number;
  total_discount: number;
}

interface UsableCouponProps {
  description: string;
  discount: number;
  expire_date: string;
  id: number;
  name: string;
  scope: string; // 'global' | 'shop'
  type: string; // 'percentage' | 'fixed' | 'shipping'
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

const RouteOnError = (code: number) => {
  const navigate = useNavigate();
  switch (code) {
    case 401:
      navigate('/unauthorized');
      break;
    case 403:
      navigate('/forbidden');
      break;
    case 404:
      navigate('/notFound');
      break;
  }
};

const Cart = ({ data, onRefetch }: Props) => {
  // get the checkout detail
  // TODO: Buyer get checkout
  // GET /buyer/cart/:cart_id/checkout
  const [canvaShow, setCanvaShow] = useState(false);
  const {
    data: checkoutData,
    status: checkoutStatus,
    refetch: refetchCheckout,
  } = useQuery({
    queryKey: ['checkoutData'],
    queryFn: async () => {
      const response = await fetch('/resources/Checkout.json');
      if (!response.ok) {
        RouteOnError(response.status);
      }
      return response.json();
    },
    select: (data) => data as CheckoutProps,
    refetchOnWindowFocus: false,
    enabled: false,
  });
  const onRefetchCheckout = () => {
    console.log('refetch checkout');
    refetchCheckout();
    if (checkoutStatus === 'pending') {
      return <div>Loading...</div>;
    }
    if (checkoutStatus === 'error') {
      return <div>Error fetching data</div>;
    }
  };

  // get the usable coupons
  // TODO: Buyer get usable coupon of cart/shop
  // GET /buyer/cart/:cart_id/coupon
  const [modalShow, setModalShow] = useState(false);
  const {
    data: usableCouponData,
    status: usableCouponStatus,
    refetch: refetchCoupon,
  } = useQuery({
    queryKey: ['usableCouponData'],
    queryFn: async () => {
      const response = await fetch('/resources/UsableCoupons.json');
      if (!response.ok) {
        RouteOnError(response.status);
      }
      return response.json();
    },
    select: (data) => data as UsableCouponProps[],
    refetchOnWindowFocus: false,
    enabled: false,
  });
  const onRefetchUsableCoupon = () => {
    console.log('refetch usable coupon');
    refetchCoupon();
    if (usableCouponStatus === 'pending') {
      return <div>Loading...</div>;
    }
    if (usableCouponStatus === 'error') {
      return <div>Error fetching data</div>;
    }
  };

  const onViewCheckout = () => {
    // TODO: Buyer get checkout
    // GET /buyer/cart/:cart_id/checkout
    onRefetchCheckout();
    setCanvaShow(true);
  };
  const onPay = () => {
    // TODO: Buyer checkout
    // POST /buyer/cart/:cart_id/checkout
    console.log('pay');
    onRefetch();
    setCanvaShow(false);
  };
  const onChooseCoupon = () => {
    // TODO: Buyer get usable coupon of cart
    // GET /buyer/cart/:cart_id/coupon
    onRefetchUsableCoupon();
    setModalShow(true);
  };
  const onApplyCoupon = (coupon_id: number) => {
    // TODO: Buyer apply coupon to cart
    // POST /buyer/cart/:cart_id/coupon/:coupon_id
    console.log(`apply coupon ${coupon_id}`);
    onRefetchCheckout();
    setModalShow(false);
  };
  const onRemoveCoupon = (coupon_id: number) => {
    // TODO: Buyer delete coupon in cart
    // DELETE /buyer/cart/:cart_id/coupon/:coupon_id
    console.log(`remove coupon ${coupon_id}`);
    onRefetchCheckout();
  };

  return (
    <>
      <div className='cart_group'>
        <div className='disappear_phone' style={{ fontSize: '20px' }}>
          <Row className='center_vertical' style={{ width: '100%', padding: '0 2%' }}>
            <Col md={6}>
              <UserItem img_path={sellerInfo.imgUrl} name={sellerInfo.name} />
            </Col>
            <Col md={3} className='center'>
              Subtotal: ${data.products.reduce((acc, cur) => acc + cur.price * cur.quantity, 0)}
            </Col>
            <Col md={3} className='center'>
              <TButton text='Checkout' action={onViewCheckout} />
            </Col>
          </Row>
        </div>

        <div className='disappear_tablet disappear_desktop'>
          <Row className='center_vertical' style={{ width: '100%', padding: '0 3%', margin: '0' }}>
            <Col xs={6} style={{ padding: '0 0 0 5%' }}>
              <UserItem img_path={sellerInfo.imgUrl} name={sellerInfo.name} />
            </Col>
            <Col xs={6} className='right'>
              Subtotal: ${data.products.reduce((acc, cur) => acc + cur.price * cur.quantity, 0)}
            </Col>
            <Col xs={12} className='center'>
              <TButton text='Checkout' action={onViewCheckout} />
            </Col>
          </Row>
        </div>

        {data.products.map((productData, index) => (
          <CartProduct
            data={productData}
            cart_id={data.cartInfo.id}
            key={index}
            onRefetch={onRefetch}
          />
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
                Delivery
                <Row style={ContentStyle}>
                  <Col xs={6}>Shipment</Col>
                  <Col xs={6} className='right'>
                    {checkoutData?.shipment}
                  </Col>
                </Row>
              </Col>

              <Col xs={12} style={LabelStyle}>
                <Row>
                  <Col xs='auto' style={{ paddingRight: '0' }}>
                    Discount
                  </Col>
                  <Col
                    xs='auto'
                    onClick={onChooseCoupon}
                    style={{ fontWeight: '900', cursor: 'pointer' }}
                  >
                    +
                  </Col>
                </Row>

                {checkoutData?.coupons.map((couponData, index) => (
                  <Row style={ContentStyle} key={index}>
                    <Col xs={6}>
                      <Row>
                        <Col
                          xs='auto'
                          onClick={() => onRemoveCoupon(couponData.id)}
                          style={{ fontWeight: '900', cursor: 'pointer' }}
                        >
                          -
                        </Col>
                        {couponData.name}
                        <Col></Col>
                      </Row>
                    </Col>
                    <Col xs={6} className='right'>
                      {couponData.discount_value}
                    </Col>
                  </Row>
                ))}
              </Col>

              <Col xs={12} style={LabelStyle}>
                Summary
                <Row style={ContentStyle}>
                  <Col xs={6}>subtotal</Col>
                  <Col xs={6} className='right'>
                    {checkoutData?.subtotal}
                  </Col>
                  <Col xs={6}>shipment</Col>
                  <Col xs={6} className='right'>
                    {checkoutData?.shipment}
                  </Col>
                  <Col xs={6}>discount</Col>
                  <Col xs={6} className='right'>
                    -{checkoutData?.total_discount}
                  </Col>
                  <Col
                    xs={6}
                    style={{ fontWeight: '900', fontSize: '20px', textDecoration: 'underline' }}
                  >
                    total
                  </Col>
                  <Col
                    xs={6}
                    className='right'
                    style={{ fontWeight: '900', fontSize: '20px', textDecoration: 'underline' }}
                  >
                    {checkoutData?.total}
                  </Col>
                </Row>
              </Col>

              <Col xs={12} style={LabelStyle}>
                Payment Method
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
            {usableCouponData?.map((couponData, index) => (
              <Col xs={12} md={6} key={index} style={{ padding: '3%' }}>
                <div onClick={() => onApplyCoupon(couponData.id)} style={{ cursor: 'pointer' }}>
                  <CouponItemTemplate data={couponData} />
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
