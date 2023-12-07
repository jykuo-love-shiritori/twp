import '@components/style.css';
import '@style/global.css';

const NotFound = () => {
  const WarpStyle = {
    display: 'flex',
    minHeight: '100vh',
    FlexDirection: 'column',
    justifyContent: 'flex-start',
    padding: '20% 10% 10% 10%',
  };

  const WrongStyle = {
    color: 'var(--white, #fff)',
    TextAlign: 'center',
    fontFamily: 'Noto Sans TC',
    fontSize: '52px',
    fontStyle: 'normal',
    fontWeight: '900',
    lineHeight: 'normal',
    WordBreak: 'break-all',
    width: '100%',
  };

  return (
    <div style={WarpStyle}>
      <div style={WrongStyle}>
        <div className='center'>404</div>
        <div className='center'>Page Not Found</div>
      </div>
    </div>
  );
};

export default NotFound;
