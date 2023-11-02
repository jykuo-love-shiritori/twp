import { faFile, faMoneyBill, faTruck, faBox } from '@fortawesome/free-solid-svg-icons';
import { Col, Row } from 'react-bootstrap';

import CartItem from '@components/CartItem';
import NotFound from '@components/NotFound';
import RecordStatus from '@components/RecordStatus';
import UserItem from '@components/UserItem';

import goodsData from '@pages/discover/goodsData.json';
import historyData from '@pages/cart/boughtData.json';

const HistoryEach = () => {
  // get the id from router
  const id = window.location.href.slice(-1);

  // find the record if the id matches
  const record = historyData.find((item) => item.recordID.toString() === id);
  let Total = 0;

  if (record != undefined) {
    Total = record.items.reduce((accumulator, item) => {
      const foundItem = goodsData.find((goods) => goods.id === item.item_id);

      if (foundItem) {
        const subtotal = item.quantity * foundItem.price;
        accumulator += subtotal;
      }
      return accumulator;
    }, 0);

    const finalTotal: number = Total + record.shipment - record.coupon;

    return (
      <div style={{ padding: '7% 10% 10% 10%' }}>
        <div className='title'>Record ID : {record.recordID} </div>
        <Row>
          <Col xs={6} md={3}>
            <RecordStatus icon={faFile} text='Order placed' date={record.order_placed} />
          </Col>
          <Col xs={6} md={3}>
            <RecordStatus
              icon={faMoneyBill}
              text='Payment confirmed'
              date={record.payment_confirmed}
            />
          </Col>
          <Col xs={6} md={3}>
            <RecordStatus icon={faTruck} text='Shipped out' date={record.shipped_out} />
          </Col>
          <Col xs={6} md={3}>
            <RecordStatus icon={faBox} text='Order received' date={record.order_received} />
          </Col>
        </Row>

        <hr className='hr' />

        <UserItem img_path='../../../images/person.png' name='Tom Johnathan' />

        {record?.items.map((data) => {
          return <CartItem item_id={data.item_id} quantity={data.quantity} isCart={false} />;
        })}

        <Row className='light'>
          <Col xs={12} md={7} />
          <Col xs={6} md={3}>
            Original Total :
          </Col>
          <Col xs={6} md={2}>
            ${Total}
          </Col>

          <Col xs={12} md={7} />
          <Col xs={6} md={3}>
            Shipment :
          </Col>
          <Col xs={6} md={2}>
            ${record.shipment}
          </Col>

          <Col xs={12} md={7} />
          <Col xs={6} md={3}>
            Coupon :
          </Col>
          <Col xs={6} md={2}>
            -${record.coupon}
          </Col>
        </Row>
        <hr className='hr' />
        <Row className='light'>
          <Col xs={12} md={7} />
          <Col xs={6} md={3}>
            Order Total :
          </Col>
          <Col xs={6} md={2}>
            ${finalTotal}
          </Col>
        </Row>
      </div>
    );
  } else {
    return <NotFound />;
  }
};

export default HistoryEach;