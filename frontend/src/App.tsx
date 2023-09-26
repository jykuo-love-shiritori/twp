import { Routes, Route, BrowserRouter } from 'react-router-dom';
import Layout from './pages/Layout';
import Home from './pages/home';
import About from './pages/about';
import EachNews from './pages/news/[newsID]';
import NotFound from './components/NotFound';
import Goods from './pages/goods';
import EachGoods from './pages/goods/[goodsID]';
import Cart from './pages/cart';
import User from './pages/user';
import Login from './pages/user/login';
import Signup from './pages/user/signup';
import Info from './components/Info';
import History from './components/History';
import HistoryEach from './pages/user/history/[historyID]';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path='/login' element={<Login />} />
        <Route path='/signup' element={<Signup />} />

        <Route path='/' element={<Layout />}>
          <Route index element={<Home />} />
          <Route path="/home" element={<Home />} />
          <Route path='/news'>
            <Route path=':news_id' element={<EachNews />} />
          </Route>
          <Route path="/about" element={<About />} />
          <Route path="/goods" element={<Goods />} />
          <Route path='/goods'>
            <Route path=':goods_id' element={<EachGoods />} />
          </Route>
          <Route path='user' element={<User />}>
            <Route index element={<Info />} />
            <Route path='/user/user_info' element={<Info />} />
            <Route path='/user/order_history' element={<History />} />
          </Route>
          <Route path='user/order_history'>
            <Route path=':history_id' element={<HistoryEach />} />
          </Route>
          <Route path='/cart' element={<Cart />} />

          <Route path="*" element={<NotFound />} />
        </Route>
      </Routes>
    </BrowserRouter >
  )
}


export default App
