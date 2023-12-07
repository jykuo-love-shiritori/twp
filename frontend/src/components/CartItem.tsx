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
  const CartItemStyle = {
    borderRadius: '30px',
    background: 'rgba(255, 255, 255, 0.6)',
    backdropFilter: 'blur(10px)',
    padding: '2% 30px 2% 30px',
    margin: '2% 0 2% 0',
  };

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
      <div style={CartItemStyle}>
        <Row>
          <Col xs={4} md={1} className='center'>
            <img src={data.imgUrl} style={{ width: '100%', borderRadius: '10px' }} />
          </Col>
          <Col xs={8} md={11} className='dark center_vertical'>
            <Row style={{ width: '100%' }}>
              <Col xs={12} md={5} className='center_vertical' style={{ wordBreak: 'break-all' }}>
                {data.name}
              </Col>

              <Col xs={12} md={isCart ? 4 : 2} className='right'>
                {isCart ? <QuantityBar /> : `x${data.quantity}`}
              </Col>

              <Col xs={isCart ? 6 : 12} md={isCart ? 2 : 4} className='right'>
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
