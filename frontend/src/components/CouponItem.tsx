import 'bootstrap/dist/css/bootstrap.min.css';
import { useState } from 'react';
import { Col, Modal, Row } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCircleXmark } from '@fortawesome/free-regular-svg-icons';
import { IconProp } from '@fortawesome/fontawesome-svg-core';
import CouponItemTemplate from '@components/CouponItemTemplate';
import TButton from '@components/TButton';
import { useQuery } from '@tanstack/react-query';
import { CheckFetchStatus } from '@lib/Status';

interface ICouponItem {
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

const LinkCouponItem = ({ data }: { data: ICouponItem }) => {
  return (
    <Link className='none' to={`${window.location.pathname}/${data.id}`}>
      <div style={{ cursor: 'pointer' }}>
        <CouponItemTemplate data={data} />
      </div>
    </Link>
  );
};

const ModalShopCouponItem = ({ data }: { data: ICouponItem }) => {
  const [show, setShow] = useState(false);
  const handleShow = () => {
    refetch();
    setShow(true);
  };
  const handleClose = () => setShow(false);

  const {
    data: fetchedData,
    status,
    refetch,
  } = useQuery({
    queryKey: ['sellerGetCouponDetail'],
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

  return (
    <>
      <div style={{ cursor: 'pointer' }} onClick={handleShow}>
        <CouponItemTemplate
          data={{
            name: fetchedData.coupon_info.name,
            type: fetchedData.coupon_info.type,
            discount: fetchedData.coupon_info.discount,
            expire_date: fetchedData.coupon_info.expire_date,
          }}
        />
      </div>
      <Modal show={show} onHide={handleClose} centered className='coupon_modal'>
        <Modal.Header>
          <div className='right' style={{ width: '100%', cursor: 'pointer' }}>
            <FontAwesomeIcon icon={faCircleXmark as IconProp} size='2x' onClick={handleClose} />
          </div>
        </Modal.Header>
        <Modal.Body>
          <Row>
            <Col xs={3} md={2} className='center'>
              <img src='/placeholder/person.png' className='user_img' />
            </Col>
            <Col xs={4} md={6} className='center_vertical left'>
              {fetchedData.coupon_info.name}
            </Col>
            <Col xs={5} md={4}>
              <TButton text='ViewShop' />
            </Col>
            <Col xs={12} className='center' style={{ padding: '4% 0 0 0' }}>
              <div style={{ minWidth: '50%' }}>
                <CouponItemTemplate data={data} />
              </div>
            </Col>
            <Col xs={12} style={{ paddingTop: '4%' }}>
              <p>{fetchedData.coupon_info.description}</p>
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
                  {fetchedData.coupon_info.start_date}
                </Col>
                <Col xs='auto' style={{ paddingRight: '0', fontWeight: '500', color: '#ffffff7f' }}>
                  to
                </Col>
                <Col xs='auto' style={{ paddingRight: '0', fontWeight: '700', color: 'white' }}>
                  {fetchedData.coupon_info.expire_date}
                </Col>
              </Row>
            </Col>
            <Col xs={12} style={{ paddingTop: '4%' }}>
              <div style={{ display: 'flex', flexDirection: 'row', flexWrap: 'wrap' }}>
                {fetchedData.tags.map((tag, index) => {
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

const ModalGlobalCouponItem = ({ data }: { data: ICouponItem }) => {
  const [show, setShow] = useState(false);
  const handleShow = () => {
    refetch();
    setShow(true);
  };
  const handleClose = () => setShow(false);

  const {
    data: fetchedData,
    status,
    refetch,
  } = useQuery({
    queryKey: ['adminGetCouponDetail'],
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
    enabled: false,
    refetchOnWindowFocus: false,
  });

  if (status !== 'success') {
    return <CheckFetchStatus status={status} />;
  }

  return (
    <>
      <div style={{ cursor: 'pointer' }} onClick={handleShow}>
        <CouponItemTemplate data={data} />
      </div>
      <Modal show={show} onHide={handleClose} centered className='coupon_modal'>
        <Modal.Header>
          <div className='right' style={{ width: '100%', cursor: 'pointer' }}>
            <FontAwesomeIcon icon={faCircleXmark as IconProp} size='2x' onClick={handleClose} />
          </div>
        </Modal.Header>
        <Modal.Body>
          <Row>
            <Col xs={3} md={2} className='center'>
              <img src='/placeholder/person.png' className='user_img' />
            </Col>
            <Col xs={4} md={6} className='center_vertical left'>
              {fetchedData.name}
            </Col>
            <Col xs={5} md={4}>
              <TButton text='ViewShop' />
            </Col>
            <Col xs={12} className='center' style={{ padding: '4% 0 0 0' }}>
              <div style={{ minWidth: '50%' }}>
                <CouponItemTemplate data={data} />
              </div>
            </Col>
            <Col xs={12} style={{ paddingTop: '4%' }}>
              <p>{fetchedData.description}</p>
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
                  {fetchedData.start_date}
                </Col>
                <Col xs='auto' style={{ paddingRight: '0', fontWeight: '500', color: '#ffffff7f' }}>
                  to
                </Col>
                <Col xs='auto' style={{ paddingRight: '0', fontWeight: '700', color: 'white' }}>
                  {fetchedData.expire_date}
                </Col>
              </Row>
            </Col>
          </Row>
        </Modal.Body>
      </Modal>
    </>
  );
};

const CouponItem = ({ data }: { data: ICouponItem }) => {
  // TODO: path.include 'seller/' change to 'seller' once implemented number [sellerID] on path
  return (
    <>
      {window.location.pathname.includes('seller/') ||
      window.location.pathname.includes('admin') ? (
        <LinkCouponItem data={data} />
      ) : data.scope === 'global' ? (
        <ModalGlobalCouponItem data={data} />
      ) : (
        <ModalShopCouponItem data={data} />
      )}
    </>
  );
};

export default CouponItem;
