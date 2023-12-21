import '@components/style.css';
import '@style/global.css';
import Logoimage_url from '@assets/images/logo.png';

const Footer = () => {
  return (
    <div className='footer center'>
      <img src={Logoimage_url} style={{ height: '100%' }}></img>
      <span className='white_word'>&nbsp;&nbsp;Copyright Ⓒ 2023 All right reserved</span>
    </div>
  );
};

export default Footer;
