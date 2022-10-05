import React from "react";
import { Outlet } from "react-router-dom";
import "./App.css";
import NavigationBar from "./components/NavigationBar";

const App = () => {
  return (
    <>
      <NavigationBar />
      <Outlet />
    </>
  );
};

export default App;
