const Unauthorized = () => {
  return (
    <div style={{ padding: '20% 10% 10% 10%' }}>
      <div className='wrong'>
        401 Unauthorized
        <hr className='white_bg' />
      </div>
      <div className='center'>Access is denied due to invalid credentials</div>
    </div>
  );
};

export default Unauthorized;
