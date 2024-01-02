import '@style/global.css';

import { Col, Row } from 'react-bootstrap';
import { useParams } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';
import { useNavigate } from 'react-router-dom';

import TButton from '@components/TButton';
import QuantityBar from '@components/QuantityBar';
import UserItem from '@components/UserItem';
import { GetResponseProps } from '@pages/user/seller/allProducts/[sellerGoodsID]';
import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import { GoodsImgStyle } from '@pages/user/seller/allProducts/NewGoods';
import { useAuth } from '@lib/Auth';

const tagStyle = {
  borderRadius: '30px',
  background: ' var(--button_light)',
  padding: '2% 3% 2% 3%',
  color: 'white',
  margin: '10px 0 0px 5px',
};

const LeftBgStyle = {
  backgroundColor: 'rgba(255, 255, 255, 0.7)',
  boxShadow: '6px 4px 10px 2px rgba(0, 0, 0, 0.25)',
};

const EachGoods = () => {
  const token = useAuth();
  const navigate = useNavigate();

  const params = useParams<{ goods_id?: string }>();
  let goods_id: number | undefined;

  if (params.goods_id) {
    goods_id = parseInt(params.goods_id);
  }

  const { status, data } = useQuery({
    queryKey: ['goodsPage', goods_id],
    queryFn: async () => {
      if (goods_id === undefined) {
        throw new Error('Invalid goods_id');
      }
      const response = await fetch(`/api/seller/product/${goods_id}`, {
        headers: {
          Accept: 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });
      if (!response.ok) {
        RouteOnNotOK(response, navigate);
      }
      return (await response.json()) as GetResponseProps;
    },
  });

  if (status != 'success') {
    return <CheckFetchStatus status={status} />;
  } else {
    const originalDate = new Date(data.product_info.expire_date);
    data.product_info.expire_date = `${originalDate.getFullYear()}-${String(
      originalDate.getMonth() + 1,
    ).padStart(2, '0')}-${String(originalDate.getDate()).padStart(2, '0')}`;
  }

  return (
    <div style={{ padding: '55px 12% 0 12%' }}>
      <Row>
        <Col xs={12} md={5} style={LeftBgStyle}>
          <div className='flex_wrapper' style={{ padding: '0 8% 10% 8%' }}>
            <div style={{ overflow: ' hidden' }}>
              <img src={data.product_info.image_url} alt='File preview' style={GoodsImgStyle} />
            </div>

            <Row xs='auto'>
              {data.tags.map((currentTag, index) => (
                <Col style={tagStyle} className='center' key={index}>
                  {currentTag.name}
                </Col>
              ))}
            </Row>

            <h4 style={{ paddingTop: '30px', color: 'black', marginBottom: '5px' }}>
              $ {data.product_info.price} TWD
            </h4>

            {data.product_info.stock != 0 ? (
              <div>
                <span style={{ color: 'black' }}>{data.product_info.stock} available</span>
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
            <div className='inpage_title'>{data.product_info.name}</div>

            <p>
              {data.product_info.description}
              <br />
              <br />
              Enjoy at its Freshest : {data.product_info.expire_date}
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
};

export default EachGoods;
