import { useContext } from 'react';
import { AuthContext } from '@components/AuthProvider';

export const useAuth = () => {
  const { tokenRef } = useContext(AuthContext);
  return tokenRef;
};
