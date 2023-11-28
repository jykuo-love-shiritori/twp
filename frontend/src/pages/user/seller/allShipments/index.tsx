import { Col, Row } from 'react-bootstrap';

import HistoryItem from '@components/HistoryItem';

import historyData from '@pages/cart/boughtData.json';

const SellerShipment = () => {
  return (
    <div>
      <div className='title'>All Shipments</div>
      <hr className='hr' />
      <br />

      <Row>
        {historyData.map((item, index) => {
          return (
            <Col xs={12} key={index}>
              <HistoryItem id={item.recordID} user='seller' />
            </Col>
          );
        })}
      </Row>
    </div>
  );
};

export default SellerShipment;
