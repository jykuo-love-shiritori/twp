const Unauthorized = () => {
  const WarpStyle = {
    display: 'flex',
    minHeight: '100vh',
    flexDirection: 'column',
    justifyContent: 'flex-start',
    padding: '20% 10% 10% 10%',
  } as const;

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
        <div className='center'>401 </div>
        <div className='center'>Unauthorized</div>

        <hr className='white_bg' />
      </div>
      <div className='center'>Access is denied due to invalid credentials</div>
    </div>
  );
};

export default Unauthorized;
