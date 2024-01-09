import { Col, Row } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTrash } from '@fortawesome/free-solid-svg-icons';
import { useAuth } from '@lib/Auth';
import { useNavigate } from 'react-router-dom';
import { RouteOnNotOK } from '@lib/Status';
import { formatFloat } from '@lib/Functions';

interface IProduct {
  enabled: boolean;
  image_url: string;
  name: string;
  price: number;
  product_id: number;
  quantity: number;
  stock: number;
}

interface Props {
  data: IProduct;
  cart_id: number;
  refresh: () => void;
}

const CartProduct = ({ data, cart_id, refresh }: Props) => {
  const token = useAuth();
  const navigate = useNavigate();

  const removeItem = async () => {
    const resp = await fetch(`/api/buyer/cart/${cart_id}/product/${data.product_id}`, {
      method: 'DELETE',
      headers: {
        Authorization: `Bearer ${token}`,
        Accept: 'application/json',
      },
    });
    if (!resp.ok) {
      RouteOnNotOK(resp);
    } else {
      refresh();
    }
  };

  const updateQuantity = async (quantity: number, isReduce = false) => {
    if (quantity === 0) {
      removeItem();
    }
    // if original quant is above stock and the user is try to reduce it, set to stock
    if (isReduce && quantity > data.stock) {
      quantity = data.stock;
    }
    if (quantity > 0 && quantity <= data.stock) {
      const resp = await fetch(`/api/buyer/cart/${cart_id}/product/${data.product_id}`, {
        method: 'PATCH',
        headers: {
          Authorization: `Bearer ${token}`,
          Accept: 'application/json',
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ quantity: quantity }),
      });
      if (!resp.ok) {
        RouteOnNotOK(resp, navigate);
      } else {
        refresh();
      }
    }
  };

  return (
    <div className='cart_item' style={{ margin: '2% 0 2% 0' }}>
      <Row>
        <Col xs={4} md={1} className='center'>
          <img src={data.image_url} style={{ width: '100%', borderRadius: '10px' }} />
        </Col>
        <Col xs={8} md={11} className='dark center_vertical'>
          <div className='disappear_phone' style={{ width: '100%' }}>
            <Row className='center_vertical' style={{ width: '100%' }}>
              <Col
                md={4}
                className='center_vertical'
                style={{ wordBreak: 'break-all', fontSize: '20px' }}
              >
                {data.name}
              </Col>
              <Col md={5} className='center' style={{ padding: '2% 0' }}>
                <Row style={{ padding: '0', margin: '0' }}>
                  <Col
                    md={3}
                    onClick={() => updateQuantity(data.quantity - 1, true)}
                    className='pointer'
                  >
                    <div className='quantity_f pointer center '>-</div>
                  </Col>
                  <Col md={6} className='center'>
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
                  <Col md={3} onClick={() => updateQuantity(data.quantity + 1)} className='pointer'>
                    <div className='quantity_f pointer center'>+</div>
                  </Col>
                </Row>
              </Col>
              <Col md={2} className='right ' style={{ padding: '2% 0', fontSize: '20px' }}>
                {formatFloat(data.price * data.quantity)} NTD
              </Col>
              <Col md={1} className='center' style={{ padding: '2% 0' }}>
                <FontAwesomeIcon icon={faTrash} size='xl' className='trash' onClick={removeItem} />
              </Col>
            </Row>
          </div>

          <div className='disappear_tablet disappear_desktop'>
            <Row className='center_vertical' style={{ width: '100%' }}>
              <Col
                xs={12}
                className='center_vertical'
                style={{ wordBreak: 'break-all', padding: '2% 0 0 0' }}
              >
                {data.name}
              </Col>
              <Col xs={12} className='center' style={{ padding: '2% 0 0 0' }}>
                <Row style={{ padding: '0', margin: '0', width: '100%' }}>
                  <Col
                    xs={3}
                    onClick={() => updateQuantity(data.quantity - 1)}
                    className='pointer'
                    style={{ padding: '0 2%', margin: '0' }}
                  >
                    <div className='quantity_f pointer center '>-</div>
                  </Col>
                  <Col xs={6} className='center' style={{ padding: '0 2%', margin: '0' }}>
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
                  <Col
                    xs={3}
                    onClick={() => updateQuantity(data.quantity + 1)}
                    className='pointer'
                    style={{ padding: '0 2%', margin: '0' }}
                  >
                    <div className='quantity_f pointer center'>+</div>
                  </Col>
                </Row>
              </Col>
              <Col xs={8} style={{ padding: '2% 0 2% 5%' }}>
                {formatFloat(data.price * data.quantity)} NTD
              </Col>

              <Col xs={4} className='right' style={{ padding: '2% 5% 2% 0' }}>
                <FontAwesomeIcon icon={faTrash} className='trash' onClick={removeItem} />
              </Col>
            </Row>
          </div>
        </Col>
      </Row>
    </div>
  );
};

export default CartProduct;
