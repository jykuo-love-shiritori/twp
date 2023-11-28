import { Row, Col } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faGear, faTrash } from '@fortawesome/free-solid-svg-icons';
import { useState } from 'react';

type UserTableRowProps = {
  data: {
    iconUrl: string;
    name: string;
    email: string;
    createDate: string;
    isAdmin: boolean;
  };
};

const UserTableRow = ({ data }: UserTableRowProps) => {
  //admin checkbox
  const [isCheckedAdmin, setIsCheckedAdmin] = useState(data.isAdmin);
  const toggleIsAdmin = () => {
    setIsCheckedAdmin(!isCheckedAdmin);
  };

  //delete checkbox
  const ColStyleOn = {
    fontSize: '16px',
    padding: '1% 1% 0% 1%',
  };
  const ColStyleOff = {
    fontSize: '16px',
    padding: '1% 1% 0% 1%',
    color: 'var(--button_dark)',
    textDecoration: 'line-through',
  };
  const [currentStyle, setCurrentStyle] = useState({ state: true, style: ColStyleOn });
  const toggleIsDelete = () => {
    setCurrentStyle(
      currentStyle.state
        ? { state: false, style: ColStyleOff }
        : { state: true, style: ColStyleOn },
    );
  };

  return (
    <div className='manage_user_table_row'>
      <Row>
        <Col md={1} xs={2} className='center' style={currentStyle.style}>
          <img src={data.iconUrl} className='user_img' />
        </Col>
        <Col md={2} xs={10} className='left center_vertical' style={currentStyle.style}>
          <h4>{data.name}</h4>
        </Col>
        <Col md={4} xs={12} className='left center_vertical' style={currentStyle.style}>
          <h4>{data.email}</h4>
        </Col>
        <Col md={3} xs={12} className='left center_vertical' style={currentStyle.style}>
          <h4>{data.createDate}</h4>
        </Col>
        <Col md={1} xs={6} className='center center_vertical' style={currentStyle.style}>
          <div onClick={toggleIsAdmin}>
            {isCheckedAdmin ? (
              <FontAwesomeIcon icon={faGear} size='2x' />
            ) : (
              <FontAwesomeIcon icon={faGear} size='2x' color='black' />
            )}
          </div>
        </Col>
        <Col md={1} xs={6} className='center center_vertical' style={currentStyle.style}>
          <div onClick={toggleIsDelete}>
            <FontAwesomeIcon icon={faTrash} size='2x' />
          </div>
        </Col>
      </Row>
    </div>
  );
};

export default UserTableRow;
