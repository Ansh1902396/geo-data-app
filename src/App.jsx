// src/App.jsx
import React from "react";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import Layout from "./pages/Layout";
import Login from "./pages/Login";
import Signup from "./pages/Signup";
import AccountManagement from "./pages/AccountManagement";
import MainPage from "./pages/MainPage";
import ErrorPage from "./pages/Error";
const checkAuth = () => {
  const token = localStorage.getItem("authToken");
  return !!token;
};
// Create routes
const router = createBrowserRouter([
  {
    path: "/",
    element: checkAuth() ? <Layout /> : <Login />,
    errorElement: <ErrorPage />,
    children: [
      {
        path: "/",
        element: <MainPage />, // Main content after login
      },
      {
        path: "/account",
        element: <AccountManagement />,
      },
    ],
  },
  {
    path: "/login",
    element: <Login />,
  },
  {
    path: "/signup",
    element: <Signup />,
  },
]);

function App() {
  return <RouterProvider router={router} />;
}

export default App;
