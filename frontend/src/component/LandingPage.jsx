// src/components/LandingPage.jsx

import React from "react";
import { useNavigate } from "react-router-dom";

const LandingPage = () => {
  const navigate = useNavigate();

  const handleLoginClick = () => {
    navigate("/login");
  };

  return (
    <div style={{ textAlign: "center", padding: "20px" }}>
      <h1>Welcome to Geo Data App</h1>
      <p>
        The Geo Data App allows you to manage and visualize geospatial data with
        ease. Upload GeoJSON or KML files, draw custom shapes on a map, and save
        them to your account.
      </p>
      <p>Get started by logging in to your account!</p>
      <button
        onClick={handleLoginClick}
        style={{ padding: "10px 20px", fontSize: "16px" }}
      >
        Login
      </button>
    </div>
  );
};

export default LandingPage;
