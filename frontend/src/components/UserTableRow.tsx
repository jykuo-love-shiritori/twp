import { Row, Col } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faTrash } from '@fortawesome/free-solid-svg-icons';

interface ICreditCard {
  CVV: string;
  name: string;
  card_number: string;
  expiry_date: string;
}

interface IUser {
  address: string;
  credit_card: [ICreditCard];
  email: string;
  enabled: true;
  icon_url: string;
  name: string;
  role: string;
  username: string;
}

const UserTableRow = ({ data, refresh }: { data: IUser; refresh: () => void }) => {
  const onDelete = async () => {
    const resp = await fetch(`/api/admin/user/${data.username}`, {
      method: 'DELETE',
      headers: { accept: 'application/json' },
    });
    if (!resp.ok) {
      const response = await resp.json();
      alert(response.message);
    } else {
      refresh();
    }
  };

  const UserTableRowSmall = () => {
    return (
      <>
        <hr />
        <Row style={{ margin: '5px', fontSize: '20px' }}>
          <Col xs={4} md={4} className={'center'} style={{ padding: '1% 1% 1% 1%' }}>
            <img src={data.icon_url} className='user_img' />
          </Col>
          <Col xs={7} md={7} className={'left center_vertical'} style={{ padding: '1% 1% 1% 1%' }}>
            <Row>
              <p style={{ padding: '1% 1% 1% 1%' }}>
                {'username: ' + data.username} <br />
                {'name: ' + data.name} <br />
                {'email: ' + data.email} <br />
              </p>
            </Row>
          </Col>
          <Col xs={1} md={1}>
            {/* delete button */}
            <div
              className='center center_vertical'
              style={{
                height: '100%',
                padding: '1% 1% 1% 1%',
              }}
            >
              <FontAwesomeIcon
                icon={faTrash}
                size='2x'
                onClick={onDelete}
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
      <Row style={{ padding: '0 0 0 0', fontSize: '24px' }}>
        <Col md={1} className={'center'} style={{ padding: '1% 1% 1% 1%' }}>
          <img src={data.icon_url} className='user_img' />
        </Col>
        <Col md={3} className={'left center_vertical'} style={{ padding: '1% 1% 1% 1%' }}>
          {data.username}
        </Col>
        <Col md={3} className={'left center_vertical'} style={{ padding: '1% 1% 1% 1%' }}>
          {data.name}
        </Col>
        <Col md={4} className={'left center_vertical'} style={{ padding: '1% 1% 1% 1%' }}>
          {data.email}
        </Col>
        <Col md={1} className={'center center_vertical'} style={{ padding: '1% 1% 1% 1%' }}>
          <FontAwesomeIcon icon={faTrash} onClick={onDelete} style={{ cursor: 'pointer' }} />
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
