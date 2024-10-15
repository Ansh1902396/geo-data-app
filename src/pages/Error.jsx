// src/pages/Error.jsx
import React from "react";
import { Link } from "react-router-dom";

const ErrorPage = () => {
  return (
    <div style={errorPageStyles}>
      <h1>404 - Page Not Found</h1>
      <p>Sorry, the page you're looking for does not exist.</p>
      <Link to="/" style={linkStyles}>
        Go Back Home
      </Link>
    </div>
  );
};

// Basic styles for the error page
const errorPageStyles = {
  textAlign: "center",
  padding: "50px",
  fontFamily: "Arial, sans-serif",
};

const linkStyles = {
  color: "#4CAF50",
  textDecoration: "none",
  fontSize: "18px",
};

export default ErrorPage;
