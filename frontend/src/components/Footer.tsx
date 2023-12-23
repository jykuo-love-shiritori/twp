import '@style/global.css';
import LogoImgUrl from '@assets/images/logo.png';

const Footer = () => {
  const FooterStyle = {
    backgroundColor: 'var(--layout)',
    bottom: '0',
    width: '100%',
    height: '60px',
    padding: '1% 0% 1% 0%',
    fontSize: '12px',
    fontFamily: 'var(--font)',
  };

  return (
    <div className='center' style={FooterStyle}>
      <img src={LogoImgUrl} style={{ height: '100%' }}></img>
      <span className='white_word'>&nbsp;&nbsp;Copyright Ⓒ 2023 All right reserved</span>
    </div>
  );
};

export default Footer;
