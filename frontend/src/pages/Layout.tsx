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
      <NavBar />
      <div className='body_down bg flex-wrapper'>
        <Outlet />
      </div>
      <Footer />
    </div>
  );
};

export default Layout;
