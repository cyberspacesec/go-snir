# Go Web Screenshot

一个用Go语言编写的网页截图工具，可以对网站进行截图并收集相关信息。

## 功能特点

- 对单个URL进行截图
- 批量处理URL列表文件
- 支持扫描CIDR网段
- 支持从Nmap和Nessus扫描结果导入目标
- 自定义截图分辨率和格式
- 可选择保存网页内容和响应头
- 支持多种输出格式（JSON、CSV、数据库）
- 内置Web服务器查看截图结果

## 安装

### 从源码安装

```bash
git clone https://github.com/cyberspacesec/go-snir.git
cd go-web-screenshot
go build
```

### 使用Docker

```bash
docker pull cyberspacesec/go-web-screenshot
docker run -it --rm cyberspacesec/go-web-screenshot scan single https://example.com
```

## 使用方法

### 对单个URL截图

```bash
go-web-screenshot scan single https://example.com
```

### 从文件批量截图

```bash
go-web-screenshot scan file -f urls.txt
```

### 扫描CIDR网段

```bash
go-web-screenshot scan cidr -c 192.168.1.0/24 --port 80,443,8080
```

### 从Nmap XML文件导入

```bash
go-web-screenshot scan nmap -f scan.xml
```

### 启动Web服务器查看结果

```bash
go-web-screenshot report serve
```

## 配置选项

可以通过命令行参数自定义工具的行为：

- `--screenshot-path`: 截图保存路径
- `--resolution`: 截图分辨率，格式为"宽x高"
- `--timeout`: 页面加载超时时间
- `--user-agent`: 自定义User-Agent
- `--chrome-path`: 自定义Chrome路径
- `--delay`: 截图前等待时间
- `--save-html`: 保存网页HTML内容
- `--save-headers`: 保存HTTP响应头

## 许可证

本项目采用MIT许可证。详见[LICENSE](LICENSE)文件。