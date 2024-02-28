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

// Customer pages
import CreateCustomer from "./pages/Customers/CreateCustomer.tsx";
import ListCustomers from "./pages/Customers/ListCustomers.tsx";
import CustomerDetails from "./pages/Customers/CustomerDetails.tsx";

// Job pages
import CreateJob from "./pages/Jobs/CreateJob.tsx";
import ListJobs from "./pages/Jobs/ListJobs.tsx";
import JobDetails from "./pages/Jobs/JobDetails.tsx";

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
            <ListCustomers />
          </PrivateRoute>
        ),
      },
      {
        path: "/customers/create",
        element: (
          <PrivateRoute>
            <CreateCustomer />
          </PrivateRoute>
        ),
      },
      {
        path: "/customers/list",
        element: (
          <PrivateRoute>
            <ListCustomers />
          </PrivateRoute>
        ),
      },
      {
        path: "/customers/:customerID",
        element: (
          <PrivateRoute>
            <CustomerDetails />
          </PrivateRoute>
        ),
      },
      {
        path: "/jobs/create",
        element: (
          <PrivateRoute>
            <CreateJob />
          </PrivateRoute>
        ),
      },
      {
        path: "/jobs/list",
        element: (
          <PrivateRoute>
            <ListJobs />
          </PrivateRoute>
        ),
      },
      {
        path: "/jobs/:jobID",
        element: (
          <PrivateRoute>
            <JobDetails />
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
