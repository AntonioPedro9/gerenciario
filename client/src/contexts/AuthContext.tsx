import { createContext, useState, useEffect, ReactNode } from "react";
import { Navigate } from "react-router-dom";

import api from "../service/api";
import Cookies from "js-cookie";

interface IAuthContext {
  isLoggedIn: boolean;
  setLoginStatus: (value: boolean) => void;
  login: (email: string, password: string) => Promise<void>;
  logout: () => void;
}

export const AuthContext = createContext<IAuthContext>({
  isLoggedIn: false,
  setLoginStatus: () => {},
  login: async () => {},
  logout: () => {},
});

interface AuthProviderProps {
  children: ReactNode;
}

export function AuthProvider({ children }: AuthProviderProps) {
  const [isLoggedIn, setLoginStatus] = useState<boolean>(false);

  useEffect(() => {
    const token = Cookies.get("Authorization");
    if (token) setLoginStatus(true);
  }, []);

  const login = async (email: string, password: string) => {
    try {
      const response = await api.post("/users/login", { email, password });

      if (response.status === 200) {
        Cookies.set("Authorization", response.data.token, { path: "/" });
        setLoginStatus(true);
      } else {
        alert("Falha ao fazer login");
      }
    } catch (error: any) {
      alert(error.response.data.error);
    }
  };

  const logout = () => {
    if (confirm("Você tem certeza que deseja sair?")) {
      Cookies.remove("Authorization");
      setLoginStatus(false);
      <Navigate to="/login" />;
    }
  };

  return <AuthContext.Provider value={{ isLoggedIn, setLoginStatus, login, logout }}>{children}</AuthContext.Provider>;
}
