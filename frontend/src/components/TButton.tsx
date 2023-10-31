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
    <div style={{ width: '100%' }}>
      <div className='button pointer'>
        <Link to={url} className='none' style={{ color: 'white' }}>
          <div className='center'>{text}</div>
        </Link>
      </div>
    </div>
  );

  return <div className='center'>{url ? urlButton : button}</div>;
};

export default TButton;
