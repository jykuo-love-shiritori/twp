import 'bootstrap/dist/css/bootstrap.min.css';
import { useState } from 'react';
import { Col, Modal, Row } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCircleXmark } from '@fortawesome/free-regular-svg-icons';
import { IconProp } from '@fortawesome/fontawesome-svg-core';
import CouponItemTemplate from '@components/CouponItemTemplate';
import { useQuery } from '@tanstack/react-query';
import { CheckFetchStatus } from '@lib/Status';

interface ICouponItemShow {
  id: number;
  scope: 'global' | 'shop';
  name: string;
  type: 'percentage' | 'fixed' | 'shipping';
  discount: number;
  expire_date: string;
}

interface IShopCouponDetail {
  coupon_info: {
    description: string;
    discount: number;
    expire_date: string;
    name: string;
    scope: 'global' | 'shop';
    start_date: string;
    type: 'percentage' | 'fixed' | 'shipping';
  };
  tags: [
    {
      name: string;
      tag_id: number;
    },
  ];
}

interface IGlobalCouponDetail {
  description: string;
  discount: number;
  expire_date: string;
  id: number;
  name: string;
  scope: 'global' | 'shop';
  start_date: string;
  type: 'percentage' | 'fixed' | 'shipping';
}

const ModalShopCouponItem = ({ data }: { data: ICouponItemShow }) => {
  const [show, setShow] = useState(false);
  const handleShow = () => {
    refetch();
    setShow(true);
  };
  const handleClose = () => setShow(false);

  const {
    data: detailCouponData,
    status,
    refetch,
  } = useQuery({
    queryKey: ['sellerGetCouponDetail', data.id],
    queryFn: async () => {
      const resp = await fetch(`/api/seller/coupon/${data.id}`, {
        method: 'GET',
        headers: {
          accept: 'application/json',
        },
      });
      const response = await resp.json();
      if (!resp.ok) {
        alert(response.message);
      } else {
        return response;
      }
    },
    select: (data) => data as IShopCouponDetail,
    enabled: true,
    refetchOnWindowFocus: false,
  });

  if (status !== 'success') {
    return <CheckFetchStatus status={status} />;
  }

  const startDate = new Date(detailCouponData.coupon_info.start_date);
  const expDate = new Date(detailCouponData.coupon_info.expire_date);

  return (
    <>
      <div style={{ cursor: 'pointer' }} onClick={handleShow}>
        <CouponItemTemplate
          data={{
            name: detailCouponData.coupon_info.name,
            type: detailCouponData.coupon_info.type,
            discount: detailCouponData.coupon_info.discount,
            expire_date: detailCouponData.coupon_info.expire_date,
          }}
        />
      </div>
      <Modal show={show} onHide={handleClose} centered className='coupon_modal'>
        <Modal.Header>
          <div className='title' style={{ whiteSpace: 'nowrap' }}>
            Shop Coupon Detail
          </div>
          <div className='right' style={{ width: '100%', cursor: 'pointer' }}>
            <FontAwesomeIcon icon={faCircleXmark as IconProp} size='2x' onClick={handleClose} />
          </div>
        </Modal.Header>
        <Modal.Body>
          <Row>
            <Col xs={12} className='center' style={{ padding: '4% 0 0 0' }}>
              <div style={{ minWidth: '50%' }}>
                <CouponItemTemplate
                  data={{
                    name: detailCouponData.coupon_info.name,
                    type: detailCouponData.coupon_info.type,
                    discount: detailCouponData.coupon_info.discount,
                    expire_date: detailCouponData.coupon_info.expire_date,
                  }}
                />
              </div>
            </Col>
            <Col xs={12} style={{ paddingTop: '4%' }}>
              <p>{detailCouponData.coupon_info.description}</p>
            </Col>
            <Col xs={12} style={{ fontSize: '20px', color: '#ffffff7f' }}>
              Valid Period
            </Col>
            <Col xs={12} style={{ fontSize: '16px' }}>
              <Row>
                <Col xs='auto' style={{ paddingRight: '0', fontWeight: '500', color: '#ffffff7f' }}>
                  From
                </Col>
                <Col xs='auto' style={{ paddingRight: '0', fontWeight: '700', color: 'white' }}>
                  {`${startDate.getFullYear()}/${startDate.getMonth() + 1}/${startDate.getDate()}`}
                </Col>
                <Col xs='auto' style={{ paddingRight: '0', fontWeight: '500', color: '#ffffff7f' }}>
                  to
                </Col>
                <Col xs='auto' style={{ paddingRight: '0', fontWeight: '700', color: 'white' }}>
                  {`${expDate.getFullYear()}/${expDate.getMonth() + 1}/${expDate.getDate()} `}
                </Col>
              </Row>
            </Col>
            <Col xs={12} style={{ paddingTop: '4%' }}>
              <div style={{ display: 'flex', flexDirection: 'row', flexWrap: 'wrap' }}>
                {detailCouponData.tags.map((tag, index) => {
                  return (
                    <div
                      className='center'
                      key={index}
                      style={{
                        backgroundColor: 'var(--button_light)',
                        borderRadius: '30px',
                        padding: '1% 4% 1% 4%',
                        margin: '1% 2% 1% 0',
                      }}
                    >
                      {tag.name}
                    </div>
                  );
                })}
              </div>
            </Col>
          </Row>
        </Modal.Body>
      </Modal>
    </>
  );
};

const ModalGlobalCouponItem = ({ data }: { data: ICouponItemShow }) => {
  const [show, setShow] = useState(false);
  const handleShow = () => {
    refetch();
    setShow(true);
  };
  const handleClose = () => setShow(false);

  const {
    data: detailCouponData,
    status,
    refetch,
  } = useQuery({
    queryKey: ['adminGetCouponDetail', data.id],
    queryFn: async () => {
      const resp = await fetch(`/api/admin/coupon/${data.id}`, {
        method: 'GET',
        headers: {
          accept: 'application/json',
        },
      });
      const response = await resp.json();
      if (!resp.ok) {
        alert(response.message);
      } else {
        return response;
      }
    },
    select: (data) => data as IGlobalCouponDetail,
    enabled: true,
    refetchOnWindowFocus: false,
  });

  if (status !== 'success') {
    return <CheckFetchStatus status={status} />;
  }

  const startDate = new Date(detailCouponData.start_date);
  const expDate = new Date(detailCouponData.expire_date);

  return (
    <>
      <div style={{ cursor: 'pointer' }} onClick={handleShow}>
        <CouponItemTemplate data={data} />
      </div>
      <Modal show={show} onHide={handleClose} centered className='coupon_modal'>
        <Modal.Header>
          <div className='title' style={{ whiteSpace: 'nowrap' }}>
            Global Coupon Detail
          </div>
          <div className='right' style={{ width: '100%', cursor: 'pointer' }}>
            <FontAwesomeIcon icon={faCircleXmark as IconProp} size='2x' onClick={handleClose} />
          </div>
        </Modal.Header>
        <Modal.Body>
          <Row>
            <Col xs={12} className='center' style={{ padding: '4% 0 0 0' }}>
              <div style={{ minWidth: '50%' }}>
                <CouponItemTemplate
                  data={{
                    name: detailCouponData.name,
                    type: detailCouponData.type,
                    discount: detailCouponData.discount,
                    expire_date: detailCouponData.expire_date,
                  }}
                />
              </div>
            </Col>
            <Col xs={12} style={{ paddingTop: '4%' }}>
              <p>{detailCouponData.description}</p>
            </Col>
            <Col xs={12} style={{ fontSize: '20px', color: '#ffffff7f' }}>
              Valid Period
            </Col>
            <Col xs={12} style={{ fontSize: '16px' }}>
              <Row>
                <Col xs='auto' style={{ paddingRight: '0', fontWeight: '500', color: '#ffffff7f' }}>
                  From
                </Col>
                <Col xs='auto' style={{ paddingRight: '0', fontWeight: '700', color: 'white' }}>
                  {`${startDate.getFullYear()}/${startDate.getMonth() + 1}/${startDate.getDate()}`}
                </Col>
                <Col xs='auto' style={{ paddingRight: '0', fontWeight: '500', color: '#ffffff7f' }}>
                  to
                </Col>
                <Col xs='auto' style={{ paddingRight: '0', fontWeight: '700', color: 'white' }}>
                  {`${expDate.getFullYear()}/${expDate.getMonth() + 1}/${expDate.getDate()}`}
                </Col>
              </Row>
            </Col>
          </Row>
        </Modal.Body>
      </Modal>
    </>
  );
};

const CouponItemShow = ({ data }: { data: ICouponItemShow }) => {
  return (
    <>
      {data.scope === 'global' ? (
        <ModalGlobalCouponItem data={data} />
      ) : (
        <ModalShopCouponItem data={data} />
      )}
    </>
  );
};

export default CouponItemShow;
