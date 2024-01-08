import { Col, Row } from 'react-bootstrap';
import { useQuery } from '@tanstack/react-query';
import { useNavigate, useParams, useSearchParams } from 'react-router-dom';
import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import { useAuth } from '@lib/Auth';
import { useState } from 'react';
import CouponItemShow from '@components/CouponItemShow';
import Pagination from '@components/Pagination';

interface ICoupon {
  description: string;
  discount: number;
  expire_date: string;
  id: number;
  name: string;
  scope: 'global' | 'shop';
  start_date: string;
  type: 'percentage' | 'fixed' | 'shipping';
}

const SellerCoupons = () => {
  const navigate = useNavigate();
  const token = useAuth();
  const { sellerName } = useParams();
  const [searchParams, setSearchParams] = useSearchParams();
  const [isMore, setIsMore] = useState(true);

  const itemLimit = 12;

  if (!searchParams.has('offset') || Number(searchParams.get('limit')) !== itemLimit + 1) {
    const newSearchParams = new URLSearchParams({
      offset: '0',
      limit: (itemLimit + 1).toString(),
    });
    setSearchParams(newSearchParams, { replace: true });
  }

  const {
    data: CouponsData,
    status: fetchCouponsStatus,
    refetch,
  } = useQuery({
    queryKey: ['GetShopCoupons', searchParams.toString()],
    queryFn: async () => {
      const resp = await fetch(`/api/shop/${sellerName}/coupon?` + searchParams.toString(), {
        method: 'GET',
        headers: {
          Authorization: `Bearer ${token}`,
          accept: 'application/json',
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
    select: (data) => data as ICoupon[],
    enabled: true,
    retry: false,
    refetchOnWindowFocus: false,
  });

  if (fetchCouponsStatus !== 'success') {
    return <CheckFetchStatus status={fetchCouponsStatus} />;
  }

  const globalCoupons = CouponsData.filter((coupon) => coupon.scope === 'global');
  const shopCoupons = CouponsData.filter((coupon) => coupon.scope === 'shop');

  return (
    <div>
      <Row>
        <Col md={12}>
          <div className='title'>Global Coupon</div>
        </Col>
        <hr className='hr' />
        {globalCoupons.map((data, index) => {
          return (
            <Col xs={12} md={4} xl={3} key={index} style={{ padding: '2%' }}>
              <CouponItemShow data={data} />
            </Col>
          );
        })}
        <Col md={12} style={{ paddingTop: '5%' }}>
          <div className='title'>Shop Coupon</div>
        </Col>
        <hr className='hr' />
        {shopCoupons.map((data, index) => {
          return (
            <Col xs={12} md={4} xl={3} key={index} style={{ padding: '2%' }}>
              <CouponItemShow data={data} />
            </Col>
          );
        })}
      </Row>
      <div className='center'>
        <Pagination
          searchParams={searchParams}
          setSearchParams={setSearchParams}
          refetch={refetch}
          limit={itemLimit}
          isMore={isMore}
        />
      </div>
    </div>
  );
};

export default SellerCoupons;
