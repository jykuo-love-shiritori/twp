import { MutableRefObject, ReactNode, createContext, useRef } from 'react';

export type Context = {
  tokenRef: MutableRefObject<string>;
};

export const AuthContext = createContext<Context>({} as Context);

interface Props {
  children?: ReactNode;
}

const AuthProvider: React.FC<Props> = ({ children }) => {
  const tokenRef = useRef('');

  return <AuthContext.Provider value={{ tokenRef: tokenRef }}>{children}</AuthContext.Provider>;
};

export default AuthProvider;
