import { ReactNode, useContext } from "react";
import { Navigate } from "react-router-dom";
import { AuthContext } from "../contexts/AuthContext";

interface RouteGuardProps {
  children: ReactNode;
}

export function PublicRoute({ children }: RouteGuardProps) {
  const { isLoggedIn } = useContext(AuthContext);

  if (isLoggedIn) {
    return <Navigate to="/" />;
  }

  return children;
}

export function PrivateRoute({ children }: RouteGuardProps) {
  const { isLoggedIn } = useContext(AuthContext);

  if (!isLoggedIn) {
    return <Navigate to="/login" />;
  }

  return children;
}
