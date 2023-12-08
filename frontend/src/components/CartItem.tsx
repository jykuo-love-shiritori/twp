import { Col, Row } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTrash } from '@fortawesome/free-solid-svg-icons';

interface Props {
  data: ProductProps;
  cart_id: number;
  onRefetch: () => void;
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

const CartItem = ({ data, cart_id, onRefetch }: Props) => {
  const removeItem = () => {
    //TODO: DELETE /buyer/cart/:cart_id/product/:product_id
    console.log(`${data.name} delete in cart ${cart_id}`);
    onRefetch();
  };

  const updateQuantity = (quantity: number) => {
    //TODO: PATCH /buyer/cart/:cart_id/product/:product_id
    // body: { quantity: number }
    if (quantity === 0) {
      removeItem();
    } else if (quantity > 0 && quantity <= data.stock) {
      console.log(`${data.name} ${quantity} in cart ${cart_id}`);
      onRefetch();
    }
  };

  return (
    <div className='cart_item' style={{ margin: '2% 0 2% 0' }}>
      <Row>
        <Col xs={4} md={1} className='center'>
          <img src={data.image_id} style={{ width: '100%', borderRadius: '10px' }} />
        </Col>
        <Col xs={8} md={11} className='dark center_vertical'>
          <Row className='center_vertical' style={{ width: '100%' }}>
            <Col
              xs={12}
              md={5}
              className='center_vertical'
              style={{ wordBreak: 'break-all', padding: '2% 0' }}
            >
              {data.name}
            </Col>

            <Col xs={6} md={4} className='right' style={{ padding: '2% 0' }}>
              <Row>
                <Col xs={3} onClick={() => updateQuantity(data.quantity - 1)} className='pointer'>
                  <div className='quantity_f pointer center'>-</div>
                </Col>

                <Col xs={6} className='center'>
                  <div>
                    <input
                      type='text'
                      className='quantity_box'
                      value={data.quantity}
                      onChange={(e) => updateQuantity(parseInt(e.target.value))}
                      style={{ textAlign: 'center' }}
                    />
                  </div>
                </Col>

                <Col xs={3} onClick={() => updateQuantity(data.quantity + 1)} className='pointer'>
                  <div className='quantity_f pointer center'>+</div>
                </Col>
              </Row>
            </Col>

            <Col xs={4} md={2} className='right'>
              {data.price * data.quantity} NTD
            </Col>

            <Col xs={2} md={1} className='center'>
              <FontAwesomeIcon icon={faTrash} className='trash' onClick={removeItem} />
            </Col>
          </Row>
        </Col>
      </Row>
    </div>
  );
};

export default CartItem;
