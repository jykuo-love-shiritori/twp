import { Col, Row } from 'react-bootstrap';
import CouponItem from '@components/CouponItem';
import TButton from '@components/TButton';
import { useNavigate } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';
import { RouteOnNotOK } from '@lib/Functions';
import { CheckFetchStatus } from '@lib/Status';

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
  const { data: fetchedData, status } = useQuery({
    queryKey: ['adminGetGlobalCoupons'],
    queryFn: async () => {
      const resp = await fetch('/api/admin/coupons', {
        method: 'GET',
        headers: {
          accept: 'application/json',
        },
      });
      if (!resp.ok) {
        RouteOnNotOK(resp, navigate);
      } else {
        const response = await resp.json();
        return response;
      }
    },
    select: (data) => data as ICoupon[],
    enabled: true,
    refetchOnWindowFocus: false,
  });

  if (status !== 'success') {
    return <CheckFetchStatus status={status} />;
  }

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
                    <CouponItem
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
                    <CouponItem
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
    </div>
  );
};

export default ManageAdminCoupons;
