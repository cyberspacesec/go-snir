import React, { useState, useEffect } from 'react';
import { useLocation, Link } from 'react-router-dom';
import styled from 'styled-components';

const DocContainer = styled.div`
  display: flex;
  max-width: 1200px;
  margin: 0 auto;
  padding: 4rem 0;
  
  @media (max-width: 768px) {
    flex-direction: column;
    padding: 2rem 1rem;
  }
`;

const Sidebar = styled.div<{ isOpen: boolean }>`
  width: 280px;
  padding: 0 1.5rem;
  position: sticky;
  top: 90px;
  height: calc(100vh - 90px);
  overflow-y: auto;
  
  @media (max-width: 768px) {
    width: 100%;
    height: auto;
    position: relative;
    top: 0;
    display: ${props => props.isOpen ? 'block' : 'none'};
    margin-bottom: 2rem;
  }
`;

const ToggleButton = styled.button`
  display: none;
  width: 100%;
  padding: 1rem;
  background-color: var(--dark-color);
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 1rem;
  font-weight: bold;
  cursor: pointer;
  margin-bottom: 1rem;
  
  @media (max-width: 768px) {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }
  
  &:after {
    content: '';
    display: inline-block;
    width: 0.5rem;
    height: 0.5rem;
    border-right: 2px solid white;
    border-bottom: 2px solid white;
    transform: rotate(45deg);
  }
`;

const NavSection = styled.div`
  margin-bottom: 2rem;
`;

const NavTitle = styled.h3`
  font-size: 1.2rem;
  margin-bottom: 1rem;
  color: var(--dark-color);
  font-weight: bold;
`;

const NavList = styled.ul`
  list-style: none;
  padding: 0;
  margin: 0;
`;

const NavItem = styled.li`
  margin-bottom: 0.5rem;
`;

const NavLink = styled(Link)<{ active: boolean }>`
  display: block;
  padding: 0.5rem 0;
  color: ${props => props.active ? 'var(--accent-color)' : 'var(--dark-color)'};
  text-decoration: none;
  font-weight: ${props => props.active ? 'bold' : 'normal'};
  border-left: 3px solid ${props => props.active ? 'var(--accent-color)' : 'transparent'};
  padding-left: 1rem;
  transition: all 0.3s;
  
  &:hover {
    color: var(--accent-color);
    border-left-color: var(--accent-color);
  }
`;

const Content = styled.div`
  flex: 1;
  padding: 0 2rem;
  
  @media (max-width: 768px) {
    padding: 0;
  }
`;

const Section = styled.section`
  margin-bottom: 3rem;
  scroll-margin-top: 100px;
`;

const SectionTitle = styled.h1`
  font-size: 2.5rem;
  margin-bottom: 1.5rem;
  color: var(--dark-color);
`;

const Paragraph = styled.p`
  line-height: 1.8;
  margin-bottom: 1.5rem;
  color: #444;
`;

const SubSection = styled.div`
  margin-bottom: 2rem;
`;

const SubTitle = styled.h2`
  font-size: 1.8rem;
  margin-bottom: 1rem;
  color: var(--dark-color);
  scroll-margin-top: 100px;
`;

const CodeBlock = styled.pre`
  background-color: #f5f5f5;
  padding: 1.5rem;
  border-radius: 4px;
  overflow-x: auto;
  margin-bottom: 1.5rem;
  font-family: 'Courier New', Courier, monospace;
  line-height: 1.5;
`;

const Note = styled.div`
  background-color: #e8f4f8;
  border-left: 4px solid var(--primary-color);
  padding: 1.5rem;
  margin: 1.5rem 0;
  border-radius: 0 4px 4px 0;
`;

const DocumentationPage: React.FC = () => {
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);
  const location = useLocation();
  
  // 监听URL的hash变化
  useEffect(() => {
    const hash = location.hash;
    if (hash) {
      const element = document.getElementById(hash.substring(1));
      if (element) {
        element.scrollIntoView({ behavior: 'smooth' });
      }
    } else {
      window.scrollTo(0, 0);
    }
  }, [location]);
  
  const isActive = (id: string) => {
    return location.hash === `#${id}`;
  };
  
  return (
    <DocContainer>
      <ToggleButton onClick={() => setIsSidebarOpen(!isSidebarOpen)}>
        {isSidebarOpen ? '隐藏目录' : '显示目录'}
      </ToggleButton>
      
      <Sidebar isOpen={isSidebarOpen}>
        <NavSection>
          <NavTitle>入门指南</NavTitle>
          <NavList>
            <NavItem>
              <NavLink to="/documentation#getting-started" active={isActive('getting-started')}>
                快速入门
              </NavLink>
            </NavItem>
            <NavItem>
              <NavLink to="/documentation#installation" active={isActive('installation')}>
                安装指南
              </NavLink>
            </NavItem>
            <NavItem>
              <NavLink to="/documentation#basic-usage" active={isActive('basic-usage')}>
                基本使用
              </NavLink>
            </NavItem>
          </NavList>
        </NavSection>
        
        <NavSection>
          <NavTitle>功能说明</NavTitle>
          <NavList>
            <NavItem>
              <NavLink to="/documentation#single-url" active={isActive('single-url')}>
                单个URL截图
              </NavLink>
            </NavItem>
            <NavItem>
              <NavLink to="/documentation#batch-processing" active={isActive('batch-processing')}>
                批量处理
              </NavLink>
            </NavItem>
            <NavItem>
              <NavLink to="/documentation#cidr-scanning" active={isActive('cidr-scanning')}>
                CIDR网段扫描
              </NavLink>
            </NavItem>
            <NavItem>
              <NavLink to="/documentation#import-targets" active={isActive('import-targets')}>
                导入目标
              </NavLink>
            </NavItem>
          </NavList>
        </NavSection>
        
        <NavSection>
          <NavTitle>高级特性</NavTitle>
          <NavList>
            <NavItem>
              <NavLink to="/documentation#concurrency" active={isActive('concurrency')}>
                并发控制
              </NavLink>
            </NavItem>
            <NavItem>
              <NavLink to="/documentation#output-formats" active={isActive('output-formats')}>
                输出格式
              </NavLink>
            </NavItem>
            <NavItem>
              <NavLink to="/documentation#web-server" active={isActive('web-server')}>
                Web服务器
              </NavLink>
            </NavItem>
            <NavItem>
              <NavLink to="/documentation#api" active={isActive('api')}>
                API参考
              </NavLink>
            </NavItem>
          </NavList>
        </NavSection>
        
        <NavSection>
          <NavTitle>其他</NavTitle>
          <NavList>
            <NavItem>
              <NavLink to="/documentation#faq" active={isActive('faq')}>
                常见问题
              </NavLink>
            </NavItem>
            <NavItem>
              <NavLink to="/documentation#troubleshooting" active={isActive('troubleshooting')}>
                故障排除
              </NavLink>
            </NavItem>
            <NavItem>
              <NavLink to="/documentation#contributing" active={isActive('contributing')}>
                贡献指南
              </NavLink>
            </NavItem>
          </NavList>
        </NavSection>
      </Sidebar>
      
      <Content>
        <Section id="getting-started">
          <SectionTitle>快速入门</SectionTitle>
          <Paragraph>
            Go-SNIR 是一个强大的网页截图与信息收集工具，使用 Go 语言开发，旨在提供高效的网页截图功能，同时收集相关信息。
            本指南将帮助您快速上手 Go-SNIR，了解其基本功能和用法。
          </Paragraph>
        </Section>
        
        <Section id="installation">
          <SubTitle>安装指南</SubTitle>
          <Paragraph>
            您可以通过多种方式安装和使用 Go-SNIR。以下是几种常见的安装方法：
          </Paragraph>
          
          <SubSection>
            <h3>从源码安装</h3>
            <Paragraph>
              如果您想从源码安装，需要确保您的系统上已经安装了 Go 语言环境（1.16 或更高版本）。
            </Paragraph>
            <CodeBlock>
{`git clone https://github.com/cyberspacesec/go-snir.git
cd go-snir
go build`}
            </CodeBlock>
          </SubSection>
          
          <SubSection>
            <h3>使用预编译二进制文件</h3>
            <Paragraph>
              您可以直接从 GitHub Releases 页面下载适合您操作系统的预编译二进制文件。
            </Paragraph>
            <CodeBlock>
{`# 下载后解压
chmod +x go-snir
./go-snir version`}
            </CodeBlock>
          </SubSection>
          
          <SubSection>
            <h3>使用 Docker</h3>
            <Paragraph>
              Go-SNIR 也提供了 Docker 镜像，您可以使用 Docker 来运行它，无需担心依赖问题。
            </Paragraph>
            <CodeBlock>
{`docker pull cyberspacesec/go-snir
docker run -it --rm cyberspacesec/go-snir scan single https://example.com`}
            </CodeBlock>
          </SubSection>
          
          <Note>
            注意：Go-SNIR 依赖 Chrome/Chromium 来进行网页截图。如果您通过源码或二进制文件安装，请确保系统中已安装了 Chrome 或 Chromium 浏览器。Docker 版本已包含所有依赖。
          </Note>
        </Section>
        
        <Section id="basic-usage">
          <SubTitle>基本使用</SubTitle>
          <Paragraph>
            Go-SNIR 的命令行界面设计简洁明了，遵循以下模式：
          </Paragraph>
          <CodeBlock>
{`go-snir <命令> <子命令> [选项]`}
          </CodeBlock>
          
          <Paragraph>
            主要命令包括：
          </Paragraph>
          <ul>
            <li><strong>scan</strong>：执行网页截图扫描任务</li>
            <li><strong>report</strong>：管理和查看扫描结果</li>
            <li><strong>config</strong>：配置工具设置</li>
            <li><strong>version</strong>：显示版本信息</li>
            <li><strong>help</strong>：显示帮助信息</li>
          </ul>
        </Section>
        
        <Section id="single-url">
          <SubTitle>单个URL截图</SubTitle>
          <Paragraph>
            最基本的用法是对单个URL进行截图：
          </Paragraph>
          <CodeBlock>
{`go-snir scan single https://example.com`}
          </CodeBlock>
          
          <Paragraph>
            您可以添加各种选项来自定义截图行为：
          </Paragraph>
          <CodeBlock>
{`go-snir scan single https://example.com --screenshot-path ./screenshots --resolution 1920x1080 --timeout 30 --user-agent "Mozilla/5.0 ..." --delay 2`}
          </CodeBlock>
        </Section>
        
        <Section id="batch-processing">
          <SubTitle>批量处理</SubTitle>
          <Paragraph>
            Go-SNIR 支持从文件中批量读取URL列表进行截图：
          </Paragraph>
          <CodeBlock>
{`go-snir scan file -f urls.txt`}
          </CodeBlock>
          
          <Paragraph>
            文件中的每一行应该包含一个URL。您也可以自定义线程数量以提高效率：
          </Paragraph>
          <CodeBlock>
{`go-snir scan file -f urls.txt --threads 10`}
          </CodeBlock>
        </Section>
        
        <Section id="cidr-scanning">
          <SubTitle>CIDR网段扫描</SubTitle>
          <Paragraph>
            Go-SNIR 可以扫描整个CIDR网段，自动发现HTTP/HTTPS服务并进行截图：
          </Paragraph>
          <CodeBlock>
{`go-snir scan cidr -c 192.168.1.0/24 --port 80,443,8080`}
          </CodeBlock>
          
          <Paragraph>
            您可以指定需要扫描的端口，以及是否进行端口扫描：
          </Paragraph>
          <CodeBlock>
{`go-snir scan cidr -c 192.168.1.0/24 --port 80,443,8080-8090 --skip-port-scan`}
          </CodeBlock>
        </Section>
        
        <Section id="faq">
          <SubTitle>常见问题</SubTitle>
          
          <SubSection>
            <h3>Q: Go-SNIR 支持代理服务器吗？</h3>
            <Paragraph>
              是的，Go-SNIR 支持通过代理服务器进行截图。您可以使用 --proxy 参数指定代理地址，例如：
              <code>--proxy http://127.0.0.1:8080</code> 或 <code>--proxy socks5://127.0.0.1:1080</code>
            </Paragraph>
          </SubSection>
          
          <SubSection>
            <h3>Q: 如何自定义截图的分辨率？</h3>
            <Paragraph>
              使用 --resolution 参数设置截图分辨率，格式为"宽x高"，例如：<code>--resolution 1920x1080</code>
            </Paragraph>
          </SubSection>
          
          <SubSection>
            <h3>Q: 如何保存网页的HTML内容？</h3>
            <Paragraph>
              使用 --save-html 参数可以同时保存网页的HTML内容。还可以使用 --save-headers 参数保存HTTP响应头。
            </Paragraph>
          </SubSection>
          
          <SubSection>
            <h3>Q: 如何增加截图超时时间？</h3>
            <Paragraph>
              使用 --timeout 参数设置页面加载超时时间（秒），例如：<code>--timeout 60</code>
            </Paragraph>
          </SubSection>
        </Section>
        
        <Section id="contributing">
          <SubTitle>贡献指南</SubTitle>
          <Paragraph>
            我们欢迎并感谢任何形式的贡献，无论是报告问题、提出功能建议，还是提交代码改进。请遵循以下步骤：
          </Paragraph>
          
          <ol>
            <li>Fork 项目仓库</li>
            <li>创建您的特性分支 (<code>git checkout -b feature/amazing-feature</code>)</li>
            <li>提交您的更改 (<code>git commit -m 'Add some amazing feature'</code>)</li>
            <li>推送到分支 (<code>git push origin feature/amazing-feature</code>)</li>
            <li>创建一个 Pull Request</li>
          </ol>
          
          <Paragraph>
            请确保您的代码遵循项目的代码风格和测试要求。更多详情，请查看项目仓库中的 CONTRIBUTING.md 文件。
          </Paragraph>
        </Section>
      </Content>
    </DocContainer>
  );
};

export default DocumentationPage;