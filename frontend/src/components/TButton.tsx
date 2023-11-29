import '@components/style.css';
import '@style/global.css';

import { Button } from 'react-bootstrap';
import { Link } from 'react-router-dom';

interface Props {
  text: string;
  url: string;
}

const TButton = ({ text, url }: Props) => {
  const button = <Button className='button pointer center'>{text}</Button>;

  const urlButton = (
    <Link to={url} style={{ color: 'white', width: '100%' }}>
      <Button className='none button pointer center'>{text}</Button>
    </Link>
  );

  return (
    <div className='center' style={{ width: '100%' }}>
      {url ? urlButton : button}
    </div>
  );
};

export default TButton;
