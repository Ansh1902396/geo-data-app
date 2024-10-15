// src/components/ErrorPage.jsx

import React from "react";
import { useNavigate } from "react-router-dom";

const ErrorPage = () => {
  const navigate = useNavigate();

  return (
    <div style={{ textAlign: "center", padding: "20px" }}>
      <h1>404 - Page Not Found</h1>
      <p>The page you're looking for doesn't exist.</p>
      <button onClick={() => navigate("/")} style={{ padding: "10px 20px" }}>
        Go to Home
      </button>
    </div>
  );
};

export default ErrorPage;
