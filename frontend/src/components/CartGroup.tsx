import { Col, Row } from 'react-bootstrap';

import CartItem from '@components/CartItem';
import UserItem from '@components/UserItem';
import sellerInfo from '@pages/user/seller/sellerInfo.json';
import TButton from './TButton';

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

const CartGroup = ({ data, onRefetch }: Props) => {
  return (
    <div className='cart_group'>
      <Row style={{ padding: '0 5%' }} className='center_vertical'>
        <Col xs={8} md={6}>
          <UserItem img_path={sellerInfo.imgUrl} name={sellerInfo.name} />
        </Col>

        <Col xs={4} md={3} className='center'>
          Subtotal: ${data.products.reduce((acc, cur) => acc + cur.price * cur.quantity, 0)}
        </Col>
        <Col xs={12} md={3} className='center'>
          <TButton text='Checkout' />
        </Col>
      </Row>
      {data.products.map((productData, index) => (
        <CartItem data={productData} cart_id={data.cartInfo.id} key={index} onRefetch={onRefetch} />
      ))}
    </div>
  );
  // }
};

export default CartGroup;
