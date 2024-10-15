// components/Layout.jsx
import React from "react";
import { Outlet, Link, useNavigate } from "react-router-dom";
import Header from "./Header";
import Footer from "./Footer";
const Layout = () => {
  // const navigate = useNavigate();
  // const handleLogout = () => {
  //   // Remove auth token from localStorage and redirect to login page
  //   localStorage.removeItem("authToken");
  //   navigate("/login");
  // };

  return (
    <div>
      <Header />
      <main style={mainStyles}>
        <Outlet />
      </main>
      <Footer />
    </div>
  );
};

// Styles for the layout components
const headerStyles = {
  backgroundColor: "#4CAF50",
  padding: "10px",
  color: "white",
};

const navStyles = {
  display: "flex",
  justifyContent: "space-between",
  alignItems: "center",
};

const logoStyles = {
  margin: "0",
  fontSize: "24px",
  fontWeight: "bold",
};

const ulStyles = {
  listStyle: "none",
  margin: "0",
  padding: "0",
  display: "flex",
};

const linkStyles = {
  color: "white",
  textDecoration: "none",
  marginRight: "15px",
};

const logoutButtonStyles = {
  backgroundColor: "transparent",
  border: "none",
  color: "white",
  cursor: "pointer",
  fontSize: "16px",
};

const mainStyles = {
  padding: "20px",
  minHeight: "80vh", // Ensures main content takes most of the viewport
};

const footerStyles = {
  textAlign: "center",
  padding: "10px",
  backgroundColor: "#f1f1f1",
  color: "#555",
};

export default Layout;
