import { Col, Row } from 'react-bootstrap';
import { useQuery } from '@tanstack/react-query';
import { useNavigate, useParams, useSearchParams } from 'react-router-dom';

import GoodsItem from '@components/GoodsItem';
import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import { useAuth } from '@lib/Auth';
import { useState } from 'react';
import Pagination from '@components/Pagination';

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
  const [searchParams, setSearchParams] = useSearchParams();
  const [isMore, setIsMore] = useState(true);

  const itemLimit = 8;

  if (!searchParams.has('offset') || Number(searchParams.get('limit')) !== itemLimit + 1) {
    const newSearchParams = new URLSearchParams({
      offset: '0',
      limit: (itemLimit + 1).toString(),
    });
    setSearchParams(newSearchParams, { replace: true });
  }

  const { status, data } = useQuery({
    queryKey: ['getShopInfoForProducts', sellerName, searchParams.toString()],
    queryFn: async () => {
      const resp = await fetch(`/api/shop/${sellerName}?` + searchParams.toString(), {
        method: 'GET',
        headers: {
          Accept: 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });
      if (!resp.ok) {
        RouteOnNotOK(resp, navigate);
      }
      const response = (await resp.json()) as IShop;
      if (response.products.length === itemLimit + 1) {
        setIsMore(true);
        response.products.pop();
      } else {
        setIsMore(false);
      }
      return response;
    },
    enabled: true,
    refetchOnWindowFocus: false,
    retry: false,
  });

  if (status != 'success') {
    return <CheckFetchStatus status={status} />;
  }

  const products = Array.isArray(data.products) ? data.products : [];

  return (
    <div>
      <div className='title'>All products</div>
      <hr className='hr' />
      <Row>
        {products.length > 0 ? (
          products.map((data, index) => {
            return (
              <Col xs={6} md={3} key={index}>
                <GoodsItem id={data.id} name={data.name} image_url={data.image_url} />
              </Col>
            );
          })
        ) : (
          <h3>No shop product 😢</h3>
        )}
      </Row>
      <div className='center' style={{ padding: '2% 0px' }}>
        <Pagination limit={itemLimit} isMore={isMore} />
      </div>
    </div>
  );
};

export default Shop;
