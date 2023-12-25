import '@style/global.css';
import '@components/style.css';

import { Outlet, useLocation } from 'react-router-dom';
import { useEffect } from 'react';

import Footer from '@components/Footer';
import NavBar from '@components/NavBar';

const Layout = () => {
  const { pathname } = useLocation();

  useEffect(() => {
    window.scrollTo({ top: 0, left: 0, behavior: 'instant' });
  }, [pathname]);

  return (
    <div>
      <div style={{ position: 'relative', zIndex: '10' }}>
        <NavBar />
      </div>
      <div className='body_down bg flex_wrapper' style={{ position: 'relative', zIndex: '1' }}>
        <Outlet />
      </div>
      <div style={{ position: 'relative', zIndex: '10' }}>
        <Footer />
      </div>
    </div>
  );
};

export default Layout;
