import CartItem from '@components/CartItem';
import TButton from '@components/TButton';

interface Props {
  item_id: number;
  quantity: number;
}

interface Input {
  title: string;
  cartContainer: Props[];
  total: number;
  updateTotal?: (subtotal: number, id: number) => void;
  removeItem?: (id: number) => void;
  isCart: boolean;
  isButtonNeeded: boolean;
}

const BoughtPage = ({
  title,
  cartContainer,
  total,
  updateTotal,
  removeItem,
  isCart,
  isButtonNeeded,
}: Input) => {
  return (
    <div style={{ padding: '10%' }}>
      <span className='titleWhite'>{title}</span>
      {cartContainer.map((data) => {
        return (
          <CartItem
            item_id={data.item_id}
            quantity={data.quantity}
            removeItem={removeItem}
            updateTotal={updateTotal}
            isCart={isCart}
          />
        );
      })}
      <hr className='white_bg' />
      <div className='center light'>Total : ${total}</div>
      {isButtonNeeded ? <TButton text='Submit' url='' /> : ''}
    </div>
  );
};

export default BoughtPage;
