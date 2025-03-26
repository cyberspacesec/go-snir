import React from 'react';
import styled from 'styled-components';

const FeaturesContainer = styled.div`
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

const FeaturesGrid = styled.div`
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(350px, 1fr));
  gap: 3rem;
  margin-bottom: 4rem;
  
  @media (max-width: 768px) {
    grid-template-columns: 1fr;
  }
`;

const FeatureCard = styled.div`
  background: white;
  border-radius: 10px;
  overflow: hidden;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.05);
  transition: all 0.3s ease;
  
  &:hover {
    transform: translateY(-5px);
    box-shadow: 0 15px 40px rgba(0, 0, 0, 0.1);
  }
`;

const FeatureImage = styled.div`
  height: 200px;
  background-color: var(--light-color);
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 4rem;
  color: var(--accent-color);
`;

const FeatureContent = styled.div`
  padding: 2rem;
`;

const FeatureTitle = styled.h2`
  font-size: 1.8rem;
  color: var(--dark-color);
  margin-bottom: 1rem;
`;

const FeatureDescription = styled.p`
  color: #666;
  line-height: 1.6;
  margin-bottom: 1.5rem;
`;

const FeatureList = styled.ul`
  list-style-type: none;
  padding: 0;
`;

const FeatureListItem = styled.li`
  margin-bottom: 0.8rem;
  padding-left: 1.5rem;
  position: relative;
  
  &:before {
    content: "✓";
    color: var(--accent-color);
    position: absolute;
    left: 0;
    font-weight: bold;
  }
`;

const SectionTitle = styled.h2`
  font-size: 2.5rem;
  color: var(--dark-color);
  margin-bottom: 2rem;
  margin-top: 5rem;
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

const ComparisonTable = styled.div`
  overflow-x: auto;
  margin-bottom: 4rem;
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
    font-weight: 600;
  }
  
  tr:nth-child(even) {
    background-color: #f8f9fa;
  }
  
  tr:hover {
    background-color: #f1f1f1;
  }
  
  td:first-child {
    font-weight: 600;
  }
`;

const FeaturesPage: React.FC = () => {
  return (
    <FeaturesContainer>
      <PageHeader>
        <PageTitle>Go-SNIR 功能特点</PageTitle>
        <PageDescription>
          深入了解Go-SNIR的核心功能与特点，看看它如何满足您的网页截图与信息收集需求
        </PageDescription>
      </PageHeader>
      
      <FeaturesGrid>
        <FeatureCard>
          <FeatureImage>📸</FeatureImage>
          <FeatureContent>
            <FeatureTitle>多样化截图方式</FeatureTitle>
            <FeatureDescription>
              支持多种目标输入方式，满足从单个URL到大规模资产的不同需求
            </FeatureDescription>
            <FeatureList>
              <FeatureListItem>单个URL精准截图</FeatureListItem>
              <FeatureListItem>从文件批量导入URL</FeatureListItem>
              <FeatureListItem>支持扫描CIDR网段</FeatureListItem>
              <FeatureListItem>从Nmap和Nessus结果导入</FeatureListItem>
              <FeatureListItem>自定义域名解析</FeatureListItem>
            </FeatureList>
          </FeatureContent>
        </FeatureCard>
        
        <FeatureCard>
          <FeatureImage>⚡</FeatureImage>
          <FeatureContent>
            <FeatureTitle>高性能并发处理</FeatureTitle>
            <FeatureDescription>
              采用Go语言的强大并发特性，高效处理大规模扫描任务
            </FeatureDescription>
            <FeatureList>
              <FeatureListItem>多线程并行处理</FeatureListItem>
              <FeatureListItem>可调节并发数量</FeatureListItem>
              <FeatureListItem>优化的资源使用</FeatureListItem>
              <FeatureListItem>支持断点续传</FeatureListItem>
              <FeatureListItem>任务进度实时显示</FeatureListItem>
            </FeatureList>
          </FeatureContent>
        </FeatureCard>
        
        <FeatureCard>
          <FeatureImage>🔍</FeatureImage>
          <FeatureContent>
            <FeatureTitle>全面信息收集</FeatureTitle>
            <FeatureDescription>
              不仅仅是截图，还能收集多种网站信息，助力安全研究
            </FeatureDescription>
            <FeatureList>
              <FeatureListItem>网站标题与图标</FeatureListItem>
              <FeatureListItem>HTTP响应头和状态码</FeatureListItem>
              <FeatureListItem>服务器指纹识别</FeatureListItem>
              <FeatureListItem>网页源码保存</FeatureListItem>
              <FeatureListItem>证书信息提取</FeatureListItem>
            </FeatureList>
          </FeatureContent>
        </FeatureCard>
        
        <FeatureCard>
          <FeatureImage>🎛️</FeatureImage>
          <FeatureContent>
            <FeatureTitle>高度可定制</FeatureTitle>
            <FeatureDescription>
              丰富的配置选项，让您能够根据特定需求自定义工具行为
            </FeatureDescription>
            <FeatureList>
              <FeatureListItem>自定义截图分辨率</FeatureListItem>
              <FeatureListItem>可调节页面加载超时</FeatureListItem>
              <FeatureListItem>自定义User-Agent</FeatureListItem>
              <FeatureListItem>代理服务器支持</FeatureListItem>
              <FeatureListItem>Cookie和会话管理</FeatureListItem>
            </FeatureList>
          </FeatureContent>
        </FeatureCard>
        
        <FeatureCard>
          <FeatureImage>💾</FeatureImage>
          <FeatureContent>
            <FeatureTitle>多样化输出格式</FeatureTitle>
            <FeatureDescription>
              支持多种输出格式，便于与其他工具和流程集成
            </FeatureDescription>
            <FeatureList>
              <FeatureListItem>JSON格式输出</FeatureListItem>
              <FeatureListItem>CSV表格数据</FeatureListItem>
              <FeatureListItem>HTML报告生成</FeatureListItem>
              <FeatureListItem>数据库存储支持</FeatureListItem>
              <FeatureListItem>自定义输出模板</FeatureListItem>
            </FeatureList>
          </FeatureContent>
        </FeatureCard>
        
        <FeatureCard>
          <FeatureImage>🔄</FeatureImage>
          <FeatureContent>
            <FeatureTitle>Web界面与API</FeatureTitle>
            <FeatureDescription>
              提供直观的Web界面和API接口，方便查看和管理结果
            </FeatureDescription>
            <FeatureList>
              <FeatureListItem>内置Web服务器</FeatureListItem>
              <FeatureListItem>截图结果在线浏览</FeatureListItem>
              <FeatureListItem>搜索和筛选功能</FeatureListItem>
              <FeatureListItem>RESTful API接口</FeatureListItem>
              <FeatureListItem>与其他平台集成</FeatureListItem>
            </FeatureList>
          </FeatureContent>
        </FeatureCard>
      </FeaturesGrid>
      
      <SectionTitle>功能对比</SectionTitle>
      <ComparisonTable>
        <Table>
          <thead>
            <tr>
              <th>功能</th>
              <th>Go-SNIR</th>
              <th>竞品A</th>
              <th>竞品B</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>单个URL截图</td>
              <td>✓</td>
              <td>✓</td>
              <td>✓</td>
            </tr>
            <tr>
              <td>批量URL截图</td>
              <td>✓</td>
              <td>✓</td>
              <td>✓</td>
            </tr>
            <tr>
              <td>CIDR网段扫描</td>
              <td>✓</td>
              <td>✗</td>
              <td>✗</td>
            </tr>
            <tr>
              <td>Nmap/Nessus导入</td>
              <td>✓</td>
              <td>✗</td>
              <td>✓</td>
            </tr>
            <tr>
              <td>自定义分辨率</td>
              <td>✓</td>
              <td>✓</td>
              <td>✓</td>
            </tr>
            <tr>
              <td>高并发处理</td>
              <td>✓</td>
              <td>✗</td>
              <td>✓</td>
            </tr>
            <tr>
              <td>代理支持</td>
              <td>✓</td>
              <td>✓</td>
              <td>✓</td>
            </tr>
            <tr>
              <td>Web界面</td>
              <td>✓</td>
              <td>✗</td>
              <td>✓</td>
            </tr>
            <tr>
              <td>API接口</td>
              <td>✓</td>
              <td>✗</td>
              <td>✗</td>
            </tr>
            <tr>
              <td>开源免费</td>
              <td>✓</td>
              <td>✗</td>
              <td>✗</td>
            </tr>
          </tbody>
        </Table>
      </ComparisonTable>
    </FeaturesContainer>
  );
};

export default FeaturesPage; 