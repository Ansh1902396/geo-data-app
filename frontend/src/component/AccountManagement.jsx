// components/AccountManagement.jsx
import React, { useState, useEffect } from "react";
import axios from "axios";

const AccountManagement = () => {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const [success, setSuccess] = useState(false);

  useEffect(() => {
    // Fetch current account details
    const fetchAccountDetails = async () => {
      try {
        const response = await axios.get(
          "https://your-backend-api.com/api/account",
          {
            headers: {
              Authorization: `Bearer ${localStorage.getItem("authToken")}`,
            },
          }
        );

        setUsername(response.data.username);
        setEmail(response.data.email);
      } catch (err) {
        setError("Failed to load account information.");
      }
    };

    fetchAccountDetails();
  }, []);

  const handleUpdate = async (e) => {
    e.preventDefault();

    try {
      const response = await axios.put(
        "https://your-backend-api.com/api/account",
        {
          username,
          email,
          password, // Send new password only if it’s updated
        },
        {
          headers: {
            Authorization: `Bearer ${localStorage.getItem("authToken")}`,
          },
        }
      );

      if (response.status === 200) {
        setSuccess(true);
      }
    } catch (err) {
      setError("Failed to update account information.");
    }
  };

  return (
    <div style={{ textAlign: "center", marginTop: "50px" }}>
      <h2>Account Management</h2>
      <form onSubmit={handleUpdate}>
        <div>
          <input
            type="text"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            placeholder="Username"
            required
            style={{ padding: "10px", marginBottom: "10px", width: "200px" }}
          />
        </div>
        <div>
          <input
            type="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            placeholder="Email"
            required
            style={{ padding: "10px", marginBottom: "10px", width: "200px" }}
          />
        </div>
        <div>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="New Password (optional)"
            style={{ padding: "10px", marginBottom: "10px", width: "200px" }}
          />
        </div>
        <button type="submit" style={{ padding: "10px 20px" }}>
          Update Account
        </button>
        {error && <p style={{ color: "red" }}>{error}</p>}
        {success && (
          <p style={{ color: "green" }}>Account updated successfully!</p>
        )}
      </form>
    </div>
  );
};

export default AccountManagement;
