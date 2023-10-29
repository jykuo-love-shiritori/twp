import TButton from '@components/TButton';
import CartGroup from '@components/CartGroup';

const Cart = () => {
  return (
    <div style={{ padding: '10%' }}>
      <span className='title'>Cart</span>

      <CartGroup />
      <CartGroup />
      <CartGroup />

      <div className='light'>Total:0</div>
      <TButton text='Submit' url='' />
    </div>
  );
};

export default Cart;
