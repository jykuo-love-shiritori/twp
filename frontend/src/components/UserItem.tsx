import { Col, Row } from 'react-bootstrap';

interface Props {
  img_path: string;
  name: string;
}

const UserItem = ({ img_path, name }: Props) => {
  return (
    <Row>
      <Col xs={2} className='center'>
        <img src={img_path} className='user' />
      </Col>
      <Col xs={10} className='center_vertical'>
        {name}
      </Col>
    </Row>
  );
};

export default UserItem;
