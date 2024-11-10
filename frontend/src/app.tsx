import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import HomePage from './components/homepage';
import LoginPage from './components/loginpage';
import RegisterPage from './components/registerpage';
import ConnectionPage from "./components/connection";
import ChatPage from "./components/chatpage";

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/register" element={<RegisterPage />} />
          <Route path="/connection" element={<ConnectionPage />} />
          <Route path="/chatpage" element={<ChatPage />} />
      </Routes>
    </Router>
  );
};

export default App;
