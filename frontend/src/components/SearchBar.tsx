import '@components/style.css';
import '@style/global.css';

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faSearch } from '@fortawesome/free-solid-svg-icons';
import { useContext } from 'react';
import { SearchContext } from './SearchProvider';

const SearchBar = () => {
  const { setQ } = useContext(SearchContext);
  const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const q = e.target.value;
    setQ(q);
  };

  return (
    <div className='input_container'>
      <input type='text' placeholder='Search' className='search' onChange={onChange} />
      <FontAwesomeIcon icon={faSearch} className='search_icon' />
    </div>
  );
};

export default SearchBar;
