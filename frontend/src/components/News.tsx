import 'bootstrap/dist/css/bootstrap.min.css';
import '@style/global.css';

import TButton from '@components/TButton';

interface Props {
  id: number;
  imgUrl: string;
  title: string;
}

const News = ({ id, imgUrl, title }: Props) => {
  const NewsComponentStyle = {
    borderRadius: '52px',
    boxShadow: '6px 6px 15px 5px rgba(0, 0, 0, 0.15)',
    marginBottom: '20px',
    width: '100%',
    border: 'var(--border) solid 1px',
    height: '250px',
  };

  return (
    <div>
      <img src={imgUrl} style={NewsComponentStyle} />

      <div style={{ padding: '1% 10% 1% 10%' }} className='center'>
        <span>
          {title.substring(0, 25)} {title.length > 25 ? '...' : ''}
        </span>
      </div>

      <TButton text='more' url={`/news/${id}`} />
      <br />
    </div>
  );
};

export default News;
