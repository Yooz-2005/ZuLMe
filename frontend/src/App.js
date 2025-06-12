import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from './pages/Home';
import LoginRegister from './pages/LoginRegister';
import Dashboard from './pages/Dashboard'; // 加上这一行
import PersonalCenter from './pages/PersonalCenter'; // 加上这一行

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/login-register" element={<LoginRegister />} />
        <Route path="/dashboard" element={<Dashboard />} /> {/* 加上这一行 */}
        <Route path="/personal-center" element={<PersonalCenter />} /> {/* 加上这一行 */}
      </Routes>
    </Router>
  );
}

export default App; 