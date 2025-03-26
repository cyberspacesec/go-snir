import React from 'react';
import { Link } from 'react-router-dom';
import styled from 'styled-components';

const FooterContainer = styled.footer`
  background-color: var(--dark-color);
  color: white;
  padding: 3rem 2rem;
`;

const FooterContent = styled.div`
  max-width: 1200px;
  margin: 0 auto;
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 2rem;
`;

const FooterColumn = styled.div`
  display: flex;
  flex-direction: column;
`;

const ColumnTitle = styled.h3`
  font-size: 1.2rem;
  margin-bottom: 1rem;
  color: var(--accent-color);
`;

const FooterLink = styled(Link)`
  color: #ddd;
  text-decoration: none;
  margin-bottom: 0.8rem;
  transition: color 0.3s;
  
  &:hover {
    color: var(--accent-color);
  }
`;

const ExternalLink = styled.a`
  color: #ddd;
  text-decoration: none;
  margin-bottom: 0.8rem;
  transition: color 0.3s;
  
  &:hover {
    color: var(--accent-color);
  }
`;

const Copyright = styled.div`
  text-align: center;
  margin-top: 2rem;
  padding-top: 2rem;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  font-size: 0.9rem;
  color: #aaa;
`;

const Footer: React.FC = () => {
  const currentYear = new Date().getFullYear();
  
  return (
    <FooterContainer>
      <FooterContent>
        <FooterColumn>
          <ColumnTitle>Go-SNIR</ColumnTitle>
          <FooterLink to="/">首页</FooterLink>
          <FooterLink to="/features">功能特点</FooterLink>
          <FooterLink to="/documentation">文档</FooterLink>
          <FooterLink to="/download">下载</FooterLink>
          <FooterLink to="/about">关于我们</FooterLink>
        </FooterColumn>
        
        <FooterColumn>
          <ColumnTitle>文档</ColumnTitle>
          <FooterLink to="/documentation#getting-started">快速入门</FooterLink>
          <FooterLink to="/documentation#installation">安装指南</FooterLink>
          <FooterLink to="/documentation#usage">使用教程</FooterLink>
          <FooterLink to="/documentation#api">API参考</FooterLink>
          <FooterLink to="/documentation#faq">常见问题</FooterLink>
        </FooterColumn>
        
        <FooterColumn>
          <ColumnTitle>社区</ColumnTitle>
          <ExternalLink href="https://github.com/cyberspacesec/go-snir" target="_blank" rel="noopener noreferrer">GitHub</ExternalLink>
          <ExternalLink href="https://github.com/cyberspacesec/go-snir/issues" target="_blank" rel="noopener noreferrer">问题反馈</ExternalLink>
          <ExternalLink href="https://github.com/cyberspacesec/go-snir/discussions" target="_blank" rel="noopener noreferrer">讨论区</ExternalLink>
          <ExternalLink href="https://github.com/cyberspacesec/go-snir/releases" target="_blank" rel="noopener noreferrer">版本发布</ExternalLink>
        </FooterColumn>
        
        <FooterColumn>
          <ColumnTitle>联系我们</ColumnTitle>
          <p>我们致力于打造高效的网页截图与信息收集工具。如有任何问题或建议，欢迎随时联系我们。</p>
          <ExternalLink href="mailto:contact@example.com">contact@example.com</ExternalLink>
        </FooterColumn>
      </FooterContent>
      
      <Copyright>
        &copy; {currentYear} Go-SNIR. 保留所有权利。 遵循 <ExternalLink href="https://github.com/cyberspacesec/go-snir/blob/main/LICENSE" target="_blank" rel="noopener noreferrer">MIT 许可证</ExternalLink>
      </Copyright>
    </FooterContainer>
  );
};

export default Footer; 