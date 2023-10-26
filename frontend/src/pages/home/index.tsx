import 'bootstrap/dist/css/bootstrap.min.css';
import '../../style/global.css';
import '../../components/style.css';

import { Col, Row } from 'react-bootstrap';

import News from '../../components/News';
import TButton from '../../components/TButton';
import GoodsItem from '../../components/GoodsItem';

import newsData from '../home/newsData.json';
import goodsData from '../discover/goodsData.json';

const Home = () => {
  return (
    <div>
      <div className='home'>
        <img src='../images/title.png' style={{ width: '100%' }}></img>
      </div>

      <div style={{ padding: '1% 15% 1% 15%' }}>
        <h2 className='title'>News</h2>
        <Row>
          {newsData.map((data) => {
            return (
              <Col xs={12} md={4}>
                <News id={data.id} imgUrl={data.imgUrl} title={data.title} />
              </Col>
            );
          })}
        </Row>

        <h2 className='title'>Popular Products</h2>

        <div style={{ padding: '0% 0% 3% 0%' }}>From most popular sellers</div>
        <Row>
          {goodsData.map((data) => {
            if (data.id < 5) {
              return (
                <Col xs={6} md={3}>
                  <GoodsItem id={data.id} name={data.name} imgUrl={data.imgUrl} />
                </Col>
              );
            }
            return null;
          })}
        </Row>

        <div style={{ padding: '3% 0% 3% 0%' }}>From local sellers</div>
        <Row>
          {goodsData.map((data) => {
            if (data.id < 9 && data.id > 4) {
              return (
                <Col xs={6} md={3}>
                  <GoodsItem id={data.id} name={data.name} imgUrl={data.imgUrl} />
                </Col>
              );
            }
            return null;
          })}
        </Row>

        <div style={{ padding: '3% 0% 3% 0%' }}>
          <TButton text='Explore more' url='/discover' />
        </div>
      </div>
    </div>
  );
};

export default Home;
