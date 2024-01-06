/* eslint-disable react-hooks/exhaustive-deps */
import '@components/style.css';
import '@style/global.css';

import { Row, Col, Button } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faSearch } from '@fortawesome/free-solid-svg-icons';
import { useEffect, useState } from 'react';
import { useNavigate, useSearchParams } from 'react-router-dom';

import { SearchProps } from '@pages/search';

// eslint-disable-next-line react-refresh/only-export-components
export const getUrl = (data: SearchProps, q: string) => {
  let url: string = '/search?q=' + q;
  url += getBehind(data);
  return url;
};

// eslint-disable-next-line react-refresh/only-export-components
export const getShopSearchUrl = (data: SearchProps, q: string) => {
  // TODO : the sellerID should be the real seller name
  let url: string = '/sellerID/shop/inside/search?q=' + q;
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
  const [searchParams, setSearchParams] = useSearchParams();
  const navigate = useNavigate();
  const [q, setQ] = useState<string>(searchParams.get('q') ?? '');

  useEffect(() => {
    if (window.location.href.includes('search')) {
      transfer();
    }
  }, [q, searchParams]);

  const getNewParams = () => {
    const minPrice = searchParams.get('minPrice');
    const maxPrice = searchParams.get('maxPrice');
    const minStock = searchParams.get('minStock');
    const maxStock = searchParams.get('maxStock');
    const haveCouponValue = searchParams.get('haveCoupon');
    const sortBy = searchParams.get('sortBy');
    const order = searchParams.get('order');

    if (minPrice) searchParams.set('minPrice', minPrice);
    if (maxPrice) searchParams.set('maxPrice', maxPrice);
    if (minStock) searchParams.set('minStock', minStock);
    if (maxStock) searchParams.set('maxStock', maxStock);
    if (haveCouponValue === 'true' || haveCouponValue === 'false') {
      searchParams.set('haveCoupon', haveCouponValue);
    }
    if (sortBy && ['price', 'stock', 'sales', 'relevancy'].includes(sortBy)) {
      searchParams.set('sortBy', sortBy);
    }
    if (order && ['asc', 'desc'].includes(order)) {
      searchParams.set('order', order);
    }
    if (q) searchParams.set('q', q);

    setSearchParams(searchParams);

    return searchParams.toString();
  };

  const transfer = () => {
    const isShopURL = /\/shop(?:\/inside\/search(?:\?.+)?)?$/;

    if (isShopURL.test(window.location.pathname)) {
      // TODO : the sellerID need to be changed to real seller
      navigate('/sellerID/shop/inside/search?' + getNewParams());
    } else {
      navigate('search?' + getNewParams());
    }
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
