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

// Client pages
import CreateClient from "./pages/Clients/CreateClient.tsx";
import ListClients from "./pages/Clients/ListClients.tsx";
import ClientDetails from "./pages/Clients/ClientDetails.tsx";

// Service pages
import CreateService from "./pages/Services/CreateService.tsx";
import ListServices from "./pages/Services/ListServices.tsx";
import ServiceDetails from "./pages/Services/ServiceDetails.tsx";

// Budget pages
import CreateBudget from "./pages/Budgets/CreateBudget.tsx";
import ListBudgets from "./pages/Budgets/ListBudgets.tsx";
import BudgetDetails from "./pages/Budgets/BudgetDetails.tsx";

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
        path: "/clients/:userID",
        element: (
          <PrivateRoute>
            <ClientDetails />
          </PrivateRoute>
        ),
      },
      {
        path: "/services/create",
        element: (
          <PrivateRoute>
            <CreateService />
          </PrivateRoute>
        ),
      },
      {
        path: "/services/list",
        element: (
          <PrivateRoute>
            <ListServices />
          </PrivateRoute>
        ),
      },
      {
        path: "/services/:userID",
        element: (
          <PrivateRoute>
            <ServiceDetails />
          </PrivateRoute>
        ),
      },
      {
        path: "/budgets/create",
        element: (
          <PrivateRoute>
            <CreateBudget />
          </PrivateRoute>
        ),
      },
      {
        path: "/budgets/list",
        element: (
          <PrivateRoute>
            <ListBudgets />
          </PrivateRoute>
        ),
      },
      {
        path: "/budgets/:budgetID",
        element: (
          <PrivateRoute>
            <BudgetDetails />
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
