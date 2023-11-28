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

  //resize
  const [winSize, setWinSize] = useState(window.innerWidth);
  window.addEventListener('resize', () => {
    setWinSize(window.innerWidth);
  });

  if (winSize >= 1024) {
    return (
      <Row style={{ padding: '0 0 0 0' }}>
        <Col md={1} className={'center'} style={currentStyle.style}>
          <img src={data.iconUrl} className='user_img' />
        </Col>
        <Col md={2} className={'left center_vertical'} style={currentStyle.style}>
          <h4>{data.name}</h4>
        </Col>
        <Col md={4} className={'left center_vertical'} style={currentStyle.style}>
          <h4>{data.email}</h4>
        </Col>
        <Col md={3} className={'left center_vertical'} style={currentStyle.style}>
          <h4>{data.createDate}</h4>
        </Col>
        <Col md={1} className={'center center_vertical'} style={currentStyle.style}>
          <div onClick={toggleIsAdmin}>
            {isCheckedAdmin ? (
              <FontAwesomeIcon icon={faGear} size='2x' />
            ) : (
              <FontAwesomeIcon icon={faGear} size='2x' color='black' />
            )}
          </div>
        </Col>
        <Col md={1} className={'center center_vertical'} style={currentStyle.style}>
          <div onClick={toggleIsDelete}>
            <FontAwesomeIcon icon={faTrash} size='2x' />
          </div>
        </Col>
      </Row>
    );
  } else {
    const ButtonFlexStyle = { display: 'flex', flexGrow: '1', justifyContent: 'center' };
    return (
      <>
        <hr />
        <Row style={{ margin: '5px' }}>
          <Col xs={4} md={4} className={'center'} style={currentStyle.style}>
            <img src={data.iconUrl} className='user_img' />
          </Col>
          <Col xs={6} md={6} className={'left center_vertical'} style={currentStyle.style}>
            <Row>
              <h5>{'name: ' + data.name}</h5>
              <h5>{'email: ' + data.email}</h5>
              <h5>{'created: ' + data.createDate}</h5>
            </Row>
          </Col>
          <Col xs={2} md={2}>
            <div style={{ display: 'flex', flexDirection: 'column', height: '100%' }}>
              <div style={{ ...currentStyle.style, ...ButtonFlexStyle }}>
                <div className='center center_vertical' onClick={toggleIsAdmin}>
                  {isCheckedAdmin ? (
                    <FontAwesomeIcon icon={faGear} size='2x' />
                  ) : (
                    <FontAwesomeIcon icon={faGear} size='2x' color='black' />
                  )}
                </div>
              </div>
              <div style={{ ...currentStyle.style, ...ButtonFlexStyle }}>
                <div className='center center_vertical' onClick={toggleIsDelete}>
                  <FontAwesomeIcon icon={faTrash} size='2x' />
                </div>
              </div>
            </div>
          </Col>
        </Row>
      </>
    );
  }
};

export default UserTableRow;
