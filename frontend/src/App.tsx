import React from "react";
import { Outlet } from "react-router-dom";
import "./App.css";
import NavigationBar from "./components/NavigationBar";

function App() {
  return (
    <>
      <NavigationBar />
      <Outlet />
    </>
  );
}

export default App;