// src/components/Login.jsx

import React, { useState } from "react";
import { useNavigate } from "react-router-dom";

const Login = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    // TODO: Add your API call for login here
    const token = "dummyAuthToken"; // Replace with real token from API response
    localStorage.setItem("authToken", token);
    navigate("/main"); // Redirect to main page after login
  };

  return (
    <div style={{ textAlign: "center" }}>
      <h2>Login to Your Account</h2>
      <form onSubmit={handleSubmit}>
        <div>
          <label>Email: </label>
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>
        <div>
          <label>Password: </label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        <button
          type="submit"
          style={{ marginTop: "10px", padding: "10px 20px" }}
        >
          Login
        </button>
      </form>
    </div>
  );
};

export default Login;
