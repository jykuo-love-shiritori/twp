import Footer from "../components/Footer"
import NavBar from "../components/NavBar"
import { Outlet } from "react-router-dom"

import '../style/global.css';
import '../components/style.css';

const Layout = () => {
    return (
        <div className="flex-wrapper">
            <NavBar />
            <Outlet />
            <Footer />
        </div>
    )
}

export default Layout
