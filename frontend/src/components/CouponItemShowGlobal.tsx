import 'bootstrap/dist/css/bootstrap.min.css';
import { useState } from 'react';
import { Col, Modal, Row } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCircleXmark } from '@fortawesome/free-regular-svg-icons';
import { IconProp } from '@fortawesome/fontawesome-svg-core';
import CouponItemTemplate from '@components/CouponItemTemplate';

interface ICouponItemShowGlobal {
  description: string;
  discount: number;
  expire_date: string;
  id: number;
  name: string;
  scope: 'global' | 'shop';
  start_date: string;
  type: 'percentage' | 'fixed' | 'shipping';
}

const CouponItemShowGlobal = ({ data }: { data: ICouponItemShowGlobal }) => {
  const [show, setShow] = useState(false);
  const handleShow = () => setShow(true);
  const handleClose = () => setShow(false);

  const startDate = new Date(data.start_date);
  const expDate = new Date(data.expire_date);

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
                    name: data.name,
                    type: data.type,
                    discount: data.discount,
                    expire_date: data.expire_date,
                  }}
                />
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

export default CouponItemShowGlobal;
