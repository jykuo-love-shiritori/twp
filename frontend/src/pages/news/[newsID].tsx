import '@style/global.css';

import { useParams } from 'react-router-dom';
import { Col, Row } from 'react-bootstrap';

import NotFound from '@components/NotFound';

import newsData from '@pages/home/newsData.json';

interface Props {
  id: number;
  image_url: string;
  title: string;
  date: string;
  subTitle: string;
  content: string;
}

const EachNews = () => {
  const NewsBgStyle = {
    borderRadius: '50px 50px 0px 0px',
    border: '1px solid var(--button_border)',
    background: 'var(--bg)',
    boxShadow: '0px 4px 30px 2px var(--title)',
    padding: '10% 7% 10% 7%',
  };

  const NewsPicStyle = {
    width: '100%',
    border: '1px solid var(--button_border, #34977f)',
  };

  // TODO : data will be assign to the data got from backend, the newsData will be removed
  const params = useParams();
  const data: Props | undefined = newsData.find((news) => news.id.toString() === params.news_id);

  if (data) {
    return (
      <div style={{ padding: '10% 10% 0% 10%' }}>
        <div className='flex_wrapper' style={NewsBgStyle}>
          <Row>
            <Col xs={12} md={4} className='center_horizontal'>
              <img src={data.image_url} style={NewsPicStyle} />
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
