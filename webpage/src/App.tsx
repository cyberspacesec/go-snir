import React from 'react';
import { HashRouter as Router, Routes, Route } from 'react-router-dom';
import styled from 'styled-components';
import Header from './components/Header.tsx';
import Footer from './components/Footer.tsx';
import HomePage from './pages/HomePage.tsx';
import FeaturesPage from './pages/FeaturesPage.tsx';
import DocumentationPage from './pages/DocumentationPage.tsx';
import DownloadPage from './pages/DownloadPage.tsx';
import AboutPage from './pages/AboutPage.tsx';

const AppContainer = styled.div`
  min-height: 100vh;
  display: flex;
  flex-direction: column;
`;

const ContentContainer = styled.main`
  flex: 1;
`;

function App() {
  return (
    <Router>
      <AppContainer>
        <Header />
        <ContentContainer>
          <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/features" element={<FeaturesPage />} />
            <Route path="/documentation" element={<DocumentationPage />} />
            <Route path="/download" element={<DownloadPage />} />
            <Route path="/about" element={<AboutPage />} />
          </Routes>
        </ContentContainer>
        <Footer />
      </AppContainer>
    </Router>
  );
}

export default App; 