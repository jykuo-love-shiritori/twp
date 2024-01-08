import { Col, Row } from 'react-bootstrap';
import { useQuery } from '@tanstack/react-query';
import { useNavigate, useParams } from 'react-router-dom';

import GoodsItem from '@components/GoodsItem';
import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import { useAuth } from '@lib/Auth';

interface IProduct {
  description: string;
  expire_date: string;
  id: number;
  image_url: string;
  name: string;
  price: number;
  sales: number;
  stock: number;
}

interface IShop {
  info: {
    description: string;
    image_url: string;
    name: string;
    seller_name: string;
  };
  products: IProduct[];
}

const Shop = () => {
  const token = useAuth();
  const navigate = useNavigate();
  const { sellerName } = useParams();

  const { status, data } = useQuery({
    queryKey: ['getShopInfoForProducts', sellerName],
    queryFn: async () => {
      const resp = await fetch(`/api/shop/${sellerName}?offset=${0}&limit=${8}`, {
        method: 'GET',
        headers: {
          Accept: 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });
      RouteOnNotOK(resp, navigate);
      return (await resp.json()) as IShop;
    },
    enabled: true,
    refetchOnWindowFocus: false,
    retry: false,
  });

  if (status != 'success') {
    return <CheckFetchStatus status={status} />;
  }

  const products = data.products;

  return (
    <div>
      <div className='title'>All products</div>
      <hr className='hr' />
      <Row>
        {products.map((data, index) => {
          return (
            <Col xs={6} md={3} key={index}>
              <GoodsItem id={data.id} name={data.name} image_url={data.image_url} />
            </Col>
          );
        })}
      </Row>
    </div>
  );
};

export default Shop;
