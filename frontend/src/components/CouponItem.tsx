import 'bootstrap/dist/css/bootstrap.min.css';
import { useState } from 'react';
import { Col, Modal, Row } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCircleXmark } from '@fortawesome/free-regular-svg-icons';
import { IconProp } from '@fortawesome/fontawesome-svg-core';
import CouponItemTemplate from './CouponItemTemplate';
import TButton from './TButton';

interface WrapCouponProps {
  data: {
    id: number;
    type: 'percentage' | 'fixed' | 'shipping';
    name: string;
    description: string;
    discount: number;
    start_date: string;
    expire_date: string;
    tags: {
      name: string;
    }[];
  };
}

const LinkCouponItem = ({ data }: WrapCouponProps) => {
  return (
    <Link className='none' to={`${window.location.pathname}/${data.id}`}>
      <div style={{ cursor: 'pointer' }}>
        <CouponItemTemplate data={data} />
      </div>
    </Link>
  );
};

const ModalCouponItem = ({ data }: WrapCouponProps) => {
  const [show, setShow] = useState(false);
  const handleShow = () => setShow(true);
  const handleClose = () => setShow(false);
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
              Name
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
              <p>{data.description}</p>
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
                  {data.start_date}
                </Col>
                <Col xs='auto' style={{ paddingRight: '0', fontWeight: '500', color: '#ffffff7f' }}>
                  to
                </Col>
                <Col xs='auto' style={{ paddingRight: '0', fontWeight: '700', color: 'white' }}>
                  {data.expire_date}
                </Col>
              </Row>
            </Col>
            <Col xs={12} style={{ paddingTop: '4%' }}>
              <div style={{ display: 'flex', flexDirection: 'row', flexWrap: 'wrap' }}>
                {data.tags.map((tag, index) => {
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

const CouponItem = ({ data }: WrapCouponProps) => {
  // TODO: path.include 'seller/' change to 'seller' once implemented number [sellerID] on path
  return (
    <>
      {window.location.pathname.includes('seller/') ||
      window.location.pathname.includes('admin') ? (
        <LinkCouponItem data={data} />
      ) : (
        <ModalCouponItem data={data} />
      )}
    </>
  );
};

export default CouponItem;
