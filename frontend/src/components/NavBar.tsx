import Nav from 'react-bootstrap/Nav';
import Navbar from 'react-bootstrap/Navbar';

import 'bootstrap/dist/css/bootstrap.min.css'

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faUser, faCartShopping } from "@fortawesome/free-solid-svg-icons";

import '../components/style.css';
import '../style/global.css';

const NavBar = () => {
    return (
        <div className='navbar_twp'>
            <Navbar expand="lg" >
                <Navbar.Brand href="/home" style={{ width: '35px' }}>
                    <img src='/images/logo.png' alt='logo' style={{ width: '100%' }}></img>
                </Navbar.Brand>
                <Navbar.Toggle aria-controls="basic-navbar-nav" />
                <Navbar.Collapse id="basic-navbar-nav">
                    <Nav className="me-auto">
                        <Nav.Link href="/home">
                            <span className='white_word'>Home</span>
                        </Nav.Link>
                        <Nav.Link href="/about" >
                            <span className='white_word'>About</span>
                        </Nav.Link>
                        <Nav.Link href="/goods">
                            <span className='white_word'>Goods</span>
                        </Nav.Link>
                    </Nav>
                    <Nav className="ms-auto">
                        <Nav.Link href="/user" >
                            <FontAwesomeIcon icon={faUser} className='white_word' />
                        </Nav.Link>
                        <Nav.Link href="/cart">
                            <FontAwesomeIcon icon={faCartShopping} className='white_word' />
                        </Nav.Link>
                    </Nav>
                </Navbar.Collapse>
            </Navbar >
        </div >
    )
}

export default NavBar
