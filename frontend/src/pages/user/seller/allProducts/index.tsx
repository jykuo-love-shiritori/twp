import { Col, Row } from 'react-bootstrap';
import { useQuery } from '@tanstack/react-query';

import SellerGoodsItem from '@components/SellerGoodsItem';
import TButton from '@components/TButton';
import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import { GoodsItemProps } from '@components/GoodsItem';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { useAuth } from '@lib/Auth';
import { useState } from 'react';
import Pagination from '@components/Pagination';

const Products = () => {
  const token = useAuth();
  const navigate = useNavigate();
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

  const { status, data: sellerShopData } = useQuery({
    queryKey: ['sellerShopView', searchParams.toString()],
    queryFn: async () => {
      const resp = await fetch(`/api/seller/product?` + searchParams.toString(), {
        headers: {
          Accept: 'application/json',
          Authorization: `Bearer ${token}`,
        },
      });
      if (!resp.ok) {
        RouteOnNotOK(resp, navigate);
        return [];
      }
      const response = await resp.json();
      if (response.length === itemLimit + 1) {
        setIsMore(true);
        response.pop();
      } else {
        setIsMore(false);
      }
      return response;
    },
    select: (data) => data as GoodsItemProps[],
    enabled: true,
    retry: false,
    refetchOnWindowFocus: false,
  });

  if (status != 'success') {
    return <CheckFetchStatus status={status} />;
  }

  return (
    <div>
      <Row>
        <Col sm={12} md={8}>
          <div className='title'>All products</div>
        </Col>
        <Col sm={12} md={4}>
          <div style={{ padding: '20px 0 0 0' }}>
            <TButton text='Add New Item' action='/user/seller/manageProducts/new' />
          </div>
        </Col>
      </Row>
      <hr className='hr' />
      <Row>
        {sellerShopData.map((data: GoodsItemProps, index: number) => {
          return (
            <Col xs={6} md={3} key={index}>
              <SellerGoodsItem id={data.id} name={data.name} image_url={data.image_url} />
            </Col>
          );
        })}
      </Row>
      <div className='center'>
        <Pagination limit={itemLimit} isMore={isMore} />
      </div>
    </div>
  );
};

export default Products;
