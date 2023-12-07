import '@style/global.css';

import { Col, Row } from 'react-bootstrap';
import { useParams } from 'react-router-dom';

import NotFound from '@components/NotFound';
import TButton from '@components/TButton';
import QuantityBar from '@components/QuantityBar';
import UserItem from '@components/UserItem';

import goodsData from '@pages/discover/goodsData.json';

interface Tag {
  name: string;
}

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
  tags: Tag[];
  quantity: number;
}

const EachGoods = () => {
  const tagStyle = {
    borderRadius: '30px',
    background: ' var(--button_light)',
    padding: '2% 3% 2% 3%',
    color: 'white',
    margin: '10px 0 0px 5px',
  };

  const params = useParams();
  const id = params.goods_id;

  const data: Props = {
    id: null,
    price: 0,
    name: '',
    introduction: '',
    sub_title: '',
    sub_content: '',
    calories: '',
    due_date: '',
    ingredients: '',
    imgUrl: '',
    tags: [],
    quantity: 0,
  };

  // to find the goods information by id
  const foundGoods = goodsData.find((goods) => goods.id.toString() === id);

  if (foundGoods) {
    Object.assign(data, foundGoods);
    console.log(data.tags);
  }

  if (foundGoods) {
    return (
      <div style={{ padding: '55px 12% 0 12%' }}>
        <Row>
          <Col xs={12} md={5} className='goods_bgW'>
            <div className='flex-wrapper' style={{ padding: '0 8% 10% 8%' }}>
              <img src={data.imgUrl} style={{ borderRadius: '0 0 30px 0' }} />

              {/* tags, price and quantity */}
              <Row xs='auto'>
                {data.tags.map((currentTag, index) => (
                  <Col style={tagStyle} className='center' key={index}>
                    {currentTag.name}
                  </Col>
                ))}
              </Row>

              <h4 style={{ paddingTop: '30px', color: 'black', marginBottom: '5px' }}>
                $ {data.price} TWD
              </h4>

              {data.quantity != 0 ? (
                <div>
                  <span style={{ color: 'black' }}>{data.quantity} available</span>
                  <hr style={{ opacity: '1' }} />
                  <QuantityBar />
                  <TButton text='Add to cart' />
                </div>
              ) : (
                <h6 style={{ color: '#ED7E6D' }}>
                  <b>SOLD OUT</b>
                </h6>
              )}
            </div>
          </Col>
          <Col xs={12} md={7}>
            <div style={{ padding: '7% 5% 7% 5%' }}>
              <div className='inpage_title'>{data.name}</div>

              {/* goods' description */}
              <p>
                {data.introduction}
                <br />
                <br />
                {data.sub_title}
                <br />
                <br />
                {data.sub_content}
                <br />
                <br />
                Calories:{data.calories}
                <br />
                <br />
                Enjoy at its Freshest:{data.due_date}
              </p>

              <hr className='hr' />
              <Row>
                <Col xs={6} className='center'>
                  <UserItem img_path='/placeholder/person.png' name='Tom Johnathan' />
                </Col>
                <Col xs={6}>
                  <TButton text='View Shop' />
                </Col>
              </Row>
            </div>
          </Col>
        </Row>
      </div>
    );
  } else {
    return <NotFound />;
  }
};

export default EachGoods;
