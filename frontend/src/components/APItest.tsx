import '@components/style.css';
import '@style/global.css';
import { useQuery } from '@tanstack/react-query';

const APItest = () => {
  const { isLoading, isError, data } = useQuery({
    queryKey: ['APItest'],
    queryFn: async () => {
      const response = await fetch('/api/ping');
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      return response.json();
    },
  });

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (isError) {
    return <div>Error</div>;
  }

  return (
    <div className='center wrong' style={{ padding: '20% 10% 10% 10%' }}>
      {data.message}
    </div>
  );
};

export default APItest;
