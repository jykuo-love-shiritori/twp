import { Col, Row } from 'react-bootstrap';
import { Dispatch, SetStateAction } from 'react';

interface Props {
  text: string;
  value: string;
  setValue: Dispatch<SetStateAction<string>> | ((value: string) => void);
}

const PasswordItem = ({ text, value, setValue }: Props) => {
  return (
    <Row style={{ margin: '2% 0% 2% 0% ' }}>
      <Col xs={12} md={4} className='center_vertical'>
        {text}
      </Col>
      <Col xs={12} md={8}>
        <input
          type='password'
          placeholder={text}
          className='input_box'
          value={value ? value : ''}
          onChange={setValue ? (e) => setValue(e.target.value) : undefined}
        />
      </Col>
    </Row>
  );
};

export default PasswordItem;
