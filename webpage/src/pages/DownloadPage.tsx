import React from 'react';
import styled from 'styled-components';
import { Link } from 'react-router-dom';

const DownloadContainer = styled.div`
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
  margin-bottom: 4rem;
`;

const SectionTitle = styled.h2`
  font-size: 2rem;
  color: var(--dark-color);
  margin-bottom: 2rem;
  text-align: center;
`;

const DownloadOptionsContainer = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 2rem;
  margin-bottom: 4rem;
`;

const DownloadCard = styled.div`
  background-color: white;
  border-radius: 10px;
  overflow: hidden;
  box-shadow: 0 5px 20px rgba(0, 0, 0, 0.05);
  transition: transform 0.3s;
  
  &:hover {
    transform: translateY(-5px);
    box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
  }
`;

const CardHeader = styled.div`
  background-color: var(--dark-color);
  color: white;
  padding: 2rem;
  text-align: center;
`;

const CardIcon = styled.div`
  font-size: 3rem;
  margin-bottom: 1rem;
`;

const CardTitle = styled.h3`
  font-size: 1.5rem;
  margin-bottom: 0.5rem;
`;

const CardSubtitle = styled.p`
  font-size: 0.9rem;
  opacity: 0.8;
`;

const CardBody = styled.div`
  padding: 2rem;
`;

const CardDescription = styled.p`
  margin-bottom: 2rem;
  color: #555;
  line-height: 1.6;
`;

const DownloadButton = styled.a`
  display: block;
  background-color: var(--accent-color);
  color: white;
  text-align: center;
  padding: 1rem;
  border-radius: 4px;
  text-decoration: none;
  font-weight: bold;
  transition: background-color 0.3s;
  
  &:hover {
    background-color: #0e946f;
  }
`;

const RequirementsSection = styled.div`
  background-color: #f8f9fa;
  padding: 3rem 2rem;
  border-radius: 10px;
  margin-bottom: 4rem;
`;

const RequirementsList = styled.ul`
  max-width: 800px;
  margin: 0 auto;
  list-style-type: none;
  padding: 0;
`;

const RequirementItem = styled.li`
  display: flex;
  align-items: flex-start;
  margin-bottom: 1.5rem;
  
  &:last-child {
    margin-bottom: 0;
  }
`;

const RequirementIcon = styled.div`
  color: var(--accent-color);
  font-size: 1.5rem;
  margin-right: 1rem;
  min-width: 25px;
`;

const RequirementText = styled.div`
  flex: 1;
`;

const RequirementTitle = styled.h3`
  font-size: 1.2rem;
  color: var(--dark-color);
  margin-bottom: 0.5rem;
`;

const RequirementDescription = styled.p`
  color: #666;
  line-height: 1.6;
`;

const NoteBox = styled.div`
  background-color: rgba(84, 160, 255, 0.1);
  border-left: 4px solid var(--primary-color);
  padding: 1.5rem;
  margin: 2rem 0;
  border-radius: 0 4px 4px 0;
`;

const VersionTable = styled.div`
  overflow-x: auto;
  margin-bottom: 2rem;
`;

const Table = styled.table`
  width: 100%;
  border-collapse: collapse;
  
  th, td {
    padding: 1rem;
    text-align: left;
    border-bottom: 1px solid #eee;
  }
  
  th {
    background-color: var(--dark-color);
    color: white;
  }
  
  tr:nth-child(even) {
    background-color: #f9f9f9;
  }
`;

const LinkButton = styled(Link)`
  display: inline-block;
  background-color: var(--primary-color);
  color: white;
  padding: 0.8rem 2rem;
  border-radius: 4px;
  text-decoration: none;
  font-weight: bold;
  transition: background-color 0.3s, transform 0.3s;
  margin-top: 1rem;
  
  &:hover {
    background-color: #1a6dbf;
    transform: translateY(-2px);
  }
`;

const DownloadPage: React.FC = () => {
  const currentVersion = "1.2.0"; // 模拟当前版本号
  
  return (
    <DownloadContainer>
      <PageHeader>
        <PageTitle>下载 Go-SNIR</PageTitle>
        <PageDescription>
          选择适合您操作系统的版本，开始使用这款强大的网页截图工具
        </PageDescription>
      </PageHeader>
      
      <Section>
        <SectionTitle>选择您的平台</SectionTitle>
        <DownloadOptionsContainer>
          <DownloadCard>
            <CardHeader>
              <CardIcon>🪟</CardIcon>
              <CardTitle>Windows</CardTitle>
              <CardSubtitle>64位系统</CardSubtitle>
            </CardHeader>
            <CardBody>
              <CardDescription>
                适用于Windows 10/11的安装包。包含所有依赖，开箱即用。
              </CardDescription>
              <DownloadButton href="https://github.com/cyberspacesec/go-snir/releases/download/v1.2.0/go-snir-windows-amd64.zip" target="_blank" rel="noopener noreferrer">
                下载Windows版 ({currentVersion})
              </DownloadButton>
            </CardBody>
          </DownloadCard>
          
          <DownloadCard>
            <CardHeader>
              <CardIcon>🍎</CardIcon>
              <CardTitle>macOS</CardTitle>
              <CardSubtitle>Intel 和 Apple Silicon</CardSubtitle>
            </CardHeader>
            <CardBody>
              <CardDescription>
                适用于macOS 10.15及以上版本，同时支持Intel和M1/M2芯片。
              </CardDescription>
              <DownloadButton href="https://github.com/cyberspacesec/go-snir/releases/download/v1.2.0/go-snir-macos-universal.dmg" target="_blank" rel="noopener noreferrer">
                下载macOS版 ({currentVersion})
              </DownloadButton>
            </CardBody>
          </DownloadCard>
          
          <DownloadCard>
            <CardHeader>
              <CardIcon>🐧</CardIcon>
              <CardTitle>Linux</CardTitle>
              <CardSubtitle>多种发行版</CardSubtitle>
            </CardHeader>
            <CardBody>
              <CardDescription>
                适用于各种Linux发行版的二进制文件和安装包。
              </CardDescription>
              <DownloadButton href="https://github.com/cyberspacesec/go-snir/releases/download/v1.2.0/go-snir-linux-amd64.tar.gz" target="_blank" rel="noopener noreferrer">
                下载Linux版 ({currentVersion})
              </DownloadButton>
            </CardBody>
          </DownloadCard>
          
          <DownloadCard>
            <CardHeader>
              <CardIcon>🐳</CardIcon>
              <CardTitle>Docker</CardTitle>
              <CardSubtitle>容器化部署</CardSubtitle>
            </CardHeader>
            <CardBody>
              <CardDescription>
                使用Docker容器，无需担心依赖问题，快速部署和使用。
              </CardDescription>
              <DownloadButton href="https://hub.docker.com/r/cyberspacesec/go-snir" target="_blank" rel="noopener noreferrer">
                Docker Hub
              </DownloadButton>
            </CardBody>
          </DownloadCard>
        </DownloadOptionsContainer>
        
        <NoteBox>
          <strong>注意：</strong> 所有下载的二进制文件都已经过数字签名。您可以在下载页面找到相应的校验和文件，以验证下载文件的完整性。
        </NoteBox>
      </Section>
      
      <Section>
        <SectionTitle>系统要求</SectionTitle>
        <RequirementsSection>
          <RequirementsList>
            <RequirementItem>
              <RequirementIcon>✓</RequirementIcon>
              <RequirementText>
                <RequirementTitle>操作系统</RequirementTitle>
                <RequirementDescription>
                  Windows 10/11, macOS 10.15+, 或现代Linux发行版 (Ubuntu 18.04+, Debian 10+, CentOS 8+, Fedora 32+)
                </RequirementDescription>
              </RequirementText>
            </RequirementItem>
            
            <RequirementItem>
              <RequirementIcon>✓</RequirementIcon>
              <RequirementText>
                <RequirementTitle>Chrome/Chromium浏览器</RequirementTitle>
                <RequirementDescription>
                  需要安装Chrome或Chromium浏览器来执行网页截图功能。如果使用Docker版本，则无需单独安装。
                </RequirementDescription>
              </RequirementText>
            </RequirementItem>
            
            <RequirementItem>
              <RequirementIcon>✓</RequirementIcon>
              <RequirementText>
                <RequirementTitle>硬件要求</RequirementTitle>
                <RequirementDescription>
                  至少2GB RAM，现代CPU，以及足够的存储空间用于保存截图。大规模并发截图可能需要更多资源。
                </RequirementDescription>
              </RequirementText>
            </RequirementItem>
            
            <RequirementItem>
              <RequirementIcon>✓</RequirementIcon>
              <RequirementText>
                <RequirementTitle>网络连接</RequirementTitle>
                <RequirementDescription>
                  需要互联网连接以访问目标网站。根据您的使用场景，可能需要调整网络配置或使用代理。
                </RequirementDescription>
              </RequirementText>
            </RequirementItem>
          </RequirementsList>
        </RequirementsSection>
      </Section>
      
      <Section>
        <SectionTitle>历史版本</SectionTitle>
        <VersionTable>
          <Table>
            <thead>
              <tr>
                <th>版本</th>
                <th>发布日期</th>
                <th>主要更新</th>
                <th>下载</th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td>v1.2.0</td>
                <td>2023-10-15</td>
                <td>添加API接口，优化并发性能，支持更多输出格式</td>
                <td><a href="https://github.com/cyberspacesec/go-snir/releases/tag/v1.2.0" target="_blank" rel="noopener noreferrer">下载</a></td>
              </tr>
              <tr>
                <td>v1.1.2</td>
                <td>2023-07-22</td>
                <td>修复Bug，改进Web界面，增强稳定性</td>
                <td><a href="https://github.com/cyberspacesec/go-snir/releases/tag/v1.1.2" target="_blank" rel="noopener noreferrer">下载</a></td>
              </tr>
              <tr>
                <td>v1.1.0</td>
                <td>2023-05-10</td>
                <td>添加CIDR扫描功能，支持从Nmap/Nessus导入</td>
                <td><a href="https://github.com/cyberspacesec/go-snir/releases/tag/v1.1.0" target="_blank" rel="noopener noreferrer">下载</a></td>
              </tr>
              <tr>
                <td>v1.0.0</td>
                <td>2023-03-01</td>
                <td>首次正式发布，基本功能完整</td>
                <td><a href="https://github.com/cyberspacesec/go-snir/releases/tag/v1.0.0" target="_blank" rel="noopener noreferrer">下载</a></td>
              </tr>
            </tbody>
          </Table>
        </VersionTable>
        
        <div style={{ textAlign: 'center' }}>
          <LinkButton to="https://github.com/cyberspacesec/go-snir/releases" target="_blank" rel="noopener noreferrer">
            查看所有版本
          </LinkButton>
        </div>
      </Section>
      
      <Section>
        <SectionTitle>从源码构建</SectionTitle>
        <Paragraph>
          如果您希望从源码编译Go-SNIR，请确保您的系统上安装了Go语言环境（1.16或更高版本）。然后按照以下步骤操作：
        </Paragraph>
        
        <CodeBlock>
{`# 克隆仓库
git clone https://github.com/cyberspacesec/go-snir.git

# 进入项目目录
cd go-snir

# 编译项目
go build

# 运行测试
go test ./...

# 运行程序
./go-snir version`}
        </CodeBlock>
        
        <div style={{ textAlign: 'center', marginTop: '2rem' }}>
          <LinkButton to="/documentation#installation">
            查看详细安装指南
          </LinkButton>
        </div>
      </Section>
    </DownloadContainer>
  );
};

// 额外的、可重用的样式组件
const Paragraph = styled.p`
  line-height: 1.8;
  margin-bottom: 1.5rem;
  color: #444;
  max-width: 800px;
  margin-left: auto;
  margin-right: auto;
`;

const CodeBlock = styled.pre`
  background-color: #f5f5f5;
  padding: 1.5rem;
  border-radius: 4px;
  overflow-x: auto;
  margin-bottom: 1.5rem;
  font-family: 'Courier New', Courier, monospace;
  line-height: 1.5;
  max-width: 800px;
  margin-left: auto;
  margin-right: auto;
`;

export default DownloadPage; 