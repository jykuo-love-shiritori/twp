import UserButton from './UserButton';

const AdminButtons = () => {
  return (
    <div>
      <UserButton url='/admin' text='Manage Users' />
      <UserButton url='/admin/manageCoupons' text='Global Coupons' />
      <UserButton url='/admin/report' text='Site Reports' />
    </div>
  );
};

export default AdminButtons;
