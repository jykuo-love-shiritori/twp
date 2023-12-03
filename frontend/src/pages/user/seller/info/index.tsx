import { useState } from 'react';

import TButton from '@components/TButton';
import InfoItem from '@components/InfoItem';
import { Col, Form, Row } from 'react-bootstrap';

const SellerInfo = () => {
  //TODO: read the initial value
  const [visiblity, setVisiblity] = useState<boolean>(true);
  const [shopName, setshopName] = useState<string>('');
  const [description, setDescription] = useState<string>('');

  const handleVisiblity = (e: React.ChangeEvent<HTMLInputElement>) => {
    setVisiblity(e.target.checked);
  };

  return (
    <div>
      <div className='title'>Shop info</div>
      <hr className='hr' />
      <Row style={{ margin: '2% 0% 2% 0% ' }}>
        <Col xs={12} md={4} className='center_vertical'>
          Visiblity
        </Col>
        <Col className='left'>
          <Form.Check
            type='checkbox'
            id='visiblity_checkbox'
            label=''
            value={visiblity ? 1 : 0}
            onChange={handleVisiblity}
          />
          {visiblity ? 'Your shop is visible to everyone.' : 'Your shop is hidden from everyone.'}
        </Col>
      </Row>
      <InfoItem text='Shop Name' isMore={false} value={shopName} setValue={setshopName} />
      <InfoItem
        text='Shop Description'
        isMore={true}
        value={description}
        setValue={setDescription}
      />
      <TButton text='Save' url='' />
    </div>
  );
};

export default SellerInfo;
