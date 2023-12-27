import '@style/global.css';

import { useParams } from 'react-router-dom';
import { Col, Row } from 'react-bootstrap';
import { useQuery } from '@tanstack/react-query';
import { CSSProperties } from 'react';

import NotFound from '@components/NotFound';
import { CheckFetchStatus } from '@lib/Status';

const EachNews = () => {
  const { news_id } = useParams();

  const NewsBgStyle = {
    borderRadius: '50px 50px 0px 0px',
    border: '1px solid var(--button_border)',
    background: 'var(--bg)',
    boxShadow: '0px 4px 30px 2px var(--title)',
    padding: '10% 7% 10% 7%',
  };

  const NewsPicStyle: CSSProperties = {
    width: '100%',
    height: '100%',
    border: '1px solid var(--button_border, #34977f)',
    objectFit: 'cover',
  };

  const { status, data } = useQuery({
    queryKey: ['newsEach'],
    queryFn: async () => {
      const response = await fetch(`/api/news/${news_id}`, {
        headers: {
          Accept: 'application/json',
        },
      });
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      return response.json();
    },
  });

  if (status != 'success') {
    return <CheckFetchStatus status={status} />;
  }

  if (!data) {
    return <NotFound />;
  }

  return (
    <div style={{ padding: '10% 10% 0% 10%' }}>
      <div className='flex_wrapper' style={NewsBgStyle}>
        <Row>
          <Col xs={12} md={4} className='center_horizontal' style={{ overflow: ' hidden' }}>
            <img src={data.image_id} style={NewsPicStyle} />
          </Col>
          <Col xs={12} md={8}>
            <h4 className='inpage_title'>{data.title}</h4> <br />
            <hr className='hr' />
            <p>{data.content}</p>
          </Col>
        </Row>
      </div>
    </div>
  );
};

export default EachNews;
