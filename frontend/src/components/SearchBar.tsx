import '@components/style.css';
import '@style/global.css';

import { Row, Col, Button } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faSearch } from '@fortawesome/free-solid-svg-icons';
import { FormEventHandler, useEffect, useState } from 'react';

import { useLocation, useNavigate, useParams, useSearchParams } from 'react-router-dom';

import { SearchProps } from '@pages/search';

// eslint-disable-next-line react-refresh/only-export-components
export const setFilter = (data: SearchProps, params: URLSearchParams) => {
  if (data.minPrice) params.set('minPrice', data.minPrice.toString());
  if (data.maxPrice) params.set('maxPrice', data.maxPrice.toString());
  if (data.minStock) params.set('minStock', data.minStock.toString());
  if (data.maxStock) params.set('maxStock', data.maxStock.toString());
  if (data.haveCoupon) params.set('haveCoupon', data.haveCoupon.toString());
  if (data.sortBy) params.set('sortBy', data.sortBy);
  if (data.order) params.set('order', data.order);
  return params;
};

const SearchBar = () => {
  const { sellerName: initialSellerName } = useParams();
  const location = useLocation();
  const sellerNameFromParams = useParams().sellerName;
  const [sellerName, setSellerName] = useState<string | undefined>(initialSellerName);
  const [searchParams] = useSearchParams();

  useEffect(() => {
    setSellerName(sellerNameFromParams);
  }, [location, sellerNameFromParams]);

  const navigate = useNavigate();
  const [q, setQ] = useState<string>(searchParams.get('q') ?? '');

  const transfer: FormEventHandler<HTMLFormElement> = (e) => {
    e.preventDefault();

    if (!q) {
      return;
    }

    const params = new URLSearchParams({ q: q });

    console.log(params.toString());

    const isShopURL = /\/shop\/(?:[^/]+\/)?products(?:\/inside\/search(?:\?.+)?)?$/;
    if (isShopURL.test(window.location.pathname) && sellerName) {
      navigate(`/shop/${sellerName}/products/inside/search?${params.toString()}`);
      return;
    }

    const url = `/search?${params.toString()}`;
    navigate(url);
  };

  return (
    <div>
      <form onSubmit={transfer}>
        <Row>
          <Col xs={12} sm={12} md={12} lg={8} className='center'>
            <div className='input_container'>
              <input
                type='text'
                placeholder='Search'
                value={q}
                onChange={(e) => setQ(e.target.value)}
                className='search'
              />
              <FontAwesomeIcon icon={faSearch} className='search_icon' />
            </div>
          </Col>
          <Col xs={4} lg={4} className='disappear_tablet disappear_phone'>
            <Button className='search_button center' type='submit'>
              Search
            </Button>
          </Col>
        </Row>
      </form>
    </div>
  );
};

export default SearchBar;
