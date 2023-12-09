import { Col, Row } from 'react-bootstrap';

import TButton from '@components/TButton';
import UserItem from '@components/UserItem';

import historyData from '@pages/user/buyer/cart/boughtData.json';
import goodsData from '@pages/discover/goodsData.json';

const HistoryItem = ({ id, user }: { id: number; user: string }) => {
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
            <hr
              style={{
                color: 'var(--button_dark)',
                opacity: '1',
                margin: '10px 0px 10px 0px',
                width: '100%',
              }}
            />
          </Col>
          <Col xs={4} md={2} lg={1} className='center'>
            <img src={firstItem?.imgUrl} style={{ width: '100%', borderRadius: '10px' }} />
          </Col>
          <Col xs={8} md={8} lg={9} className='center_vertical'>
            <Row style={{ width: '100%' }}>
              <Col xs={12} lg={6}>
                {firstItem?.name}
              </Col>
              <Col xs={12} lg={6}>
                Order Total : ${Total}
              </Col>
            </Row>
          </Col>
          <Col xs={12} md={2} className='right'>
            <TButton
              text='Detail'
              action={
                user === 'buyer'
                  ? `/user/buyer/order/${record.recordID}`
                  : `/user/seller/order/${record.recordID}`
              }
            />
          </Col>
        </Row>
      </div>
    );
  }
};

export default HistoryItem;
