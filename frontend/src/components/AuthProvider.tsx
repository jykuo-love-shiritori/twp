import { useQuery } from '@tanstack/react-query';
import { MutableRefObject, createContext, useRef } from 'react';
import { Navigate, Outlet } from 'react-router-dom';

export type Token = {
  access_token: string;
};

export type Context = {
  tokenRef: MutableRefObject<string>;
};

export const AuthContext = createContext<Context>({} as Context);

const refresh = async () => {
  console.log('refresh');
  const refreshUrl = '/api/oauth/refresh';
  const resp = await fetch(refreshUrl, { method: 'POST' });
  return (await resp.json()) as Token;
};

const AuthProvider = () => {
  const tokenRef = useRef('');

  const { isLoading, isError, data } = useQuery({
    queryKey: ['refresh'],
    queryFn: refresh,
    enabled: tokenRef.current.length === 0,
  });

  if (isLoading) {
    return <div>Loading</div>;
  }

  if (isError) {
    return <Navigate to='/login'></Navigate>;
  }

  if (data) {
    tokenRef.current = data.access_token;
  }

  return (
    <AuthContext.Provider value={{ tokenRef: tokenRef }}>
      <Outlet />
    </AuthContext.Provider>
  );
};

export default AuthProvider;
