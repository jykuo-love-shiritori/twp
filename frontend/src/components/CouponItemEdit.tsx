import 'bootstrap/dist/css/bootstrap.min.css';
import { Link } from 'react-router-dom';
import CouponItemTemplate from '@components/CouponItemTemplate';

interface ICouponItemEdit {
  id: number;
  scope: 'global' | 'shop';
  name: string;
  type: 'percentage' | 'fixed' | 'shipping';
  discount: number;
  expire_date: string;
}

const CouponItemEdit = ({ data }: { data: ICouponItemEdit }) => {
  return (
    <Link className='none' to={`${window.location.pathname}/${data.id}`}>
      <div style={{ cursor: 'pointer' }}>
        <CouponItemTemplate data={data} />
      </div>
    </Link>
  );
};

export default CouponItemEdit;
