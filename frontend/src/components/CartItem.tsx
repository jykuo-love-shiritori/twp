import { Col, Row } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTrash } from '@fortawesome/free-solid-svg-icons';

import goodsData from '../pages/discover/goodsData.json';
import QuantityBar from './QuantityBar';

interface Input {
  item_id: number;
  quantity: number;
  updateTotal?: (subtotal: number, id: number) => void;
  removeItem?: (id: number) => void;
  isCart: boolean;
}

interface Props extends Input {
  name: string;
  imgUrl: string;
  subtotal: number;
}

const CartItem = ({ item_id, quantity, updateTotal, removeItem, isCart }: Input) => {
  let data: Props = {
    item_id,
    quantity,
    updateTotal,
    name: '',
    imgUrl: '',
    subtotal: 0,
    isCart: false,
  };
  const matchingGood = goodsData.find((goods) => goods.id === data.item_id);

  if (matchingGood) {
    data.name = matchingGood.name;
    data.imgUrl = matchingGood.imgUrl;
    data.subtotal = matchingGood.price * data.quantity;
  }

  if (updateTotal) {
    updateTotal(data.subtotal, data.item_id);
  }

  return (
    <div className='cart_item' style={{ margin: '2% 0 2% 0' }}>
      <Row>
        <Col xs={3} md={1} className='center'>
          <img src={data.imgUrl} style={{ width: '100%' }} />
        </Col>
        <Col xs={9} md={11} className='dark'>
          <Row>
            <Col xs={9} md={4} className='center_vertical'>
              {data.name}
            </Col>

            <Col xs={12} md={4} className='right'>
              {isCart ? <QuantityBar /> : `x${data.quantity}`}
            </Col>

            <Col xs={6} md={2} className='center'>
              ${data.subtotal}
            </Col>

            <Col xs={6} md={1} className='center'>
              {isCart && removeItem && (
                <FontAwesomeIcon
                  icon={faTrash}
                  className='trash'
                  onClick={() => removeItem(data.item_id)}
                />
              )}
            </Col>
          </Row>
        </Col>
      </Row>
    </div>
  );
};

export default CartItem;
