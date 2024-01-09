import '@style/global.css';

import { Col, Row } from 'react-bootstrap';
import { useParams } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';
import { useNavigate } from 'react-router-dom';

import TButton from '@components/TButton';
import UserItem from '@components/UserItem';
import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import { GoodsImgStyle } from '@pages/user/seller/allProducts/NewGoods';
import { useAuth } from '@lib/Auth';
import { formatDate } from '@lib/Functions';
import { useState } from 'react';
import NotFound from '@components/NotFound';

interface IProduct {
  description: string;
  expire_date: string;
  id: number;
  product_image_url: string;
  name: string;
  price: number;
  sales: number;
  stock: number;
  seller_name: string;
  shop_name: string;
  shop_image_url: string;
  tags: {
    id: number;
    name: string;
  }[];
}

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
  const [quantity, setQuantity] = useState(1);

  const { goods_id } = useParams();

  const { status, data } = useQuery({
    queryKey: ['goodsPage', goods_id],
    queryFn: async () => {
      const response = await fetch(`/api/product/${goods_id}`, {
        headers: {
          Accept: 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });
      RouteOnNotOK(response, navigate);
      return (await response.json()) as IProduct;
    },
    enabled: true,
    retry: false,
    refetchOnWindowFocus: false,
  });

  const addToCart = async () => {
    const resp = await fetch(`/api/buyer/cart/product/${goods_id}`, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token}`,
        accept: 'application/json',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        quantity: quantity,
      }),
    });
    if (!resp.ok) {
      RouteOnNotOK(resp);
    } else {
      alert('Success');
    }
  };

  if (goods_id === undefined) {
    return <NotFound />;
  }

  if (status !== 'success') {
    return <CheckFetchStatus status={status} />;
  }

  const handleAdd = () => {
    if (quantity < data.stock) {
      setQuantity(quantity + 1);
    }
  };
  const handleMinus = () => {
    if (quantity > 1) {
      setQuantity(quantity - 1);
    }
  };
  const handleQuantityChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const newQuantity = Number(e.target.value);
    if (!isNaN(newQuantity) && newQuantity > 0 && newQuantity <= data.stock) {
      setQuantity(newQuantity);
    }
  };

  const expireDate = formatDate(data.expire_date);
  return (
    <div style={{ padding: '55px 12% 0 12%' }}>
      <Row>
        <Col xs={12} md={5} style={LeftBgStyle}>
          <div className='flex_wrapper' style={{ padding: '0 8% 10% 8%' }}>
            <div style={{ overflow: ' hidden' }}>
              <img src={data.product_image_url} alt='File preview' style={GoodsImgStyle} />
            </div>

            <Row xs='auto'>
              {data.tags.map((tag, index) => (
                <Col style={tagStyle} className='center' key={index}>
                  {tag.name}
                </Col>
              ))}
            </Row>

            <h4 style={{ paddingTop: '30px', color: 'black', marginBottom: '5px' }}>
              $ {data.price} NTD
            </h4>
            {data.stock != 0 ? (
              <div>
                <span style={{ color: 'black' }}>{data.stock} available</span>
                <hr style={{ opacity: '1' }} />
                <Row>
                  <Col xs={3} onClick={handleMinus} className='pointer'>
                    <div className='quantity_f pointer center'>-</div>
                  </Col>

                  <Col xs={6} className='center'>
                    <div>
                      <input
                        type='text'
                        placeholder={`Quantity: ${quantity}`}
                        className='quantity_box'
                        value={quantity}
                        onChange={handleQuantityChange}
                        style={{ textAlign: 'center' }}
                      />
                    </div>
                  </Col>

                  <Col xs={3} onClick={handleAdd} className='pointer'>
                    <div className='quantity_f pointer center'>+</div>
                  </Col>
                </Row>
                <TButton text='Add to cart' action={addToCart} />
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

            <p>
              {data.description}
              <br />
              <br />
              Enjoy at its Freshest : {expireDate}
            </p>

            <hr className='hr' />
            <Row>
              <Col xs={6} className='center'>
                <UserItem img_path={data.shop_image_url} name={data.shop_name} />
              </Col>
              <Col xs={6}>
                <TButton text='View Shop' action={`/shop/${data.seller_name}`} />
              </Col>
            </Row>
          </div>
        </Col>
      </Row>
    </div>
  );
};

export default EachGoods;
