import { Routes, Route, BrowserRouter } from 'react-router-dom';

import Layout from '@pages/Layout';
import Home from '@pages/home';
import EachNews from '@pages/news/[newsID]';
import Discover from '@pages/discover';
import EachGoods from '@pages/discover/[goodsID]';
import Cart from '@pages/cart';
import User from '@pages/user';
import Login from '@pages/user/login';
import Signup from '@pages/user/signup';
import HistoryEach from '@pages/user/history/[historyID]';
import Info from '@components/Info';
import History from '@components/History';
import NotFound from '@components/NotFound';
import SearchNotFound from '@components/SearchNotFound';
import APItest from '@components/APItest';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path='/login' element={<Login />} />
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
          <Route path='user' element={<User />}>
            <Route index element={<Info />} />
            <Route path='/user/info' element={<Info />} />
            <Route path='/user/buyer/order' element={<History />} />
          </Route>
          <Route path='user/buyer/order'>
            <Route path=':history_id' element={<HistoryEach />} />
          </Route>
          <Route path='buyer/cart' element={<Cart />} />

          <Route path='/search?' element={<SearchNotFound />} />

          <Route path='*' element={<NotFound />} />

          <Route path='/APItest' element={<APItest />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
