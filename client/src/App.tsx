import { Outlet } from "react-router-dom";

import TopNavbar from "./components/TopNavbar";
import BottomNavbar from "./components/BottomNavbar";

export default function App() {
  return (
    <>
      <TopNavbar />
      <Outlet />
      <BottomNavbar />
    </>
  );
}
