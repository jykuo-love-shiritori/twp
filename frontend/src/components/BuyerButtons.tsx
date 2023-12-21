import UserButton from '@components/UserButton';

const BuyerButtons = () => {
  return (
    <div>
      <UserButton url='/user/info' text='Personal info' />
      <UserButton url='/user/security' text='Security' />
      <UserButton url='/user/buyer/order' text='Order history' />
    </div>
  );
};

export default BuyerButtons;
