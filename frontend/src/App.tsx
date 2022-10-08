import React from "react";
import { Outlet } from "react-router-dom";
import NavigationBar from "./components/NavigationBar";

const App = () => {
  return (
    <div className="w-screen h-screen bg-white">
      <NavigationBar />
      <Outlet />
    </div>
  );
};

export default App;
