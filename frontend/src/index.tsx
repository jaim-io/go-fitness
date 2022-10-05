import React from "react";
import ReactDOM from "react-dom/client";
import {
  Route,
  BrowserRouter,
  Routes,
} from "react-router-dom";
import "./index.css";
import App from "./App";
import reportWebVitals from "./reportWebVitals";
import Exercise from "./routes/Exercise";
import { Exercises } from "./routes/Exercises";

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement
);

root.render(
  <React.StrictMode>
    {/* <RouterProvider router={router} /> */}
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<App />}>
          <Route path="/exercise/:id" element={<Exercise />} />
          <Route path="/exercise" element={<Exercises />} />
        </Route>
      </Routes>
    </BrowserRouter>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
