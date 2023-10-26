import { Col, Row } from 'react-bootstrap';

import HistoryItem from './HistoryItem';

import historyData from '../pages/cart/boughtData.json';

const History = () => {
  return (
    <div>
      <div className='title'>Order history</div>
      <hr className='hr' />
      <br />

      <Row>
        {historyData.map((item) => {
          return (
            <Col xs={12}>
              <HistoryItem id={item.recordID} />
            </Col>
          );
        })}
      </Row>
    </div>
  );
};

export default History;
