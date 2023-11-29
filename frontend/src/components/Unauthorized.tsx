const Unauthorized = () => {
  const WarpStyle = {
    display: 'flex',
    minHeight: '100vh',
    flexDirection: 'column',
    justifyContent: 'flex-start',
    padding: '20% 10% 10% 10%',
  } as const;

  return (
    <div style={WarpStyle}>
      <div className='wrong'>
        401 <br />
        Unauthorized
        <hr className='white_bg' />
      </div>
      <div className='center'>Access is denied due to invalid credentials</div>
    </div>
  );
};

export default Unauthorized;
