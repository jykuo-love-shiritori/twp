import { Routes, Route, BrowserRouter } from 'react-router-dom';

import Layout from '@pages/Layout';
import Home from '@pages/home';
import EachNews from '@pages/news/[newsID]';
import Discover from '@pages/discover';
import EachGoods from '@pages/discover/[goodsID]';
import Cart from '@pages/cart';
import User from '@pages/user/buyer/index';
import Login from '@pages/user/login';
import Signup from '@pages/user/signup';
import HistoryEach from '@pages/user/buyer/history/[historyID]';
import Info from '@components/Info';
import History from '@components/History';
import NotFound from '@components/NotFound';
import SearchNotFound from '@components/SearchNotFound';
import APItest from '@components/APItest';
import Seller from '@pages/user/seller';
import Forbidden from '@components/Forbidden';
import Unauthorized from '@components/Unauthorized';
import Security from '@pages/user/buyer/security';
import Password from '@pages/user/buyer/security/Password';
import CreditCard from '@pages/user/buyer/security/CreditCard';
import NewCard from '@pages/user/buyer/security/NewCard';
import Shop from '@pages/user/buyer/Shop';
import Products from '@pages/user/seller/Products';
import NewGoods from '@pages/user/seller/NewGoods';
import Authorize from '@pages/user/authorize';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path='/login' element={<Login />} />
        <Route path='/authorize' element={<Authorize />} />
        <Route path='/signup' element={<Signup />} />

        <Route path='/' element={<Layout />}>
          <Route index element={<Home />} />
          <Route path='/' element={<Home />} />
          <Route path='/news'>
            <Route path=':news_id' element={<EachNews />} />
          </Route>
          <Route path='/discover' element={<Discover />} />
          <Route path='/discover'>
            <Route path=':goods_id' element={<EachGoods />} />
          </Route>
          <Route path='/user' element={<User />}>
            <Route index element={<Info />} />
            <Route path='/user/info' element={<Info />} />
            <Route path='/user/security' element={<Security />} />
            <Route path='/user/security/password' element={<Password />} />
            <Route path='/user/security/manageCreditCard' element={<CreditCard />} />
            <Route path='/user/security/manageCreditCard/newCard' element={<NewCard />} />
            <Route path='/user/buyer/order' element={<History />} />
          </Route>
          <Route path='/user/buyer/order'>
            <Route path=':history_id' element={<HistoryEach />} />
          </Route>

          <Route path='/user/seller' element={<Seller />}>
            <Route path='/user/seller/info' element={<NotFound />} />
            <Route path='/user/seller/manageProducts' element={<Products />} />
            <Route path='/user/seller/manageCoupons' element={<NotFound />} />
            <Route path='/user/seller/orders' element={<NotFound />} />
            <Route path='/user/seller/reports' element={<NotFound />} />
          </Route>

          <Route path='/user/seller/manageProducts/new' element={<NewGoods />} />

          <Route path='/sellerID/shop' element={<Shop />} />

          <Route path='/buyer/cart' element={<Cart />} />

          <Route path='/search?' element={<SearchNotFound />} />

          <Route path='*' element={<NotFound />} />
          <Route path='/forbidden' element={<Forbidden />} />
          <Route path='/unauthorized' element={<Unauthorized />} />

          <Route path='/APItest' element={<APItest />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
