import React from "react";
import ReactDOM from "react-dom/client";
import { createBrowserRouter, RouterProvider } from "react-router-dom";

import App from "./App.tsx";
import "./global.css";

import { AuthProvider } from "./contexts/AuthContext.tsx";
import { PrivateRoute, PublicRoute } from "./components/RouteGuard.tsx";

import Register from "./pages/Register.tsx";
import Login from "./pages/Login.tsx";
import ErrorPage from "./pages/ErrorPage.tsx";
import CreateClient from "./pages/Clients/CreateClient.tsx";
import ListClients from "./pages/Clients/ListClients.tsx";
import ClientDetails from "./pages/Clients/ClientDetails.tsx";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      {
        path: "/register",
        element: (
          <PublicRoute>
            <Register />
          </PublicRoute>
        ),
      },
      {
        path: "/login",
        element: (
          <PublicRoute>
            <Login />
          </PublicRoute>
        ),
      },
      {
        path: "/",
        element: (
          <PrivateRoute>
            <ListClients />
          </PrivateRoute>
        ),
      },
      {
        path: "/clients/create",
        element: (
          <PrivateRoute>
            <CreateClient />
          </PrivateRoute>
        ),
      },
      {
        path: "/clients/list",
        element: (
          <PrivateRoute>
            <ListClients />
          </PrivateRoute>
        ),
      },
      {
        path: "/clients/:clientID",
        element: (
          <PrivateRoute>
            <ClientDetails />
          </PrivateRoute>
        ),
      },
    ],
  },
]);

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <AuthProvider>
      <RouterProvider router={router} />
    </AuthProvider>
  </React.StrictMode>
);
