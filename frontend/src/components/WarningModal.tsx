import { Col, Modal, Row } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faCircleXmark } from '@fortawesome/free-regular-svg-icons';
import { IconProp } from '@fortawesome/fontawesome-svg-core';

interface WarningModalProps {
  text: string;
  show: boolean;
  onHide: () => void;
}

const WarningModal = ({ text, show, onHide }: WarningModalProps) => {
  return (
    <Modal show={show} onHide={onHide} centered className='coupon_modal'>
      <Modal.Header style={{ border: 'none' }}>
        <Row className='center_vertical' style={{ width: '100%' }}>
          <Col xs={8} md={11} className='title' style={{ padding: '10px' }}>
            <div style={{ fontSize: '40px' }}>Warning</div>
          </Col>
          <Col xs={4} md={1} className='right' style={{ padding: '0' }}>
            <FontAwesomeIcon
              icon={faCircleXmark as IconProp}
              size='2x'
              onClick={onHide}
              style={{ cursor: 'pointer' }}
            />
          </Col>
        </Row>
      </Modal.Header>
      <hr className='hr' style={{ margin: '0 0 5px 0' }} />
      <Modal.Body>{text}</Modal.Body>
    </Modal>
  );
};

export default WarningModal;
