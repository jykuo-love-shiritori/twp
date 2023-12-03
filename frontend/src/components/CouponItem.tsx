import 'bootstrap/dist/css/bootstrap.min.css';
import { useState } from 'react';
import { Col, Modal, Row } from 'react-bootstrap';
import { Link } from 'react-router-dom';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCircleXmark } from '@fortawesome/free-regular-svg-icons';
import TButton from './TButton';

interface CouponItemProps {
  data: {
    id: number;
    name: string;
    policy: string;
    date: string;
    tags: { name: string }[];
  };
}

const couponStyle = {
  backgroundColor: 'var(--button_dark)',
  boxShadow: '3px 5px 10px 0px rgba(0, 0, 0, 0.25)',
  borderRadius: '30px',
  padding: '5%',
  border: 'var(--border) solid 2px',
};

const PreCouponItem = ({ data }: CouponItemProps) => {
  return (
    <Row style={{ height: '100%', padding: '2% 0' }}>
      <Col xs={8} md={9} xl={8} className='center'>
        <div>
          <div className='center' style={{ fontSize: '16px', fontWeight: '700', color: 'white' }}>
            {data.name}
          </div>
          <div className='center' style={{ fontSize: '12px', fontWeight: '500', color: 'white' }}>
            {data.policy}
          </div>
        </div>
      </Col>
      <Col xs={4} md={3} xl={4} className='center' style={{ borderLeft: '2px dashed #AAAAAA' }}>
        <div>
          <div className='center' style={{ fontSize: '16px', fontWeight: '500', color: 'white' }}>
            exp
          </div>
          <div className='center' style={{ fontSize: '12px', fontWeight: '500', color: 'white' }}>
            {data.date}
          </div>
        </div>
      </Col>
    </Row>
  );
};

const LinkCouponItem = ({ data }: CouponItemProps) => {
  return (
    <Link className='none' to={`${window.location.pathname}/${data.id}`}>
      <div style={{ ...couponStyle, cursor: 'pointer' }}>
        <PreCouponItem data={data} />
      </div>
    </Link>
  );
};

const ModalCouponItem = ({ data }: CouponItemProps) => {
  const [show, setShow] = useState(false);
  const handleShow = () => setShow(true);
  const handleClose = () => setShow(false);
  return (
    <>
      <div style={{ ...couponStyle, cursor: 'pointer' }} onClick={handleShow}>
        <PreCouponItem data={data} />
      </div>
      <Modal show={show} onHide={handleClose} centered className='coupon_modal'>
        <Modal.Header>
          <div className='right' style={{ width: '100%' }}>
            <FontAwesomeIcon icon={faCircleXmark} size='2x' onClick={handleClose} />
          </div>
        </Modal.Header>
        <Modal.Body>
          <Row>
            <Col xs={3} md={2} className='center' style={{ paddingTop: '4%' }}>
              <img src='../../images/person.png' className='user_img' />
            </Col>
            <Col xs={4} md={6} className='center_vertical left' style={{ padding: '4% 0 0 0' }}>
              Name
            </Col>
            <Col xs={5} md={4} style={{ padding: '4% 0 0 0' }}>
              <TButton text='ViewShop' url={''} />
            </Col>
            <Col xs={12} className='center' style={{ padding: '4% 0 0 0' }}>
              <div style={{ ...couponStyle, minWidth: '50%' }}>
                <PreCouponItem data={data} />
              </div>
            </Col>
            <Col xs={12} style={{ paddingTop: '4%' }}>
              <p>
                Gather your loved ones and enjoy a 30% discount on your entire bill when you order
                dessert for four or more. It's the perfect excuse for a dessert night out.
              </p>
            </Col>
            <Col xs={12} style={{ fontSize: '20px', color: '#ffffff7f' }}>
              Expire at: {data.date}
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
                        padding: '2% 4% 2% 4%',
                        margin: '2% 2% 2% 0',
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

const CouponItem = ({ data }: CouponItemProps) => {
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
