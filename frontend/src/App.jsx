import React, { useEffect, useState } from "react";
import {
  createBrowserRouter,
  RouterProvider,
  Navigate,
} from "react-router-dom";
import LandingPage from "./component/LandingPage";
import Login from "./component/Login";
import MainPage from "./component/MainPage";
import Layout from "./component/Layout";
import Signup from "./component/Signup";
import AccountManagement from "./component/AccountManagement";
import ErrorPage from "./Error";

// Create a router using createBrowserRouter
const App = () => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  useEffect(() => {
    const token = localStorage.getItem("authToken");
    if (token) {
      setIsAuthenticated(true);
    }
  }, []);

  const router = createBrowserRouter([
    {
      path: "/",
      element: <Layout />, // Wrap in Layout with header
      children: [
        {
          index: true,
          element: isAuthenticated ? <Navigate to="/main" /> : <LandingPage />, // Redirect to main if logged in
        },
        {
          path: "/login",
          element: isAuthenticated ? <Navigate to="/main" /> : <Login />, // Redirect to main if already logged in
        },
        {
          path: "/signup",
          element: <Signup />,
        },
        {
          path: "/account",
          element: isAuthenticated ? (
            <Navigate to="/main" />
          ) : (
            <AccountManagement />
          ),
        },
        {
          path: "/main",
          element: isAuthenticated ? <MainPage /> : <Navigate to="/" />, // Redirect to landing if not authenticated
        },
        {
          path: "*",
          element: <ErrorPage />, // Fallback error page for unknown routes
        },
      ],
    },
  ]);

  return <RouterProvider router={router} />;
};

export default App;
