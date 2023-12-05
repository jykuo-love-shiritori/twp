import { useState } from 'react';

import TButton from '@components/TButton';
import InfoItem from '@components/InfoItem';
import { Col, Form, Row } from 'react-bootstrap';

const SellerInfo = () => {
  //TODO: read the initial value
  const [visibility, setVisibility] = useState<boolean>(true);
  const [shopName, setShopName] = useState<string>('');
  const [description, setDescription] = useState<string>('');

  const handleVisibility = (e: React.ChangeEvent<HTMLInputElement>) => {
    setVisibility(e.target.checked);
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
            value={visibility ? 1 : 0}
            onChange={handleVisibility}
          />
          {visibility ? 'Your shop is visible to everyone.' : 'Your shop is hidden from everyone.'}
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
