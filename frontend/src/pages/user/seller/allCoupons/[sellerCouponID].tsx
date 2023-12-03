import { useState } from 'react';
import { useParams } from 'react-router-dom';
import { Row, Col } from 'react-bootstrap';
import DatePicker from 'react-datepicker';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faPen, faTrash } from '@fortawesome/free-solid-svg-icons';
import 'react-datepicker/dist/react-datepicker.css';

import TButton from '@components/TButton';
import InfoItem from '@components/InfoItem';
import CouponItem from '@components/CouponItem';
import couponData from '@pages/coupon/couponData.json';

interface Props {
  id: number | null;
  name: string;
  policy: string;
  date: string;
  introduction: string;
  tags: { name: string }[];
}

const EachSellerCoupon = () => {
  const params = useParams();
  const id = params.coupon_id;

  const data: Props = {
    id: null,
    name: '',
    policy: '',
    date: '',
    introduction: '',
    tags: [],
  };

  const foundCoupon = couponData.find((coupon) => coupon.id.toString() === id);

  if (foundCoupon) {
    Object.assign(data, foundCoupon);
  }

  const tagStyle = {
    borderRadius: '30px',
    background: ' var(--button_light)',
    padding: '1% 1% 1% 3%',
    color: 'white',
    margin: '5px 0 5px 5px',
    width: '100%',
  };

  const [tag, setTag] = useState('');
  const [tagContainer, setTagContainer] = useState<string[]>(data.tags.map((tag) => tag.name));
  const [modification, setModification] = useState<boolean[]>(Array(data.tags.length).fill(false));
  const [name, setName] = useState<string>(data.name);
  const [policy, setPolicy] = useState<string>(data.policy);
  const [introduction, setIntroduction] = useState<string>(data.introduction);
  const [date, setDate] = useState<string>(data.date);

  const addNewTag = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.keyCode === 229) return;

    if (event.key === 'Enter') {
      const input = event.currentTarget.value.trim();
      console.log(event.currentTarget.value);

      if (input !== '') {
        setTagContainer((prevTags) => [...prevTags, input]);
        setModification((prevModification) => [...prevModification, false]);
        setTag('');
      }
      console.log(tagContainer, modification);
    }
  };

  const deleteTag = (index: number) => {
    setTagContainer((prevTags) => prevTags.filter((_, i) => i !== index));
    setModification((prevModifications) => prevModifications.filter((_, i) => i !== index));
  };

  const changeModification = (index: number) => {
    setModification((prevModifications) =>
      prevModifications.map((mod, i) => (i === index ? !mod : mod)),
    );
  };

  const changeTag = (index: number, value: string) => {
    setTagContainer((prevTags) => prevTags.map((tag, i) => (i === index ? value : tag)));
  };

  const changeDate = (date: Date) => {
    setDate(date.toISOString().split('T')[0].replace(/-/g, '/'));
  };

  return (
    <div style={{ padding: '55px 12% 0 12%' }}>
      <Row>
        <Col xs={12} md={5} className='goods_bgW'>
          <div className='flex-wrapper' style={{ padding: '0 8% 10% 8%' }}>
            <div style={{ padding: '15% 10%' }}>
              <CouponItem data={data} />
            </div>
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
                  <Row style={{ width: '100%' }} className='center'>
                    <Col xs={1} className='center'>
                      <FontAwesomeIcon
                        icon={faTrash}
                        className='white_word pointer'
                        onClick={() => deleteTag(index)}
                      />
                    </Col>
                    <Col xs={1} className='center'>
                      <FontAwesomeIcon
                        icon={faPen}
                        className='white_word pointer'
                        onClick={() => changeModification(index)}
                      />
                    </Col>
                    <Col xs={8} lg={10}>
                      {modification[index] ? (
                        <input
                          type='text'
                          placeholder={currentTag}
                          value={currentTag}
                          onChange={(e) => changeTag(index, e.target.value)}
                          style={{
                            border: 'var(--bg) 1px solid',
                            borderRadius: '30px',
                            padding: '0 10px 0 10px',
                            backgroundColor: 'transparent',
                            color: 'white',
                            width: '100%',
                          }}
                        />
                      ) : (
                        <span style={{ wordBreak: 'break-all' }}>{currentTag}</span>
                      )}
                    </Col>
                  </Row>
                </Col>
              ))}
            </Row>

            <div style={{ height: '50px' }} />
            <TButton text='Delete Coupon' url='' />
            <TButton text='Confirm Changes' url='' />
          </div>
        </Col>
        <Col xs={12} md={7}>
          <div style={{ padding: '7% 0% 7% 0%' }}>
            <InfoItem text='Coupon Name' isMore={false} value={name} setValue={setName} />
            <InfoItem text='Coupon Policy' isMore={false} value={policy} setValue={setPolicy} />
            <Row style={{ margin: '2% 0% 2% 0% ' }}>
              <Col xs={12} md={4} className='center_vertical'>
                Date
              </Col>
              <Col xs={12} md={8} className='coupon_date_picker'>
                <DatePicker value={date} onChange={changeDate} />
              </Col>
            </Row>
            <InfoItem
              text='Coupon Introduction'
              isMore={true}
              value={introduction}
              setValue={setIntroduction}
            />
          </div>
        </Col>
      </Row>
    </div>
  );
};

export default EachSellerCoupon;