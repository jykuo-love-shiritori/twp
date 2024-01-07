import { createContext, ReactNode, useState } from 'react';

export type Context = {
  q: string;
  setQ: React.Dispatch<React.SetStateAction<string>>;
};

export const SearchContext = createContext<Context>({} as Context);

interface Props {
  children?: ReactNode;
}

const SearchProvider: React.FC<Props> = ({ children }) => {
  const [q, setQ] = useState('');
  return <SearchContext.Provider value={{ q: q, setQ: setQ }}>{children}</SearchContext.Provider>;
};

export default SearchProvider;
