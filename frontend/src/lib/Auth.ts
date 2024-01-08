import { useContext } from 'react';
import { AuthContext } from '@components/AuthProvider';
import { jwtDecode, JwtPayload } from 'jwt-decode';

interface CustomJwtPayload extends JwtPayload {
  role: string;
  username: string;
}

export type Token = {
  access_token: string;
};

export const TryRefresh = async () => {
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

export const IsAdmin = () => {
  const { tokenRef } = useContext(AuthContext);
  try {
    const decoded = jwtDecode<CustomJwtPayload>(tokenRef.current);
    return decoded.role === 'admin';
  } catch (error) {
    console.error('Invalid token', error);
  }
};

export const GetUserName = () => {
  const { tokenRef } = useContext(AuthContext);
  try {
    const decoded = jwtDecode<CustomJwtPayload>(tokenRef.current);
    return decoded.username;
  } catch (error) {
    console.error('Invalid token', error);
  }
};
