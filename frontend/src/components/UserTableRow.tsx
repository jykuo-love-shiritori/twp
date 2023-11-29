import { Row, Col } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTrash } from '@fortawesome/free-solid-svg-icons';
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
  //delete checkbox
  const ColStyleOn = {
    fontSize: '16px',
    padding: '1% 1% 1% 1%',
  };
  const ColStyleOff = {
    fontSize: '16px',
    padding: '1% 1% 1% 1%',
    color: 'var(--layout)',
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

  const UserTableRowSmall = () => {
    return (
      <>
        <hr />
        <Row style={{ margin: '5px' }}>
          <Col xs={4} md={4} className={'center'} style={currentStyle.style}>
            <img src={data.iconUrl} className='user_img' />
          </Col>
          <Col xs={7} md={7} className={'left center_vertical'} style={currentStyle.style}>
            <Row>
              <p style={currentStyle.style}>
                {'name: ' + data.name} <br />
                {'email: ' + data.email} <br />
                {'created: ' + data.createDate}
              </p>
            </Row>
          </Col>
          <Col xs={1} md={1}>
            {/* delete button */}
            <div
              className='center center_vertical'
              style={{
                height: '100%',
                ...currentStyle.style,
              }}
            >
              <FontAwesomeIcon
                icon={faTrash}
                size='2x'
                onClick={toggleIsDelete}
                style={{ cursor: 'pointer' }}
              />
            </div>
          </Col>
        </Row>
      </>
    );
  };

  const UserTableRowBig = () => {
    return (
      <Row style={{ padding: '0 0 0 0' }}>
        <Col md={1} className={'center'} style={currentStyle.style}>
          <img src={data.iconUrl} className='user_img' />
        </Col>
        <Col md={3} className={'left center_vertical'} style={currentStyle.style}>
          <h4>{data.name}</h4>
        </Col>
        <Col md={5} className={'left center_vertical'} style={currentStyle.style}>
          <h4>{data.email}</h4>
        </Col>
        <Col md={2} className={'left center_vertical'} style={currentStyle.style}>
          <h4>{data.createDate}</h4>
        </Col>
        <Col md={1} className={'center center_vertical'} style={currentStyle.style}>
          <FontAwesomeIcon
            icon={faTrash}
            size='2x'
            onClick={toggleIsDelete}
            style={{ cursor: 'pointer' }}
          />
        </Col>
      </Row>
    );
  };

  return (
    <>
      <div className='disappear_tablet disappear_phone'>
        <UserTableRowBig />
      </div>
      <div className='disappear_desktop'>
        <UserTableRowSmall />
      </div>
    </>
  );
};

export default UserTableRow;
