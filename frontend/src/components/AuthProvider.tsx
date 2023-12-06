import { createContext, useState } from 'react';
import { Outlet } from 'react-router-dom';

export type Context = {
  token: string;
  setToken: (v: string) => void;
};

export const AuthContext = createContext<Context>({} as Context);

export const AuthProvider = () => {
  const [token, setToken] = useState('');

  return (
    <AuthContext.Provider value={{ token, setToken }}>
      <Outlet />
    </AuthContext.Provider>
  );
};
