import '../components/style.css';
import '../style/global.css';

interface Props {
    text: string;
    url: string;
}

const PButton = ({ text, url }: Props) => {
    let button = (
        <div className='news_button pointer button2_bg'>
            {text}
        </div>
    );

    let urlButton = (
        <div className='news_button pointer button2_bg'>
            <a href={url} className='none' style={{ color: 'white' }}>{text}</a>
        </div>
    );

    return (
        <div className='center'>
            {url ? urlButton : button}
        </div >
    )
}

export default PButton
