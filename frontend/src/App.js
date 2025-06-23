import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { ConfigProvider } from 'antd';
import zhCN from 'antd/locale/zh_CN';
import Home from './pages/Home';
import LoginRegister from './pages/LoginRegister';
import VehicleList from './pages/VehicleList';
import VehicleDetail from './pages/VehicleDetail';
import SearchResults from './pages/SearchResults';

import Dashboard from './pages/Dashboard';
import PersonalCenter from './pages/PersonalCenter';
import ImageTest from './components/ImageTest';
import VehicleDebug from './pages/VehicleDebug';
import PaymentSuccess from './pages/PaymentSuccess';
import PaymentTest from './pages/PaymentTest';
import { createGlobalStyle } from 'styled-components';

const GlobalStyle = createGlobalStyle`
  * {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
  }

  body {
    font-family: 'SimHei', 'Microsoft YaHei', Arial, -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue',
      'Noto Sans', sans-serif, 'Apple Color Emoji', 'Segoe UI Emoji', 'Segoe UI Symbol',
      'Noto Color Emoji';
  }
`;

function App() {
  return (
    <ConfigProvider locale={zhCN}>
      <GlobalStyle />
      <Router>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login-register" element={<LoginRegister />} />
          <Route path="/vehicles" element={<VehicleList />} />
          <Route path="/vehicle/:id" element={<VehicleDetail />} />
          <Route path="/search" element={<SearchResults />} />

          <Route path="/dashboard" element={<Dashboard />} />
          <Route path="/personal-center" element={<PersonalCenter />} />
          <Route path="/payment/success" element={<PaymentSuccess />} />
          <Route path="/payment/test" element={<PaymentTest />} />
          <Route path="/image-test" element={<ImageTest />} />
          <Route path="/vehicle-debug" element={<VehicleDebug />} />
        </Routes>
      </Router>
    </ConfigProvider>
  );
}

export default App;