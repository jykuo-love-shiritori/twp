import 'bootstrap/dist/css/bootstrap.min.css'

import '../components/style.css';
import '../style/global.css';
import PButton from './PButton';

interface Props {
    id: number;
    imgUrl: string;
    title: string;
}

const News = ({ id, imgUrl, title }: Props) => {
    return (
        <div>
            <img src={imgUrl} className='new_pic_c' />

            <div style={{ padding: '1% 15% 1% 15%' }} className='center'>
                <span className='news_title_c '>{title}</span>
            </div>

            <PButton text="more" url={`/news/${id}`} />
            <br />
        </div>
    )
}

export default News