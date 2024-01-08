import '@components/style.css';
import '@style/global.css';

import { Row, Col, Button } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faSearch } from '@fortawesome/free-solid-svg-icons';
import { useEffect, useState } from 'react';

import { useLocation, useNavigate, useParams } from 'react-router-dom';

import { SearchProps } from '@pages/search';

// eslint-disable-next-line react-refresh/only-export-components
export const getUrl = (data: SearchProps, q: string) => {
  let url: string = '/search?q=' + q;
  url += getBehind(data);
  return url;
};

// eslint-disable-next-line react-refresh/only-export-components
export const getBehind = (data: SearchProps) => {
  let url = '';

  if (data.minPrice !== null && data.minPrice.toString() !== '') {
    url += '&minPrice=' + data.minPrice;
  }
  if (data.maxPrice !== null && data.maxPrice.toString() !== '') {
    url += '&maxPrice=' + data.maxPrice;
  }
  if (data.minStock !== null) {
    url += '&minStock=' + data.minStock;
  }
  if (data.maxStock !== null) {
    url += '&maxStock=' + data.maxStock;
  }
  if (data.haveCoupon !== null) {
    url += '&haveCoupon=' + data.haveCoupon.toString();
  }
  if (data.sortBy !== null) {
    url += '&sortBy=' + data.sortBy;
  }
  if (data.order !== null) {
    url += '&order=' + data.order;
  }
  url = url.replace(/ /g, '%20');
  return url;
};

const SearchBar = () => {
  const { sellerName: initialSellerName } = useParams();
  const location = useLocation();
  const sellerNameFromParams = useParams().sellerName;
  const [sellerName, setSellerName] = useState<string | undefined>(initialSellerName);

  useEffect(() => {
    setSellerName(sellerNameFromParams);
  }, [location, sellerNameFromParams]);

  const navigate = useNavigate();
  const [q, setQ] = useState<string>('');

  const transfer = () => {
    if (q === '') {
      return;
    }

    let fullUrl = '';
    const searchQuery = q ? `?q=${encodeURIComponent(q)}` : '';

    const isShopURL = /\/shop\/(?:[^/]+\/)?products(?:\/inside\/search(?:\?.+)?)?$/;
    if (isShopURL.test(window.location.pathname)) {
      if (sellerName !== undefined) {
        fullUrl = `${`/shop/${sellerName}/products/inside/search`}${searchQuery}`;
      }
    } else {
      fullUrl = `${`/search`}${searchQuery}`;
    }
    navigate(fullUrl);
  };

  const handleKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.keyCode === 229) return;

    if (event.key === 'Enter') {
      transfer();
    }
  };

  const handleClick = () => {
    transfer();
  };

  return (
    <div>
      <Row>
        <Col xs={12} sm={12} md={12} lg={8} className='center'>
          <div className='input_container'>
            <input
              type='text'
              placeholder='Search'
              value={q}
              onChange={(e) => setQ(e.target.value)}
              onKeyDown={handleKeyDown}
              className='search'
            />
            <FontAwesomeIcon icon={faSearch} className='search_icon' />
          </div>
        </Col>
        <Col xs={4} lg={4} className='disappear_tablet disappear_phone'>
          <Button className='search_button center' onClick={handleClick}>
            Search
          </Button>
        </Col>
      </Row>
    </div>
  );
};

export default SearchBar;
