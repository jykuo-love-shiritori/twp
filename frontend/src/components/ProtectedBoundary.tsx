import { useQuery } from '@tanstack/react-query';
import { AuthContext } from './AuthProvider';
import { useContext } from 'react';
import { Navigate, Outlet } from 'react-router-dom';
import { TryRefresh } from '@lib/Auth';

const ProtectedBoundary = () => {
  const { tokenRef } = useContext(AuthContext);

  const { isLoading, isError, data } = useQuery({
    queryKey: ['refresh'],
    queryFn: TryRefresh,
    enabled: !tokenRef.current || tokenRef.current.length === 0,
    retry: false,
  });

  if (isLoading) {
    return <div>Loading</div>;
  }

  if (isError) {
    return <Navigate to='/login' />;
  }

  if (data) {
    tokenRef.current = data.access_token;
  }

  return <Outlet />;
};

export default ProtectedBoundary;
