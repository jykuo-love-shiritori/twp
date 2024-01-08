import UserButton from '@components/UserButton';
import { IsAdmin } from '@lib/Auth';

const BuyerButtons = () => {
  const isAdmin = IsAdmin();
  return (
    <div>
      <UserButton url='/user/info' text='Personal info' />
      {!isAdmin ? (
        <div>
          <UserButton url='/user/security' text='Security' />
          <UserButton url='/user/buyer/order' text='Order history' />
        </div>
      ) : null}
    </div>
  );
};

export default BuyerButtons;
