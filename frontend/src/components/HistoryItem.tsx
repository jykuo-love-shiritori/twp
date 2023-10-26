import { Col, Row } from 'react-bootstrap';

import TButton from './TButton';
import UserItem from './UserItem';

import historyData from '../pages/cart/boughtData.json';
import goodsData from '../pages/discover/goodsData.json';

const HistoryItem = ({ id }: { id: number }) => {
  const record = historyData.find((item) => item.recordID === id);
  const firstItem = goodsData.find((item) => record?.items[0].item_id === item.id);

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

    return (
      <div>
        <Row className='history_container dark'>
          <Col sm={12} md={6}>
            <UserItem img_path={record.seller_img} name={record?.seller} />
          </Col>
          <Col sm={12} md={6} className='right'>
            Record ID : {record?.recordID}
          </Col>
          <Col sm={12} md={12}>
            <hr style={{ color: 'var(--button_dark)', opacity: '1', margin: '10px' }} />
          </Col>
          <Col xs={2} md={2} className='center'>
            <img src={firstItem?.imgUrl} style={{ width: '70px' }} />
          </Col>
          <Col xs={4} md={4} className='center_vertical'>
            {firstItem?.name}
          </Col>
          <Col xs={4} md={4} className='center_vertical'>
            Order Total : ${Total}
          </Col>
          <Col xs={12} md={2}>
            <TButton text='Detail' url={`/user/buyer/order/${record.recordID}`} />
          </Col>
        </Row>
      </div>
    );
  }
};

export default HistoryItem;
