import { Routes, Route, BrowserRouter } from 'react-router-dom';

import Layout from '@pages/Layout';
import Home from '@pages/home';
import EachNews from '@pages/news/[newsID]';
import Discover from '@pages/discover';
import EachGoods from '@pages/discover/[goodsID]';
import BuyerCarts from '@pages/user/buyer/cart';
import User from '@pages/user/buyer';
import Login from '@pages/user/login';
import Signup from '@pages/user/signup';
import BuyerHistoryEach from '@pages/user/buyer/history/[historyID]';
import SellerHistoryEach from '@pages/user/seller/allShipments/[historyID]';
import Info from '@pages/user/buyer/info';
import History from '@pages/user/buyer/history';
import Admin from '@pages/user/admin';
import ManageUser from '@pages/user/admin/manageUsers';
import ManageAdminCoupons from '@pages/user/admin/allCoupons';
import NewAdminCoupon from '@pages/user/admin/allCoupons/newCoupon';
import EachAdminCoupon from '@pages/user/admin/allCoupons/[adminCouponID]';
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
import UserViewShop from '@pages/user/shop';
import SellerInfo from '@pages/user/seller/info';
import Products from '@pages/user/seller/allProducts';
import NewGoods from '@pages/user/seller/allProducts/NewGoods';
import ManageSellerCoupons from '@pages/user/seller/allCoupons';
import EachSellerCoupon from '@pages/user/seller/allCoupons/[sellerCouponID]';
import NewSellerCoupon from '@pages/user/seller/allCoupons/newCoupon';
import Authorize from '@pages/user/authorize';
import EachSellerGoods from '@pages/user/seller/allProducts/[sellerGoodsID]';
import SellerShipment from '@pages/user/seller/allShipments';
import Shop from '@pages/user/shop/Shop';
import AdminReport from '@pages/user/admin/allReports';
import AdminReportEach from '@pages/user/admin/allReports/[adminReportID]';
import SellerReport from '@pages/user/seller/allReports';
import SellerReportEach from '@pages/user/seller/allReports/[sellerReportID]';
import Callback from '@pages/user/callback';
import SellerCoupons from '@pages/user/shop/SellerCoupons';
import ProtectedBoundary from '@components/ProtectedBoundary';
import Search from '@pages/search';
import ShopSearch from '@pages/user/shop/ShopSearch';

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path='/login' element={<Login />} />
        <Route path='/authorize' element={<Authorize />} />
        <Route path='/signup' element={<Signup />} />

        <Route path='/callback' element={<Callback />} />

        <Route path='/' element={<ProtectedBoundary />}>
          <Route path='/' element={<Layout />}>
            <Route index element={<Home />} />
            <Route path='/' element={<Home />} />
            <Route path='/news'>
              <Route path=':news_id' element={<EachNews />} />
            </Route>
            <Route path='/discover' element={<Discover />} />
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
              <Route path=':history_id' element={<BuyerHistoryEach />} />
            </Route>

            <Route path='/user/seller' element={<Seller />}>
              <Route index element={<SellerInfo />} />
              <Route path='/user/seller/info' element={<SellerInfo />} />
              <Route path='/user/seller/manageProducts' element={<Products />} />
              <Route path='/user/seller/manageCoupons' element={<ManageSellerCoupons />} />
              <Route path='/user/seller/orders' element={<SellerShipment />} />
              <Route path='/user/seller/reports' element={<SellerReport />} />
            </Route>

            <Route path='/user/seller/manageProducts/new' element={<NewGoods />} />
            <Route path='/user/seller/manageProducts'>
              <Route path=':goods_id' element={<EachSellerGoods />} />
            </Route>

            <Route path='/user/seller/manageCoupons'>
              <Route path='new' element={<NewSellerCoupon />} />
              <Route path=':coupon_id' element={<EachSellerCoupon />} />
            </Route>

            <Route path='shop/:sellerName' element={<UserViewShop />}>
              <Route index element={<Shop />} />
              <Route path='products' element={<Shop />} />
              <Route path='coupons' element={<SellerCoupons />} />
              <Route path='products/inside/search' element={<ShopSearch />} />
            </Route>

            <Route path='/product/:goods_id' element={<EachGoods />} />

            <Route path='/user/seller/order'>
              <Route path=':history_id' element={<SellerHistoryEach />} />
            </Route>

            <Route path='/user/seller/reports'>
              <Route path=':year/:month' element={<SellerReportEach />} />
            </Route>

            <Route path='/buyer/cart' element={<BuyerCarts />} />

            <Route path='/admin' element={<Admin />}>
              <Route index element={<ManageUser />} />
              <Route path='/admin/manageUser' element={<ManageUser />} />
              <Route path='/admin/manageCoupons' element={<ManageAdminCoupons />} />
              <Route path='/admin/reports' element={<AdminReport />} />
            </Route>

            <Route path='/admin/manageCoupons'>
              <Route path='new' element={<NewAdminCoupon />} />
              <Route path=':coupon_id' element={<EachAdminCoupon />} />
            </Route>

            <Route path='/admin/reports'>
              <Route path=':year/:month' element={<AdminReportEach />} />
            </Route>

            <Route path='/searchNotFound' element={<SearchNotFound />} />
            <Route path='/search' element={<Search />} />

            <Route path='*' element={<NotFound />} />
            <Route path='/forbidden' element={<Forbidden />} />
            <Route path='/unauthorized' element={<Unauthorized />} />

            <Route path='/APItest' element={<APItest />} />
          </Route>
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
