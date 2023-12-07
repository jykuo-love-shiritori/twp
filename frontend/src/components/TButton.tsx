import '@components/style.css';
import '@style/global.css';
import { useNavigate } from 'react-router-dom';
import { Button } from 'react-bootstrap';

interface Props {
  text: string;
  action?: string | (() => void);
}

const TButton = ({ text, action }: Props) => {
  const navigate = useNavigate();

  function handleClick() {
    if (typeof action === 'string') {
      navigate(action);
    }
    if (typeof action === 'function') {
      action();
    }
  }

  return (
    <div className='center none' style={{ width: '100%' }}>
      <Button className='none button pointer center' onClick={handleClick}>
        {text}
      </Button>
    </div>
  );
};

export default TButton;
