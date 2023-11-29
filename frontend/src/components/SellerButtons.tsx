import UserButton from './UserButton';

const SellerButtons = () => {
  return (
    <div>
      <UserButton url='/user/seller/info' text='Shop Info' />
      <UserButton url='/user/seller/manageProducts' text='All Products' />
      <UserButton url='/user/seller/manageCoupons' text='All Coupons' />
      <UserButton url='/user/seller/orders' text='All Shipments' />
      <UserButton url='/user/seller/reports' text='All Reports' />
    </div>
  );
};

export default SellerButtons;
