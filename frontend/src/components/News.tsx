import 'bootstrap/dist/css/bootstrap.min.css';
import '@components/style.css';
import '@style/global.css';

import TButton from '@components/TButton';

interface Props {
  id: number;
  imgUrl: string;
  title: string;
}

const News = ({ id, imgUrl, title }: Props) => {
  return (
    <div>
      <img src={imgUrl} className='news_pic_c' />

      <div style={{ padding: '1% 10% 1% 10%' }} className='center'>
        <span>
          {title.substring(0, 25)} {title.length > 25 ? '...' : ''}
        </span>
      </div>

      <TButton text='more' action={`/news/${id}`} />
      <br />
    </div>
  );
};

export default News;
