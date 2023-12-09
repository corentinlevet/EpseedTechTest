import React, { useState } from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";

import Login from "../Login/Login";
import Home from "../Home/Home";
import Navbar from "../Navbar/Navbar";
import Notes from "../Notes/Notes";

import "./App.css";

import 'bootstrap/dist/css/bootstrap.min.css';

function App() {
  const [token, setToken] = useState(localStorage.getItem("token") || "");

  if (!token) {
    return <Login setToken={setToken} />;
  }

  return (
    <div className="wrapper">
      <Router>
        <Navbar />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/notes" element={<Notes />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
