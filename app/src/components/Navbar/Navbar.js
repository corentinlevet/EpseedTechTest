import { useRef } from 'react';
import { FaBars, FaTimes } from 'react-icons/fa';
import { Link } from 'react-router-dom';

import './Navbar.css';

import Button from 'react-bootstrap/Button'

function Navbar() {
  const navRef = useRef(null);

  const showNavbar = () => {
    navRef.current.classList.toggle("responsive-nav");
  }

  const logOut = () => {
    localStorage.setItem("token", "");
    localStorage.setItem("current_user_id", "");
    window.location.reload();
  }

  return (
    <header id="navHeader">
      <h3>Epseed Tech Test</h3>

      <nav id="navBar" ref={navRef}>
        <Link to="/">Home</Link>
        <Link to="/notes">My notes</Link>
        <button className="nav-btn nav-close-btn" onClick={showNavbar}>
          <FaTimes />
        </button>
        <Button variant="danger" onClick={logOut}>Logout</Button>
      </nav>
      <button id="navBurger" className="nav-btn" onClick={showNavbar}>
        <FaBars />
      </button>
    </header>
  );
}

export default Navbar;
