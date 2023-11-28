import '@components/style.css';
import '@style/global.css';

const NotFound = () => {
  const WarpStyle = {
    display: 'flex',
    minHeight: '100vh',
    flexDirection: 'column',
    justifyContent: 'flex-start',
    padding: '20% 10% 10% 10%',
  } as const;

  return (
    <div className='center wrong' style={WarpStyle}>
      404 <br />
      Page Not Found
    </div>
  );
};

export default NotFound;
