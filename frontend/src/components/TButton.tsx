import '../components/style.css';
import '../style/global.css';

interface Props {
  text: string;
  url: string;
}

const TButton = ({ text, url }: Props) => {
  let button = <div className='button pointer center'>{text}</div>;

  let urlButton = (
    <div style={{ width: '100%' }}>
      <div className='button pointer'>
        <a href={url} className='none' style={{ color: 'white' }}>
          <div className='center'>{text}</div>
        </a>
      </div>
    </div>
  );

  return <div className='center'>{url ? urlButton : button}</div>;
};

export default TButton;
