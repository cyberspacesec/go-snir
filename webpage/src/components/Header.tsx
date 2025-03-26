import React, { useState } from 'react';
import { Link, useLocation } from 'react-router-dom';
import styled from 'styled-components';

const HeaderContainer = styled.header`
  background-color: var(--dark-color);
  color: white;
  padding: 0 2rem;
  position: sticky;
  top: 0;
  z-index: 100;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
`;

const NavContainer = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 70px;
  max-width: 1200px;
  margin: 0 auto;
`;

const Logo = styled(Link)`
  font-size: 1.8rem;
  font-weight: bold;
  color: white;
  text-decoration: none;
  display: flex;
  align-items: center;
  
  span {
    color: var(--accent-color);
  }
`;

const NavLinks = styled.nav<{ isOpen: boolean }>`
  display: flex;
  align-items: center;
  
  @media (max-width: 768px) {
    position: fixed;
    top: 70px;
    left: 0;
    width: 100%;
    background-color: var(--dark-color);
    flex-direction: column;
    align-items: flex-start;
    padding: 1rem;
    transform: ${props => props.isOpen ? 'translateY(0)' : 'translateY(-100%)'};
    opacity: ${props => props.isOpen ? 1 : 0};
    visibility: ${props => props.isOpen ? 'visible' : 'hidden'};
    transition: all 0.3s ease-in-out;
    z-index: 99;
  }
`;

const NavLink = styled(Link)<{ active: boolean }>`
  color: ${props => props.active ? 'var(--accent-color)' : 'white'};
  text-decoration: none;
  padding: 0.5rem 1rem;
  font-weight: ${props => props.active ? 'bold' : 'normal'};
  position: relative;
  
  &:after {
    content: '';
    position: absolute;
    width: ${props => props.active ? '100%' : '0'};
    height: 2px;
    bottom: 0;
    left: 0;
    background-color: var(--accent-color);
    transition: width 0.3s;
  }
  
  &:hover:after {
    width: 100%;
  }
  
  @media (max-width: 768px) {
    padding: 1rem;
    width: 100%;
    border-bottom: 1px solid rgba(255, 255, 255, 0.1);
  }
`;

const DownloadButton = styled(Link)`
  background-color: var(--accent-color);
  color: white;
  padding: 0.5rem 1.2rem;
  border-radius: 4px;
  margin-left: 1rem;
  font-weight: bold;
  text-decoration: none;
  transition: background-color 0.3s;
  
  &:hover {
    background-color: #0e946f;
  }
  
  @media (max-width: 768px) {
    margin: 1rem 0;
    width: 100%;
    text-align: center;
  }
`;

const MenuButton = styled.button`
  display: none;
  background: none;
  border: none;
  color: white;
  font-size: 1.5rem;
  cursor: pointer;
  
  @media (max-width: 768px) {
    display: block;
  }
`;

const Header: React.FC = () => {
  const [isMenuOpen, setIsMenuOpen] = useState(false);
  const location = useLocation();
  
  const toggleMenu = () => {
    setIsMenuOpen(!isMenuOpen);
  };
  
  const isActive = (path: string) => {
    return location.pathname === path;
  };
  
  return (
    <HeaderContainer>
      <NavContainer>
        <Logo to="/">
          Go-<span>SNIR</span>
        </Logo>
        
        <MenuButton onClick={toggleMenu}>
          {isMenuOpen ? '✕' : '☰'}
        </MenuButton>
        
        <NavLinks isOpen={isMenuOpen}>
          <NavLink to="/" active={isActive('/')}>
            首页
          </NavLink>
          <NavLink to="/features" active={isActive('/features')}>
            功能特点
          </NavLink>
          <NavLink to="/documentation" active={isActive('/documentation')}>
            文档
          </NavLink>
          <NavLink to="/about" active={isActive('/about')}>
            关于
          </NavLink>
          <DownloadButton to="/download">
            立即下载
          </DownloadButton>
        </NavLinks>
      </NavContainer>
    </HeaderContainer>
  );
};

export default Header; 