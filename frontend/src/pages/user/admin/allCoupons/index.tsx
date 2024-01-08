import { Col, Row } from 'react-bootstrap';
import CouponItemEdit from '@components/CouponItemEdit';
import TButton from '@components/TButton';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';
import { RouteOnNotOK } from '@lib/Status';
import { CheckFetchStatus } from '@lib/Status';
import { useAuth } from '@lib/Auth';
import { useState } from 'react';
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

const ManageAdminCoupons = () => {
  const navigate = useNavigate();
  const token = useAuth();
  const [searchParams, setSearchParams] = useSearchParams();
  const [isMore, setIsMore] = useState(true);

  const itemLimit = 6;

  if (!searchParams.has('offset') || Number(searchParams.get('limit')) !== itemLimit + 1) {
    const newSearchParams = new URLSearchParams({
      offset: '0',
      limit: (itemLimit + 1).toString(),
    });
    setSearchParams(newSearchParams, { replace: true });
  }

  const { data: fetchedData, status } = useQuery({
    queryKey: ['adminGetGlobalCoupons', searchParams.toString()],
    queryFn: async () => {
      const resp = await fetch('/api/admin/coupon?' + searchParams.toString(), {
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
      console.log(response.length);
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
    refetchOnWindowFocus: false,
    staleTime: 0,
    retry: false,
  });

  // console.log(status);

  if (status !== 'success') {
    return <CheckFetchStatus status={status} />;
  }

  // console.log(isMore);

  return (
    <div>
      <Row>
        {/* display title for desktop */}
        <div className='disappear_phone disappear_tablet' style={{ paddingTop: '6%' }}>
          <Row>
            <Col xl={6} className='left'>
              <div className='title'>All Global Coupon</div>
            </Col>
            <Col xl={3} />
            <Col xl={3} className='right'>
              <TButton text='New Coupon' action='/admin/manageCoupons/new' />
            </Col>
          </Row>
        </div>
        {/* display title for tablet */}
        <div className='disappear_desktop disappear_phone'>
          <Row>
            <Col md={8} className='left'>
              <div className='title'>All Global Coupon</div>
            </Col>
            <Col md={4} className='right'>
              <TButton text='New Coupon' action='/admin/manageCoupons/new' />
            </Col>
          </Row>
        </div>
        {/* display title for phone */}
        <div className='disappear_desktop disappear_tablet'>
          <Row>
            <Col xs={12} className='center' style={{ padding: '0' }}>
              <div className='title'>All Global Coupon</div>
            </Col>
            <Col />
            <Col xs={6} className='center' style={{ padding: '0 0 2% 0' }}>
              <TButton text='New Coupon' action='/admin/manageCoupons/new' />
            </Col>
            <Col />
          </Row>
        </div>

        <hr className='hr' />
        <Row>
          <div className='disappear_phone'>
            <Row>
              {fetchedData.map((data, index) => {
                return (
                  <Col md={6} xl={4} key={index} style={{ padding: '2%' }}>
                    <CouponItemEdit
                      data={{
                        id: data.id,
                        scope: data.scope,
                        name: data.name,
                        type: data.type,
                        discount: data.discount,
                        expire_date: data.expire_date,
                      }}
                    />
                  </Col>
                );
              })}
            </Row>
          </div>
          <div className='disappear_desktop disappear_tablet'>
            <Row>
              {fetchedData.map((data, index) => {
                return (
                  <Col xs={12} key={index} style={{ padding: '2% 10%' }}>
                    <CouponItemEdit
                      data={{
                        id: data.id,
                        scope: data.scope,
                        name: data.name,
                        type: data.type,
                        discount: data.discount,
                        expire_date: data.expire_date,
                      }}
                    />
                  </Col>
                );
              })}
            </Row>
          </div>
        </Row>
      </Row>
      <div className='center'>
        <Pagination limit={itemLimit} isMore={isMore} />
      </div>
    </div>
  );
};

export default ManageAdminCoupons;
