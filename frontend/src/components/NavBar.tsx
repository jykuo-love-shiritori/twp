import 'bootstrap/dist/css/bootstrap.min.css';
import '@components/style.css';
import '@style/global.css';

import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import { Row, Col, NavbarBrand, Button } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faUser, faCartShopping } from '@fortawesome/free-solid-svg-icons';
import { Link } from 'react-router-dom';

import SearchBar from '@components/SearchBar';

const NavBar = () => {
  return (
    <div className='navbar_twp'>
      <Navbar expand='lg' style={{ padding: '0px 8% 0px 8%' }}>
        <NavbarBrand href='/' className='disappear_desktop'>
          <img src='/images/logo.png' alt='logo' style={{ width: '35px' }} />
        </NavbarBrand>
        <Nav>
          <div className='disappear_desktop'>
            <SearchBar />
          </div>
        </Nav>
        <Navbar.Toggle aria-controls='basic-navbar-nav' />
        <Navbar.Collapse id='basic-navbar-nav' className='navbar-dark'>
          <Nav className='mt-auto'>
            <Link to='/user/seller' className='nav_link none'>
              Sell
            </Link>
            <Link to='/coupons' className='nav_link none'>
              Coupons
            </Link>
            <Link to='/admin' className='nav_link none'>
              Admin
            </Link>
          </Nav>
          <Nav className='ms-auto'>
            <Link to='/user' className='nav_link none'>
              <FontAwesomeIcon icon={faUser} />
            </Link>
            <Link to='/buyer/cart' className='nav_link none'>
              <FontAwesomeIcon icon={faCartShopping} />
            </Link>
          </Nav>
        </Navbar.Collapse>
      </Navbar>

      <div className='disappear_phone'>
        <hr style={{ color: 'white', opacity: '0.5', margin: '5px' }} />

        <Row className='center' style={{ padding: '0px 8% 0px 8%' }}>
          <Col sm={3}>
            <Link to='/' className='none'>
              <div className='center_vertical'>
                <img src='/images/logo.png' alt='logo' style={{ width: '35px' }} />
                &nbsp;&nbsp; <span className='nav_title'>Too White Powder</span>
              </div>
            </Link>
          </Col>
          <Col sm={6}>
            <SearchBar />
          </Col>
          <Col sm={3}>
            <Button className='search_button center'>Search</Button>
          </Col>
        </Row>
      </div>
    </div>
  );
};

export default NavBar;
