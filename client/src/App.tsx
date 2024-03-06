import { Outlet } from "react-router-dom";

import TopNavbar from "./components/TopNavbar";
import BottomNavbar from "./components/BottomNavbar";
import SpaceContainer from "./components/SpaceContainer";

export default function App() {
  return (
    <>
      <TopNavbar />
      <Outlet />
      <SpaceContainer />
      <BottomNavbar />
    </>
  );
}
