import '@style/global.css';

import { Col, Row } from 'react-bootstrap';

import NotFound from '@components/NotFound';

import newsData from '@pages/home/newsData.json';

interface Props {
  id: number | null;
  imgUrl: string;
  title: string;
  date: string;
  subTitle: string;
  content: string;
}

const EachNews = () => {
  const id = window.location.href.slice(-1);
  console.log(id);
  const data: Props = { id: null, imgUrl: '', title: '', date: '', subTitle: '', content: '' };
  const foundNews = newsData.find((news) => news.id.toString() === id);

  if (foundNews) {
    Object.assign(data, foundNews);
  }
  const isNewsExist = !!foundNews;

  if (isNewsExist) {
    return (
      <div style={{ padding: '10% 10% 0% 10%' }}>
        <div className='news_bg flex-wrapper'>
          <Row>
            <Col xs={12} md={4}>
              <img src={data.imgUrl} className='news_pic' />
            </Col>
            <Col xs={12} md={8}>
              <h4 className='inpage_title'>{data.title}</h4> <br />
              <span className='right'>{data.date}</span>
              <hr className='hr' />
              <p>{data.subTitle}</p>
              <p>{data.content}</p>
            </Col>
          </Row>
        </div>
      </div>
    );
  } else {
    return <NotFound />;
  }
};

export default EachNews;
