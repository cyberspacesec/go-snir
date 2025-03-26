import React from 'react';
import styled from 'styled-components';
import { Link } from 'react-router-dom';

const AboutContainer = styled.div`
  max-width: 1200px;
  margin: 0 auto;
  padding: 4rem 2rem;
`;

const PageHeader = styled.div`
  text-align: center;
  margin-bottom: 4rem;
`;

const PageTitle = styled.h1`
  font-size: 3rem;
  color: var(--dark-color);
  margin-bottom: 1.5rem;
  
  @media (max-width: 768px) {
    font-size: 2.5rem;
  }
`;

const PageDescription = styled.p`
  font-size: 1.2rem;
  color: #666;
  max-width: 800px;
  margin: 0 auto;
  line-height: 1.6;
`;

const Section = styled.section`
  margin-bottom: 5rem;
`;

const SectionTitle = styled.h2`
  font-size: 2rem;
  color: var(--dark-color);
  margin-bottom: 2rem;
  text-align: center;
  
  &:after {
    content: '';
    display: block;
    width: 80px;
    height: 4px;
    background-color: var(--accent-color);
    margin: 1rem auto 0;
  }
`;

const StorySection = styled.div`
  display: flex;
  align-items: center;
  gap: 3rem;
  margin-bottom: 4rem;
  
  @media (max-width: 992px) {
    flex-direction: column;
  }
`;

const StoryContent = styled.div`
  flex: 1;
`;

const StoryTitle = styled.h3`
  font-size: 1.8rem;
  color: var(--dark-color);
  margin-bottom: 1.5rem;
`;

const StoryText = styled.p`
  line-height: 1.8;
  color: #555;
  margin-bottom: 1.5rem;
`;

const StoryImage = styled.div`
  flex: 1;
  border-radius: 10px;
  overflow: hidden;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
  
  img {
    width: 100%;
    height: auto;
    display: block;
  }
`;

const TeamGrid = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 2rem;
  margin-top: 3rem;
`;

const TeamMember = styled.div`
  text-align: center;
`;

const MemberAvatar = styled.div`
  width: 150px;
  height: 150px;
  border-radius: 50%;
  overflow: hidden;
  margin: 0 auto 1.5rem;
  
  img {
    width: 100%;
    height: 100%;
    object-fit: cover;
  }
`;

const MemberName = styled.h3`
  font-size: 1.3rem;
  color: var(--dark-color);
  margin-bottom: 0.5rem;
`;

const MemberRole = styled.p`
  color: var(--accent-color);
  font-weight: 500;
  margin-bottom: 1rem;
`;

const MemberBio = styled.p`
  font-size: 0.9rem;
  color: #666;
  line-height: 1.6;
`;

const SocialLinks = styled.div`
  display: flex;
  justify-content: center;
  gap: 1rem;
  margin-top: 1rem;
`;

const SocialLink = styled.a`
  color: var(--dark-color);
  font-size: 1.2rem;
  transition: color 0.3s;
  
  &:hover {
    color: var(--accent-color);
  }
`;

const HighlightBox = styled.div`
  background-color: var(--light-color);
  padding: 3rem;
  border-radius: 10px;
  margin-bottom: 3rem;
  text-align: center;
`;

const HighlightTitle = styled.h3`
  font-size: 2rem;
  color: var(--dark-color);
  margin-bottom: 1.5rem;
`;

const HighlightText = styled.p`
  font-size: 1.2rem;
  color: #555;
  max-width: 800px;
  margin: 0 auto 2rem;
  line-height: 1.8;
`;

const StatGrid = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 2rem;
  margin-top: 3rem;
`;

const StatItem = styled.div`
  text-align: center;
`;

const StatNumber = styled.div`
  font-size: 3rem;
  font-weight: bold;
  color: var(--accent-color);
  margin-bottom: 1rem;
`;

const StatLabel = styled.p`
  font-size: 1.1rem;
  color: var(--dark-color);
`;

const ContactSection = styled.div`
  background-color: var(--light-color);
  padding: 3rem;
  border-radius: 10px;
  text-align: center;
`;

const Button = styled(Link)`
  display: inline-block;
  background-color: var(--accent-color);
  color: white;
  padding: 1rem 2rem;
  border-radius: 4px;
  text-decoration: none;
  font-weight: bold;
  transition: all 0.3s;
  
  &:hover {
    background-color: #0e946f;
    transform: translateY(-2px);
  }
`;

const EmailButton = styled.a`
  display: inline-block;
  background-color: var(--accent-color);
  color: white;
  padding: 1rem 2rem;
  border-radius: 4px;
  text-decoration: none;
  font-weight: bold;
  transition: all 0.3s;
  
  &:hover {
    background-color: #0e946f;
    transform: translateY(-2px);
  }
`;

const AboutPage: React.FC = () => {
  return (
    <AboutContainer>
      <PageHeader>
        <PageTitle>关于 Go-SNIR</PageTitle>
        <PageDescription>
          了解我们的历程、使命和背后的团队
        </PageDescription>
      </PageHeader>
      
      <Section>
        <SectionTitle>我们的故事</SectionTitle>
        <StorySection>
          <StoryContent>
            <StoryTitle>始于需求，成于创新</StoryTitle>
            <StoryText>
              Go-SNIR 项目源于网络安全研究中的实际需求。在进行大规模网站安全评估时，我们需要一个能够高效批量截图并收集网站信息的工具，但现有的解决方案要么功能有限，要么性能不足。
            </StoryText>
            <StoryText>
              2023年初，我们决定从零开始，利用 Go 语言的并发优势，打造一款专为安全研究和网站监控设计的高性能工具。我们的目标是创建一个既易用又强大的解决方案，能够满足从个人研究者到企业安全团队的各种需求。
            </StoryText>
            <StoryText>
              在开源社区的支持和贡献下，Go-SNIR 已经发展成为一个功能全面、性能卓越的网页截图与信息收集工具，被广泛应用于网络安全评估、资产管理、竞争情报收集等多个领域。
            </StoryText>
          </StoryContent>
          <StoryImage>
            <img src="https://via.placeholder.com/600x400?text=Our+Story" alt="Go-SNIR 发展历程" />
          </StoryImage>
        </StorySection>
        
        <StorySection>
          <StoryImage>
            <img src="https://via.placeholder.com/600x400?text=Our+Mission" alt="Go-SNIR 使命" />
          </StoryImage>
          <StoryContent>
            <StoryTitle>我们的使命</StoryTitle>
            <StoryText>
              我们的使命是提供最高效、最可靠的网页截图与信息收集解决方案，帮助安全研究人员和组织更好地了解和监控网络资产。
            </StoryText>
            <StoryText>
              我们相信，优秀的工具应该既强大又易用。Go-SNIR 的设计理念是简化复杂任务，让用户能够专注于分析和决策，而不是工具操作。
            </StoryText>
            <StoryText>
              作为一个开源项目，我们致力于保持透明度和社区参与，鼓励用户反馈和贡献，共同推动项目不断完善和创新。
            </StoryText>
          </StoryContent>
        </StorySection>
      </Section>
      
      <Section>
        <SectionTitle>项目团队</SectionTitle>
        <PageDescription style={{ marginBottom: '3rem' }}>
          Go-SNIR 由一群热爱网络安全和开源技术的专业人士创建和维护。我们的团队成员来自不同背景，但都有着共同的热情和目标。
        </PageDescription>
        
        <TeamGrid>
          <TeamMember>
            <MemberAvatar>
              <img src="https://via.placeholder.com/300?text=A" alt="团队成员头像" />
            </MemberAvatar>
            <MemberName>张志强</MemberName>
            <MemberRole>项目发起人 & 核心开发者</MemberRole>
            <MemberBio>
              资深网络安全研究员，Go语言爱好者，拥有10年安全领域经验。喜欢解决复杂问题，热衷于开源贡献。
            </MemberBio>
            <SocialLinks>
              <SocialLink href="#" target="_blank" rel="noopener noreferrer">
                <i className="fab fa-github"></i>
              </SocialLink>
              <SocialLink href="#" target="_blank" rel="noopener noreferrer">
                <i className="fab fa-twitter"></i>
              </SocialLink>
              <SocialLink href="#" target="_blank" rel="noopener noreferrer">
                <i className="fab fa-linkedin"></i>
              </SocialLink>
            </SocialLinks>
          </TeamMember>
          
          <TeamMember>
            <MemberAvatar>
              <img src="https://via.placeholder.com/300?text=B" alt="团队成员头像" />
            </MemberAvatar>
            <MemberName>李明</MemberName>
            <MemberRole>UI/UX设计师</MemberRole>
            <MemberBio>
              专注于创造简洁有效的用户界面，拥有丰富的设计经验。负责Go-SNIR的Web界面设计与用户体验优化。
            </MemberBio>
            <SocialLinks>
              <SocialLink href="#" target="_blank" rel="noopener noreferrer">
                <i className="fab fa-github"></i>
              </SocialLink>
              <SocialLink href="#" target="_blank" rel="noopener noreferrer">
                <i className="fab fa-dribbble"></i>
              </SocialLink>
              <SocialLink href="#" target="_blank" rel="noopener noreferrer">
                <i className="fab fa-behance"></i>
              </SocialLink>
            </SocialLinks>
          </TeamMember>
          
          <TeamMember>
            <MemberAvatar>
              <img src="https://via.placeholder.com/300?text=C" alt="团队成员头像" />
            </MemberAvatar>
            <MemberName>王晓</MemberName>
            <MemberRole>后端开发者</MemberRole>
            <MemberBio>
              专注于高性能系统开发，对并发和分布式系统有深入研究。负责Go-SNIR的核心引擎和性能优化。
            </MemberBio>
            <SocialLinks>
              <SocialLink href="#" target="_blank" rel="noopener noreferrer">
                <i className="fab fa-github"></i>
              </SocialLink>
              <SocialLink href="#" target="_blank" rel="noopener noreferrer">
                <i className="fab fa-stack-overflow"></i>
              </SocialLink>
              <SocialLink href="#" target="_blank" rel="noopener noreferrer">
                <i className="fab fa-medium"></i>
              </SocialLink>
            </SocialLinks>
          </TeamMember>
          
          <TeamMember>
            <MemberAvatar>
              <img src="https://via.placeholder.com/300?text=D" alt="团队成员头像" />
            </MemberAvatar>
            <MemberName>赵婷</MemberName>
            <MemberRole>文档与社区管理</MemberRole>
            <MemberBio>
              技术写作专家，热衷于将复杂概念简化。负责项目文档、教程和社区支持，确保用户能够充分利用Go-SNIR。
            </MemberBio>
            <SocialLinks>
              <SocialLink href="#" target="_blank" rel="noopener noreferrer">
                <i className="fab fa-github"></i>
              </SocialLink>
              <SocialLink href="#" target="_blank" rel="noopener noreferrer">
                <i className="fab fa-twitter"></i>
              </SocialLink>
              <SocialLink href="#" target="_blank" rel="noopener noreferrer">
                <i className="fab fa-medium"></i>
              </SocialLink>
            </SocialLinks>
          </TeamMember>
        </TeamGrid>
      </Section>
      
      <Section>
        <HighlightBox>
          <HighlightTitle>项目成果与影响</HighlightTitle>
          <HighlightText>
            自项目启动以来，Go-SNIR 已经帮助众多安全研究人员和组织提升工作效率，简化网页截图和信息收集流程。
            从个人安全研究到企业级应用，Go-SNIR 正在各个领域发挥作用。
          </HighlightText>
          
          <StatGrid>
            <StatItem>
              <StatNumber>25,000+</StatNumber>
              <StatLabel>下载次数</StatLabel>
            </StatItem>
            <StatItem>
              <StatNumber>300+</StatNumber>
              <StatLabel>GitHub Stars</StatLabel>
            </StatItem>
            <StatItem>
              <StatNumber>50+</StatNumber>
              <StatLabel>代码贡献者</StatLabel>
            </StatItem>
            <StatItem>
              <StatNumber>15+</StatNumber>
              <StatLabel>企业用户</StatLabel>
            </StatItem>
          </StatGrid>
        </HighlightBox>
      </Section>
      
      <Section>
        <SectionTitle>开源与社区</SectionTitle>
        <StorySection>
          <StoryContent>
            <StoryTitle>开源精神</StoryTitle>
            <StoryText>
              Go-SNIR 是一个完全开源的项目，采用 MIT 许可证，允许任何人自由使用、修改和分发。我们相信开源不仅是一种开发模式，更是一种共享知识和共同进步的方式。
            </StoryText>
            <StoryText>
              通过开放源代码，我们希望能够吸引更多人参与到项目中来，共同解决问题，添加新功能，并确保代码质量和安全性。同时，开源也让用户能够根据自己的需求定制工具，更好地满足特定场景的要求。
            </StoryText>
            <StoryText>
              我们欢迎各种形式的贡献，无论是代码提交、文档改进、问题报告还是功能建议，都能帮助项目变得更好。
            </StoryText>
          </StoryContent>
          <StoryImage>
            <img src="https://via.placeholder.com/600x400?text=Open+Source" alt="开源与社区" />
          </StoryImage>
        </StorySection>
      </Section>
      
      <Section>
        <ContactSection>
          <HighlightTitle>联系我们</HighlightTitle>
          <HighlightText>
            如果您有任何问题、建议或合作意向，欢迎随时与我们联系。您可以通过GitHub Issues提交问题，或通过以下方式直接联系我们。
          </HighlightText>
          <EmailButton href="mailto:contact@example.com">发送邮件</EmailButton>
          <div style={{ marginTop: '2rem' }}>
            <SocialLink href="https://github.com/cyberspacesec/go-snir" target="_blank" rel="noopener noreferrer" style={{ fontSize: '2rem', margin: '0 1rem' }}>
              <i className="fab fa-github"></i>
            </SocialLink>
            <SocialLink href="#" target="_blank" rel="noopener noreferrer" style={{ fontSize: '2rem', margin: '0 1rem' }}>
              <i className="fab fa-twitter"></i>
            </SocialLink>
            <SocialLink href="#" target="_blank" rel="noopener noreferrer" style={{ fontSize: '2rem', margin: '0 1rem' }}>
              <i className="fab fa-linkedin"></i>
            </SocialLink>
          </div>
        </ContactSection>
      </Section>
    </AboutContainer>
  );
};

export default AboutPage; 