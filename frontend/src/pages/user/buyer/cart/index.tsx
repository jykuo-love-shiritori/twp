import CartGroup from '@components/Cart';
import { useQuery } from '@tanstack/react-query';

interface CartProps {
  cartInfo: { id: number; image_id: string; seller_name: string; shop_name: string };
  coupons: CouponProps[];
  products: ProductProps[];
}

interface CouponProps {
  description: string;
  discount: number;
  id: number;
  name: string;
  type: string; // 'percentage' | 'fixed' | 'shipping'
  scope: string; // 'global' | 'shop'
}

interface ProductProps {
  enabled: boolean;
  image_id: string;
  name: string;
  price: number;
  product_id: number;
  quantity: number;
  stock: number;
}

const BuyerCarts = () => {
  const { data, isLoading, isError, refetch } = useQuery({
    queryKey: ['jsonData'],
    queryFn: async () => {
      const response = await fetch('/resources/Carts.json');
      if (!response.ok) {
        throw new Error('Failed to fetch data');
      }
      return response.json();
    },
  });
  const onRefetch = () => {
    console.log('refetch');
    refetch();
  };
  if (isLoading) {
    return <div>Loading...</div>;
  }
  if (isError) {
    return <div>Error fetching data</div>;
  }
  const fetchedData = data as CartProps[];

  return (
    <div style={{ padding: '10% 5% 10% 5%' }}>
      <span className='title'>Cart</span>

      {fetchedData.map((cartData, index) => (
        <CartGroup data={cartData} key={index} onRefetch={onRefetch} />
      ))}
    </div>
  );
};

export default BuyerCarts;
