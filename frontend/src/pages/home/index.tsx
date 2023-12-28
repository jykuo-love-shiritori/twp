import 'bootstrap/dist/css/bootstrap.min.css';
import '@style/global.css';
import '@components/style.css';

import { Col, Row } from 'react-bootstrap';
import { useQuery } from '@tanstack/react-query';
import { useNavigate } from 'react-router-dom';

import News from '@components/News';
import TButton from '@components/TButton';
import GoodsItem from '@components/GoodsItem';

import { useAuth } from '@lib/Auth';

import TitleImgUrl from '@assets/images/title.png';
import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import { NewsProps } from '@components/News';
import { GoodsItemProps } from '@components/GoodsItem';

const Home = () => {
  const token = useAuth();
  const navigate = useNavigate();

  const { status: newsStatus, data: newsData } = useQuery({
    queryKey: ['news'],
    queryFn: async () => {
      const response = await fetch(`/api/news`, {
        headers: {
          Accept: 'application/json',
        },
      });
      if (!response.ok) {
        RouteOnNotOK(response, navigate);
      }
      return (await response.json()) as NewsProps[];
    },
  });

  const { status: recommendStatus, data: recommendData } = useQuery({
    queryKey: ['recommend'],
    queryFn: async () => {
      const response = await fetch(`/api/popular`, {
        headers: {
          Accept: 'application/json',
        },
      });
      if (!response.ok) {
        RouteOnNotOK(response, navigate);
      }
      return await response.json();
    },
  });

  if (newsStatus != 'success') {
    return <CheckFetchStatus status={newsStatus} />;
  }

  if (recommendStatus != 'success') {
    return <CheckFetchStatus status={recommendStatus} />;
  }

  console.log(token);

  return (
    <div>
      <div className='home'>
        <img src={TitleImgUrl} style={{ width: '100%' }}></img>
      </div>

      <div style={{ padding: '1% 15% 1% 15%' }}>
        <h2 className='title'>News</h2>
        <Row>
          {newsData.map((data: NewsProps, index: number) => {
            return (
              <Col xs={12} md={4} key={index}>
                <News id={data.id} image_id={data.image_id} title={data.title} />
              </Col>
            );
          })}
        </Row>

        <h2 className='title'>Popular Products</h2>

        <div style={{ padding: '0% 0% 3% 0%' }}>From most popular sellers</div>
        <Row>
          {recommendData.popular_products.map((data: GoodsItemProps, index: number) => (
            <Col xs={6} md={3} key={index}>
              <GoodsItem id={data.id} name={data.name} image_url={data.image_url} />
            </Col>
          ))}
        </Row>

        <div style={{ padding: '3% 0% 3% 0%' }}>From local sellers</div>
        <Row>
          {recommendData.local_products.map((data: GoodsItemProps, index: number) => (
            <Col xs={6} md={3} key={index}>
              <GoodsItem id={data.id} name={data.name} image_url={data.image_url} />
            </Col>
          ))}
        </Row>

        <div style={{ padding: '3% 0% 3% 0%' }}>
          <TButton text='Explore more' action='/discover' />
        </div>
      </div>
    </div>
  );
};

export default Home;
