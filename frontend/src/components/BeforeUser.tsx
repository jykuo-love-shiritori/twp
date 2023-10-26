import { Col, Row } from 'react-bootstrap';

import Footer from './Footer';

interface Props {
  imgUrl: string;
  title: string;
  content: string;
  subContent: string;
  way: string;
  path: string;
  url: string;
}

interface Data {
  data: Props | undefined;
}

const BeforeUser = ({ data }: Data) => {
  if (data) {
    return (
      <div>
        <div className='center' style={{ backgroundColor: 'var(--bg)' }}>
          <Row>
            <Col xs={12} md={6}>
              <img src={data?.imgUrl} style={{ height: '100%', width: '100%' }} />
            </Col>
            <Col xs={12} md={6} style={{ padding: '10% 10% 10% 10%' }}>
              <div className='title'>{data?.title}</div>
              <div style={{ padding: '20% 0 40% 0' }}>
                <p>{data?.content}</p>
              </div>

              <div className='center' style={{ margin: '5%' }}>
                <div className='forTest' />
              </div>

              <div className='center' style={{ fontSize: '12px' }}>
                <p>{data?.subContent}</p>
              </div>
              <br />

              <Row>
                <Col xs={4}>
                  <hr style={{ color: 'white' }} />
                </Col>
                <Col xs={4} className='center'>
                  <p>Or With</p>
                </Col>
                <Col xs={4}>
                  <hr style={{ color: 'white' }} />
                </Col>
              </Row>

              <div className='center'>
                <span style={{ color: 'white' }}>{data?.way} &nbsp; </span>
                <span>
                  <u>
                    <a href={data?.url}>{data?.path}</a>
                  </u>
                </span>
              </div>
            </Col>
          </Row>
        </div>
        <Footer />
      </div>
    );
  }
};

export default BeforeUser;
