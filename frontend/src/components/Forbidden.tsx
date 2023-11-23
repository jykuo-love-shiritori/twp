const Forbidden = () => {
  return (
    <div style={{ padding: '20% 10% 10% 10%' }}>
      <div className='wrong'>
        403 Forbidden
        <hr className='white_bg' />
      </div>
      <div className='center'>Access to this resource on the server is denied</div>
    </div>
  );
};

export default Forbidden;
