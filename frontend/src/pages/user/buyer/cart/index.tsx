import Cart from '@components/Cart';
import { useAuth } from '@lib/Auth';
import { CheckFetchStatus, RouteOnNotOK } from '@lib/Status';
import { useQuery } from '@tanstack/react-query';
import { useNavigate } from 'react-router-dom';

interface ICoupon {
  description: string;
  discount: number;
  expire_date: string;
  id: number;
  name: string;
  scope: 'global' | 'shop';
  type: 'percentage' | 'fixed' | 'shipping';
}

interface IProduct {
  enabled: true;
  image_url: string;
  name: string;
  price: number;
  product_id: number;
  quantity: number;
  stock: number;
}

interface ICart {
  CartInfo: {
    id: number;
    seller_name: string;
    shop_image_url: string;
    shop_name: string;
  };
  Coupons: ICoupon[];
  Products: IProduct[];
}

const BuyerCarts = () => {
  const navigate = useNavigate();
  const token = useAuth();
  const { data, status, refetch } = useQuery({
    queryKey: ['buyerGetCart'],
    queryFn: async () => {
      const resp = await fetch('/api/buyer/cart', {
        method: 'GET',
        headers: {
          Authorization: `Bearer ${token}`,
          Accept: 'application/json',
        },
      });
      if (!resp.ok) {
        RouteOnNotOK(resp, navigate);
        return [];
      }
      return (await resp.json()) as ICart[];
    },
    enabled: true,
    refetchOnWindowFocus: false,
  });

  if (status !== 'success') {
    return <CheckFetchStatus status={status} />;
  }
  console.log(data);

  return (
    <div style={{ padding: '10% 5% 10% 5%' }}>
      <span className='title'>Cart</span>

      {data.length > 0 ? (
        data.map((cart, index) => (
          <Cart products={cart.Products} cartInfo={cart.CartInfo} key={index} refresh={refetch} />
        ))
      ) : (
        <h3>No unpaid cart 😎</h3>
      )}
    </div>
  );
};

export default BuyerCarts;
