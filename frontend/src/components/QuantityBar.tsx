import '../components/style.css';
import '../style/global.css';

import { Col, Row } from 'react-bootstrap';
import { useState } from 'react';

const QuantityBar = () => {
  const [quality, setQuality] = useState(1);
  const handleAdd = () => setQuality(quality + 1);
  const handleMinus = () => setQuality(quality - 1 < 0 ? quality : quality - 1);

  return (
    <Row>
      <Col xs={3} onClick={handleMinus} className='pointer'>
        <div className='quantity_f pointer'>-</div>
      </Col>

      <Col xs={6} className='center'>
        <div>
          <input
            type='text'
            placeholder={`Quantity: ${quality}`}
            className='quantity_box'
            value={quality}
            onChange={(e) => setQuality(parseInt(e.target.value) || 0)}
          />
        </div>
      </Col>

      <Col xs={3} onClick={handleAdd} className='pointer'>
        <div className='quantity_f pointer'>+</div>
      </Col>
    </Row>
  );
};

export default QuantityBar;
