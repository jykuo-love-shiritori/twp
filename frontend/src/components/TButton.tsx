import '@components/style.css';
import '@style/global.css';

import { Link } from 'react-router-dom';

interface Props {
  text: string;
  url: string;
}

const TButton = ({ text, url }: Props) => {
  const button = <div className='button pointer center'>{text}</div>;

  const urlButton = (
    <Link to={url} className='none button pointer center' style={{ color: 'white' }}>
      <div className='center'>{text}</div>
    </Link>
  );

  return (
    <div className='center' style={{ width: '100%' }}>
      {url ? urlButton : button}
    </div>
  );
};

export default TButton;
