/* eslint-disable react-hooks/exhaustive-deps */
import '@components/style.css';
import '@style/global.css';

import { Row, Col, Button } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faSearch } from '@fortawesome/free-solid-svg-icons';
import { useEffect, useState } from 'react';
import { useNavigate, useSearchParams } from 'react-router-dom';

import { SearchProps, defaultFilterValues } from '@pages/search';

// eslint-disable-next-line react-refresh/only-export-components
export const getUrl = (data: SearchProps, q: string) => {
  let url: string = '/search?q=' + q;
  if (data.minPrice !== null) {
    url += '&minPrice=' + data.minPrice;
  }
  if (data.maxPrice !== null) {
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
  const [searchParams] = useSearchParams();
  const navigate = useNavigate();
  const [q, setQ] = useState<string>(searchParams.get('q') ?? '');

  useEffect(() => {
    if (window.location.href.includes('search')) {
      transfer();
    }
  }, [q, navigate]);

  const transfer = () => {
    const data: SearchProps = { ...defaultFilterValues, q: q };

    const maxPrice = parseInt(searchParams.get('maxPrice') || '');
    if (!isNaN(maxPrice)) {
      data.maxPrice = maxPrice;
    }

    const minPrice = parseInt(searchParams.get('minPrice') || '');
    if (!isNaN(minPrice)) {
      data.minPrice = minPrice;
    }

    const maxStock = parseInt(searchParams.get('maxStock') || '');
    if (!isNaN(maxStock)) {
      data.maxStock = maxStock;
    }

    const minStock = parseInt(searchParams.get('minStock') || '');
    if (!isNaN(minStock)) {
      data.minStock = minStock;
    }

    const haveCouponValue = searchParams.get('haveCoupon');
    if (haveCouponValue !== null && haveCouponValue !== '') {
      if (haveCouponValue === 'true' || haveCouponValue === 'false') {
        data.haveCoupon = JSON.parse(haveCouponValue);
      }
    }

    const sortBy = searchParams.get('sortBy');
    if (
      sortBy !== null &&
      sortBy !== '' &&
      (sortBy === 'price' || sortBy === 'stock' || sortBy === 'sales' || sortBy === 'relevancy')
    ) {
      data.sortBy = sortBy;
    }

    const order = searchParams.get('order');
    if (order !== null && order !== '' && (order === 'asc' || order === 'desc')) {
      data.order = order;
    }

    navigate(getUrl(data, q));
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
              // onChange={(e) => queryOnChange(e.target.value)}
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
