// src/components/Header.jsx
import React from "react";
import { Link } from "react-router-dom";

const Header = () => {
  return (
    <header
      style={{
        padding: "10px",
        backgroundColor: "#333",
        color: "#fff",
        textAlign: "center",
      }}
    >
      <nav>
        <Link to="/" style={{ margin: "0 15px", color: "#fff" }}>
          Geo Data App
        </Link>
        <Link to="/login" style={{ margin: "0 15px", color: "#fff" }}>
          Login
        </Link>
        <Link to="/signup" style={{ margin: "0 15px", color: "#fff" }}>
          Sign Up
        </Link>
      </nav>
    </header>
  );
};

export default Header;
