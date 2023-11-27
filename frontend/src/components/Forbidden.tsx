const Forbidden = () => {
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
        403 <br />
        Forbidden
        <hr className='white_bg' />
      </div>
      <div className='center'>Access to this resource on the server is denied</div>
    </div>
  );
};

export default Forbidden;
