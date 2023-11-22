import { Outlet } from "react-router-dom";

import CustomNavbar from "./components/CustomNavbar";

export default function App() {
  return (
    <>
      <CustomNavbar />
      <Outlet />
    </>
  );
}
