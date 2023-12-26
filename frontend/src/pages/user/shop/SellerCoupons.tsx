import { Col, Row } from 'react-bootstrap';
import { useQuery } from '@tanstack/react-query';
import { useNavigate } from 'react-router-dom';
import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import CouponItem from '@components/CouponItem';

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

  // TODO
  // const {sellerName} = useParams();
  const sellerName = 'user1';

  const { data: fetchedData, status } = useQuery({
    queryKey: ['GetShopCoupons'],
    queryFn: async () => {
      const resp = await fetch(`/api/shop/${sellerName}/coupon?offset=0&limit=10`, {
        method: 'GET',
        headers: {
          accept: 'application/json',
        },
      });
      if (!resp.ok) {
        RouteOnNotOK(resp, navigate);
      } else {
        return await resp.json();
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
        <Col md={12}>
          <div className='title'>All Coupon</div>
        </Col>
        <hr />
        <Row>
          <div className='disappear_phone'>
            <Row>
              {fetchedData.map((data, index) => {
                return (
                  <Col md={4} xl={3} key={index} style={{ padding: '2%' }}>
                    <CouponItem
                      data={{
                        discount: data.discount,
                        expire_date: data.expire_date,
                        id: data.id,
                        name: data.name,
                        scope: data.scope,
                        type: data.type,
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
                        discount: data.discount,
                        expire_date: data.expire_date,
                        id: data.id,
                        name: data.name,
                        scope: data.scope,
                        type: data.type,
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

export default SellerCoupons;
