import { useContext } from 'react';
import { AuthContext } from '@components/AuthProvider';

export type Token = {
  access_token: string;
};

export const TryRefresh = async () => {
  console.log('refresh');
  const refreshUrl = '/api/oauth/refresh';
  const resp = await fetch(refreshUrl, { method: 'POST' });
  if (!resp.ok) {
    throw new Error('failed refreshing token');
  }
  return (await resp.json()) as Token;
};

export const useAuth = () => {
  const { tokenRef } = useContext(AuthContext);
  return tokenRef.current;
};
