import Form from 'react-bootstrap/Form';
import { useState } from 'react';
import { Col, Row } from 'react-bootstrap';

import CartItem from '@components/CartItem';
import UserItem from '@components/UserItem';

interface Props {
  item_id: number;
  quantity: number;
  subtotal: number;
}

const CartGroup = () => {
  const initialData: Props[] = [];
  for (let i = 0; i < 3; i++) {
    const newData: Props = {
      item_id: 2 + i,
      quantity: 3 + 3 * i,
      subtotal: 0,
    };
    initialData.push(newData);
  }
  const [cartContainer, setCartContainer] = useState<Props[]>(initialData);

  const removeItem = (id: number) => {
    setCartContainer((prevCartContainer) => {
      // const item = prevCartContainer.filter((item) => item.item_id === id);
      const updateCartContainer = prevCartContainer.filter((item) => item.item_id !== id);
      return updateCartContainer;
    });
  };

  if (cartContainer.length != 0) {
    return (
      <div className='cart_group'>
        <Row style={{ padding: '0 0 0 5%' }}>
          <Col xs={1} className='right'>
            <Form.Check type={'checkbox'} />
          </Col>
          <Col xs={11}>
            <UserItem img_path='../images/person.png' name='Tom Johnathan' />
          </Col>
        </Row>

        {cartContainer.map((data) => {
          return (
            <CartItem
              item_id={data.item_id}
              quantity={data.quantity}
              removeItem={removeItem}
              // updateTotal={updateTotal}
              isCart={true}
            />
          );
        })}
      </div>
    );
  }
};

export default CartGroup;
