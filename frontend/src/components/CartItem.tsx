import { Col, Row } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTrash } from '@fortawesome/free-solid-svg-icons';

import goodsData from '@pages/discover/goodsData.json';
import QuantityBar from '@components/QuantityBar';

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
  const data: Props = {
    item_id,
    quantity,
    updateTotal,
    name: '',
    imgUrl: '',
    subtotal: 0,
    isCart: false,
  };

  if (updateTotal) {
    updateTotal(data.subtotal, data.item_id);
  }

  const matchingGood = goodsData.find((goods) => goods.id === data.item_id);

  if (matchingGood) {
    data.name = matchingGood.name;
    data.imgUrl = matchingGood.imgUrl;
    data.subtotal = matchingGood.price * data.quantity;

    return (
      <div className='cart_item' style={{ margin: '2% 0 2% 0' }}>
        <Row>
          <Col xs={3} md={1} className=''>
            <img src={data.imgUrl} style={{ width: '70%' }} />
          </Col>
          <Col xs={9} md={11} className='dark'>
            <Row>
              <Col xs={9} md={5} className='center_vertical'>
                {data.name}
              </Col>

              <Col xs={12} md={4} className='center'>
                {isCart ? <QuantityBar /> : `x${data.quantity}`}
              </Col>

              <Col xs={6} md={2} className='center_vertical'>
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
  }
};

export default CartItem;
