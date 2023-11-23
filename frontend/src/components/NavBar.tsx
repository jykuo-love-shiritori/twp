import 'bootstrap/dist/css/bootstrap.min.css';
import '@components/style.css';
import '@style/global.css';

import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';
import { Row, Col, NavbarBrand, Button, Dropdown } from 'react-bootstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faUser, faCartShopping } from '@fortawesome/free-solid-svg-icons';
import { Link } from 'react-router-dom';

import SearchBar from '@components/SearchBar';

const NavBar = () => {
  const DropDownStyle = {
    borderRadius: '25px',
    border: '1px solid var(--border)',
    background: ' var(--layout)',
    padding: '10% 5% 10% 5%',
    color: 'white',
  };

  const DropButtonStyle = {
    background: ' var(--layout)',
    border: 'none',
  };

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
        <Navbar.Toggle aria-controls='seller-nav' />
        <Navbar.Collapse id='seller-nav' className='navbar-dark'>
          <Nav className='mt-auto'>
            <Dropdown>
              <Dropdown.Toggle id='dropdown-custom-1' style={DropButtonStyle} className='nav_link'>
                Sell
              </Dropdown.Toggle>
              <Dropdown.Menu style={DropDownStyle}>
                <Link to='/user/seller/info' className='none nav_link' style={{ padding: '0' }}>
                  <div style={{ padding: '5px 10% 5px 10%' }}>Shop Info</div>
                </Link>
                <Link
                  to='/user/seller/manageProducts'
                  className='none nav_link'
                  style={{ padding: '0%' }}
                >
                  <div style={{ padding: '5px 10% 5px 10%' }}>All Products</div>
                </Link>
                <Link
                  to='/user/seller/manageCoupons'
                  className='none nav_link'
                  style={{ padding: '0%' }}
                >
                  <div style={{ padding: '5px 10% 5px 10%' }}>All Coupons</div>
                </Link>
                <Link to='/user/seller/orders' className='none nav_link' style={{ padding: '0%' }}>
                  <div style={{ padding: '5px 10% 5px 10%' }}>All Shipments</div>
                </Link>
                <Link to='/user/seller/reports' className='none nav_link' style={{ padding: '0%' }}>
                  <div style={{ padding: '5px 10% 5px 10%' }}>All Reports</div>
                </Link>
              </Dropdown.Menu>
            </Dropdown>
            <Link to='/coupons' className='nav_link none' style={{ paddingLeft: '10px' }}>
              Coupons
            </Link>
          </Nav>
          <Nav className='ms-auto'>
            <Dropdown>
              <Dropdown.Toggle id='dropdown-custom-1' style={DropButtonStyle} className='nav_link'>
                <FontAwesomeIcon icon={faUser} />
              </Dropdown.Toggle>
              <Dropdown.Menu style={DropDownStyle}>
                <Link to='/user/info' className='none nav_link' style={{ padding: '0' }}>
                  <div style={{ padding: '5px 10% 5px 10%' }}>Personal Info</div>
                </Link>
                <Link to='/user/security' className='none nav_link' style={{ padding: '0%' }}>
                  <div style={{ padding: '5px 10% 5px 10%' }}>Security</div>
                </Link>
                <Link to='/user/buyer/order' className='none nav_link' style={{ padding: '0%' }}>
                  <div style={{ padding: '5px 10% 5px 10%' }}>Order History</div>
                </Link>
                <hr className='hr' style={{ padding: '0', margin: '5px' }} />
                <Link to='/login' className='none nav_link' style={{ padding: '0%' }}>
                  <div style={{ padding: '5px 10% 5px 10%' }}>Logout</div>
                </Link>
              </Dropdown.Menu>
            </Dropdown>
            <Link to='/buyer/cart' className='nav_link none' style={{ paddingLeft: '10px' }}>
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
