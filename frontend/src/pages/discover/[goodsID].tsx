import '../../style/global.css';

import { Col, Row } from 'react-bootstrap';

import NotFound from '../../components/NotFound';
import TButton from '../../components/TButton';
import QuantityBar from '../../components/QuantityBar';
import UserItem from '../../components/UserItem';

import goodsData from './goodsData.json';

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

const EachGoods = () => {
  // to get the goods' id from router
  const id = window.location.href.slice(-1);
  let data: Props = {
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
  };

  // to find the goods information by id
  const foundGoods = goodsData.find((goods) => goods.id.toString() === id);

  if (foundGoods) {
    Object.assign(data, foundGoods);
  }

  const isGoodsExist = !!foundGoods;

  if (isGoodsExist) {
    return (
      <div style={{ padding: '55px 12% 0 12%' }}>
        <Row>
          <Col xs={12} md={5} className='goods_bgW'>
            <div className='flex-wrapper' style={{ padding: '0 8% 10% 8%' }}>
              <img src={data.imgUrl} style={{ borderRadius: '0 0 30px 0' }} />

              {/* tags, price and quantity */}

              <br />

              <hr style={{ opacity: '1' }} />

              <QuantityBar />
              <TButton text='Add to cart' url='' />
            </div>
          </Col>
          <Col xs={12} md={7}>
            <div style={{ padding: '7% 0% 7% 10%' }}>
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
                  <UserItem img_path='../images/person.png' name='Tom Johnathan' />
                </Col>
                <Col xs={6}>
                  <TButton text='View Shop' url='' />
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
