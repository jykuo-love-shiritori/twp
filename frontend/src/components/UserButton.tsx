import { Link } from 'react-router-dom';

interface Props {
  url: string;
  text: string;
}

const UserButton = ({ url, text }: Props) => {
  return (
    <Link to={url} className='none'>
      <div className='user_button'>
        <span className='white_word'>{text}</span>
      </div>
    </Link>
  );
};

export default UserButton;
