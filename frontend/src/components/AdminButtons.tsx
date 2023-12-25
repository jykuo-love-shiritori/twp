import UserButton from '@components/UserButton';

const AdminButtons = () => {
  return (
    <div>
      <UserButton url='/admin/manageUser' text='Manage Users' />
      <UserButton url='/admin/manageCoupons' text='Global Coupons' />
      <UserButton url='/admin/reports' text='Site Reports' />
    </div>
  );
};

export default AdminButtons;
