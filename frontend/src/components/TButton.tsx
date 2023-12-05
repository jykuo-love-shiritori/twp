import '@components/style.css';
import '@style/global.css';
import { useNavigate } from 'react-router-dom';
import { Button } from 'react-bootstrap';

interface Props {
  text: string;
  url?: string;
  onClick?: () => void;
}

const TButton = ({ text, url, onClick }: Props) => {
  const navigate = useNavigate();

  function handleClick() {
    if (url) {
      navigate(url);
    } else if (onClick) {
      onClick();
    }
  }

  return (
    <Button className='none button pointer center' onClick={handleClick}>
      {text}
    </Button>
  );
};

export default TButton;
