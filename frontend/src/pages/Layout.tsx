import '../style/global.css';
import '../components/style.css';

import { Outlet } from 'react-router-dom';

import Footer from '../components/Footer';
import NavBar from '../components/NavBar';

const Layout = () => {
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
