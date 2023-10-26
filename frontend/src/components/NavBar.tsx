import 'bootstrap/dist/css/bootstrap.min.css';
import '../components/style.css';
import '../style/global.css';

import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import { Row, Col, NavbarBrand } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faUser, faCartShopping } from '@fortawesome/free-solid-svg-icons';

import SearchBar from './SearchBar';

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
            <Nav.Link href='/user/seller'>Sell</Nav.Link>
            <Nav.Link href='/coupons'>Coupons</Nav.Link>
          </Nav>
          <Nav className='ms-auto'>
            <Nav.Link href='/user'>
              <FontAwesomeIcon icon={faUser} />
            </Nav.Link>
            <Nav.Link href='/buyer/cart'>
              <FontAwesomeIcon icon={faCartShopping} />
            </Nav.Link>
          </Nav>
        </Navbar.Collapse>
      </Navbar>

      <div className='disappear_phone'>
        <hr style={{ color: 'white', opacity: '0.5', margin: '5px' }} />

        <Row className='center' style={{ padding: '0px 8% 0px 8%' }}>
          <Col sm={3}>
            <a href='/' className='none'>
              <div className='center_vertical'>
                <img src='/images/logo.png' alt='logo' style={{ width: '35px' }} />
                &nbsp;&nbsp; <span className='nav_title'>Too White Powder</span>
              </div>
            </a>
          </Col>
          <Col sm={6}>
            <SearchBar />
          </Col>
          <Col sm={3}>
            <div className='search_button center'>Search</div>
          </Col>
        </Row>
      </div>
    </div>
  );
};

export default NavBar;
