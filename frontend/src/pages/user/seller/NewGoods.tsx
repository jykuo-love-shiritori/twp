import '@style/global.css';

import { Col, Row } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faFileUpload, faPen, faTrash } from '@fortawesome/free-solid-svg-icons';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import InputGroup from 'react-bootstrap/InputGroup';

import TButton from '@components/TButton';
import QuantityBar from '@components/QuantityBar';
import InfoItem from '@components/InfoItem';
import { useState } from 'react';

interface Props {
  id: number | null;
  price: number;
  name: string;
  introduction: string;
  sub_title: string;
  sub_content: string;
  calories: string;
  due_date: string;
  ingredients: string;
  imgUrl: string;
}

const EmptyGoods = () => {
  const tagStyle = {
    borderRadius: '30px',
    background: ' var(--button_light)',
    padding: '1% 1% 1% 3%',
    color: 'white',
    margin: '5px 0 5px 5px',
    width: '100%',
  };

  const [tag, setTag] = useState('');
  const [tagContainer, setTagContainer] = useState<string[]>([]);

  const addNewTag = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.keyCode === 229) return;

    if (event.key === 'Enter') {
      const input = event.currentTarget.value.trim();

      if (input !== '') {
        setTag('');
        setTagContainer((prevTags) => [...prevTags, input]);
      }
    }
  };

  const changeTag = (currentTag: string) => {
    const index = tagContainer.indexOf(currentTag);
  };

  return (
    <div style={{ padding: '55px 12% 0 12%' }}>
      <Row>
        <Col xs={12} md={5} className='goods_bgW'>
          <div className='flex-wrapper' style={{ padding: '0 8% 10% 8%' }}>
            <div
              style={{
                backgroundColor: 'black',
                padding: '30% 3% 5% 3%',
                borderRadius: '0 0 30px 0',
              }}
            >
              <div className='center'>
                <FontAwesomeIcon icon={faFileUpload} size='6x' />
              </div>
              <br />
              <Row>
                <Col xs={9}></Col>
                <Col xs={3}>
                  <form method='post' encType='multipart/form-data'>
                    <label htmlFor='file' className='custom-file-upload'>
                      <div className='button pointer center' style={{ padding: '10px' }}>
                        <FontAwesomeIcon icon={faPen} className='white_word' />
                      </div>
                    </label>

                    <input id='file' name='file' type='file' style={{ display: 'none' }} />
                  </form>
                </Col>
              </Row>
            </div>
            <br />
            <span className='dark'>add more tags</span>

            <input
              type='text'
              placeholder='Input tags'
              className='quantity_box'
              value={tag}
              onChange={(e) => setTag(e.target.value)}
              onKeyDown={addNewTag}
              style={{ marginBottom: '10px' }}
            />

            <Row xs='auto'>
              {tagContainer.map((currentTag, index) => (
                <Col style={tagStyle} key={index} className='center'>
                  <Col sm={true} style={tagStyle} key={index}>
                    <Row style={{ width: '100%' }} className='center'>
                      <Col xs={1} style={{ backgroundColor: '' }} className='center'>
                        <FontAwesomeIcon icon={faTrash} className='white_word' />
                      </Col>
                      <Col xs={1} style={{ backgroundColor: '' }} className='center'>
                        <FontAwesomeIcon icon={faPen} className='white_word' />
                      </Col>
                      <Col xs={10} style={{ backgroundColor: '' }}>
                        {/* {currentTag} */}
                        <input
                          type='text'
                          placeholder={currentTag}
                          // value={currentTag}
                          // onChange={(e) => setTag(e.target.value)}
                          style={{
                            border: 'var(--bg) 1px solid',
                            borderRadius: '30px',
                            padding: '0 10px 0 10px',
                            backgroundColor: 'transparent',
                            color: 'white',
                            width: '100%',
                          }}
                        />
                      </Col>
                    </Row>
                  </Col>
                </Col>
              ))}
            </Row>

            <div style={{ height: '50px' }} />
            <TButton text='Delete Product' url='' />
            <TButton text='Confirm Changes' url='' />
          </div>
        </Col>
        <Col xs={12} md={7}>
          <div style={{ padding: '7% 0% 7% 0%' }}>
            <InfoItem text='Product Name' isMore={false} />
            <InfoItem text='Product Price' isMore={false} />
            <InfoItem text='Product Quality' isMore={false} />
            <InfoItem text='Product Introduction' isMore={true} />
            <InfoItem text='Best Before Date' isMore={true} />
          </div>
        </Col>
      </Row>
    </div>
  );
};

export default EmptyGoods;
