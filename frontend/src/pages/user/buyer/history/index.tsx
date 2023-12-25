import { Col, Row } from 'react-bootstrap';

import HistoryItem from '@components/HistoryItem';

import historyData from '@pages/user/buyer/cart/boughtData.json';

const History = () => {
  return (
    <div>
      <div className='title'>Order history</div>
      <hr className='hr' />
      <br />

      <Row>
        {historyData.map((item, index) => {
          return (
            <Col xs={12} key={index}>
              <HistoryItem id={item.recordID} user='buyer' />
            </Col>
          );
        })}
      </Row>
    </div>
  );
};

export default History;
