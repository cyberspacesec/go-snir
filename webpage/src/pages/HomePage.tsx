import React from 'react';
import { Link } from 'react-router-dom';
import styled from 'styled-components';

const HomeContainer = styled.div`
  width: 100%;
`;

const HeroSection = styled.section`
  background: linear-gradient(135deg, var(--dark-color) 0%, #1e3a8a 100%);
  color: white;
  padding: 6rem 2rem;
  text-align: center;
`;

const HeroContent = styled.div`
  max-width: 1200px;
  margin: 0 auto;
`;

const HeroTitle = styled.h1`
  font-size: 3.5rem;
  margin-bottom: 1.5rem;
  
  @media (max-width: 768px) {
    font-size: 2.5rem;
  }
  
  span {
    color: var(--accent-color);
  }
`;

const HeroSubtitle = styled.p`
  font-size: 1.5rem;
  margin-bottom: 2rem;
  max-width: 800px;
  margin-left: auto;
  margin-right: auto;
  
  @media (max-width: 768px) {
    font-size: 1.2rem;
  }
`;

const ButtonContainer = styled.div`
  display: flex;
  justify-content: center;
  gap: 1rem;
  margin-bottom: 3rem;
  
  @media (max-width: 768px) {
    flex-direction: column;
    align-items: center;
  }
`;

const Button = styled(Link)`
  padding: 0.8rem 2rem;
  font-size: 1.1rem;
  border-radius: 4px;
  text-decoration: none;
  font-weight: bold;
  transition: all 0.3s;
  
  @media (max-width: 768px) {
    width: 100%;
    max-width: 300px;
  }
`;

const PrimaryButton = styled(Button)`
  background-color: var(--accent-color);
  color: white;
  
  &:hover {
    background-color: #0e946f;
    transform: translateY(-2px);
  }
`;

const SecondaryButton = styled(Button)`
  background-color: transparent;
  color: white;
  border: 2px solid white;
  
  &:hover {
    background-color: rgba(255, 255, 255, 0.1);
    transform: translateY(-2px);
  }
`;

const FeatureSection = styled.section`
  padding: 5rem 2rem;
  background-color: white;
`;

const SectionTitle = styled.h2`
  font-size: 2.5rem;
  text-align: center;
  margin-bottom: 3rem;
  color: var(--dark-color);
  
  &:after {
    content: '';
    display: block;
    width: 80px;
    height: 4px;
    background-color: var(--accent-color);
    margin: 1rem auto 0;
  }
`;

const FeatureGrid = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
  max-width: 1200px;
  margin: 0 auto;
`;

const FeatureCard = styled.div`
  background: white;
  border-radius: 8px;
  padding: 2rem;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.05);
  transition: transform 0.3s;
  
  &:hover {
    transform: translateY(-5px);
  }
`;

const FeatureIcon = styled.div`
  font-size: 2.5rem;
  margin-bottom: 1rem;
  color: var(--accent-color);
`;

const FeatureTitle = styled.h3`
  font-size: 1.5rem;
  margin-bottom: 1rem;
  color: var(--dark-color);
`;

const FeatureDescription = styled.p`
  color: #666;
  line-height: 1.6;
`;

const DemoSection = styled.section`
  padding: 5rem 2rem;
  background-color: #f8f9fa;
`;

const DemoContent = styled.div`
  max-width: 1200px;
  margin: 0 auto;
  display: flex;
  align-items: center;
  gap: 3rem;
  
  @media (max-width: 992px) {
    flex-direction: column;
  }
`;

const DemoTextContent = styled.div`
  flex: 1;
`;

const DemoTitle = styled.h2`
  font-size: 2.5rem;
  margin-bottom: 1.5rem;
  color: var(--dark-color);
`;

const DemoDescription = styled.p`
  color: #666;
  line-height: 1.8;
  font-size: 1.1rem;
  margin-bottom: 2rem;
`;

const DemoImageContainer = styled.div`
  flex: 1;
  box-shadow: 0 15px 50px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
  overflow: hidden;
`;

const DemoImage = styled.img`
  width: 100%;
  height: auto;
  display: block;
`;

const CTASection = styled.section`
  padding: 5rem 2rem;
  background: linear-gradient(135deg, var(--primary-color) 0%, var(--accent-color) 100%);
  color: white;
  text-align: center;
`;

const CTATitle = styled.h2`
  font-size: 2.5rem;
  margin-bottom: 1.5rem;
`;

const CTADescription = styled.p`
  font-size: 1.2rem;
  max-width: 800px;
  margin: 0 auto 2rem;
`;

const HomePage: React.FC = () => {
  return (
    <HomeContainer>
      <HeroSection>
        <HeroContent>
          <HeroTitle>
            Go-<span>SNIR</span> 网页截图工具
          </HeroTitle>
          <HeroSubtitle>
            强大的网页截图与信息收集工具，为安全研究和网站监控提供全面支持
          </HeroSubtitle>
          <ButtonContainer>
            <PrimaryButton to="/download">立即下载</PrimaryButton>
            <SecondaryButton to="/documentation">查看文档</SecondaryButton>
          </ButtonContainer>
          
          {/* 将来可以添加一张屏幕截图或工具界面图片 */}
        </HeroContent>
      </HeroSection>
      
      <FeatureSection>
        <SectionTitle>核心功能</SectionTitle>
        <FeatureGrid>
          <FeatureCard>
            <FeatureIcon>📸</FeatureIcon>
            <FeatureTitle>批量截图</FeatureTitle>
            <FeatureDescription>
              从单个URL到海量URL列表，甚至整个CIDR网段，都能轻松截图，满足不同规模的需求。
            </FeatureDescription>
          </FeatureCard>
          
          <FeatureCard>
            <FeatureIcon>🔍</FeatureIcon>
            <FeatureTitle>信息收集</FeatureTitle>
            <FeatureDescription>
              不仅仅是截图，还能收集网站标题、响应头、状态码等关键信息，帮助您更全面地了解目标网站。
            </FeatureDescription>
          </FeatureCard>
          
          <FeatureCard>
            <FeatureIcon>🚀</FeatureIcon>
            <FeatureTitle>高并发处理</FeatureTitle>
            <FeatureDescription>
              强大的并发能力，支持同时处理多个目标，大大提高工作效率。
            </FeatureDescription>
          </FeatureCard>
          
          <FeatureCard>
            <FeatureIcon>🧩</FeatureIcon>
            <FeatureTitle>灵活配置</FeatureTitle>
            <FeatureDescription>
              自定义截图分辨率、超时时间、UA信息，满足各种特定需求。
            </FeatureDescription>
          </FeatureCard>
          
          <FeatureCard>
            <FeatureIcon>🌐</FeatureIcon>
            <FeatureTitle>导入导出</FeatureTitle>
            <FeatureDescription>
              支持从Nmap和Nessus扫描结果导入目标，导出结果至多种格式，无缝融入您的工作流程。
            </FeatureDescription>
          </FeatureCard>
          
          <FeatureCard>
            <FeatureIcon>📊</FeatureIcon>
            <FeatureTitle>结果展示</FeatureTitle>
            <FeatureDescription>
              内置Web服务器，让您可以通过浏览器直观地查看和筛选所有截图结果。
            </FeatureDescription>
          </FeatureCard>
        </FeatureGrid>
      </FeatureSection>
      
      <DemoSection>
        <DemoContent>
          <DemoTextContent>
            <DemoTitle>直观强大的命令行工具</DemoTitle>
            <DemoDescription>
              Go-SNIR提供简洁明了的命令行界面，让您能够快速上手，轻松完成各种网页截图任务。无论是安全评估、竞品监控，还是网站内容归档，Go-SNIR都能满足您的需求。
            </DemoDescription>
            <DemoDescription>
              支持多种操作系统，包括Windows、macOS和Linux，无论您使用什么平台，都能轻松部署和使用。
            </DemoDescription>
            <SecondaryButton to="/documentation">了解更多功能</SecondaryButton>
          </DemoTextContent>
          
          <DemoImageContainer>
            {/* 实际项目中应替换为真实的示例截图 */}
            <DemoImage src="https://via.placeholder.com/600x400?text=Go-SNIR+Screenshot+Demo" alt="Go-SNIR工具界面演示" />
          </DemoImageContainer>
        </DemoContent>
      </DemoSection>
      
      <CTASection>
        <CTATitle>开始使用Go-SNIR</CTATitle>
        <CTADescription>
          立即下载并体验这款强大的网页截图工具，提升您的网络安全研究和网站监控效率
        </CTADescription>
        <PrimaryButton to="/download">免费下载</PrimaryButton>
      </CTASection>
    </HomeContainer>
  );
};

export default HomePage; 