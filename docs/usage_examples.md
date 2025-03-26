# go-snir 使用示例

本文档提供了 `go-snir` 工具的常见使用示例，帮助您快速上手使用而无需记忆众多参数。

## 基础扫描示例

### 1. 扫描单个网站（最简单用法）

```bash
# 最基本的扫描命令，使用默认设置
./snir scan example.com
```

### 2. 扫描单个网站并自定义截图保存位置

```bash
./snir scan example.com --screenshot-path screenshots/custom_folder
```

### 3. 扫描单个网站并增加超时时间（对于加载较慢的网站）

```bash
./snir scan example.com --timeout 60
```

### 4. 扫描单个网站并增加截图前等待时间（确保页面完全加载）

```bash
./snir scan example.com --delay 3
```

### 5. 使用代理扫描网站

```bash
./snir scan example.com --proxy http://127.0.0.1:8080
```

### 6. 自定义浏览器窗口大小

```bash
./snir scan example.com --resolution-x 1920 --resolution-y 1080
```

## 批量扫描示例

### 1. 从文件批量扫描多个网站

```bash
# 从文件 urls.txt 读取网站地址列表进行扫描
./snir scan file -f urls.txt
```

### 2. 扫描整个网段

```bash
# 扫描指定网段的所有主机
./snir scan cidr 192.168.1.0/24
```

### 3. 批量扫描并调整并发数

```bash
./snir scan file -f urls.txt --threads 5
```

## 结果输出示例

### 1. 将结果保存为 CSV 格式

```bash
./snir scan example.com --write-csv
```

### 2. 将结果保存为 JSONL 格式

```bash
./snir scan example.com --write-jsonl
```

### 3. 同时保存 HTML 内容和响应头

```bash
./snir scan example.com --save-html --save-headers
```

### 4. 将结果保存到数据库

```bash
./snir scan example.com --db
```

## 高级使用示例

### 1. 截图前执行 JavaScript 脚本

```bash
./snir scan example.com --js "document.querySelector('.cookie-banner').style.display='none';"
```

### 2. 使用自定义 User-Agent

```bash
./snir scan example.com --user-agent "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
```

### 3. 禁用默认黑名单

```bash
./snir scan example.com --default-blacklist=false
```

### 4. 多次重试扫描失败的网站

```bash
./snir scan example.com --max-retries 3
```

### 5. 使用非无头模式（显示浏览器界面）

```bash
./snir scan example.com --headless=false
```

## 常见问题解决方案

### 网站加载过慢导致超时

增加超时时间和延迟：

```bash
./snir scan slow-website.com --timeout 60 --delay 5
```

### 无法正确加载需要登录的网站

使用已登录的 Cookie：

```bash
./snir scan members-only-site.com --cookie-file cookies.json
```

### 截图中出现弹窗或通知

截图前执行 JavaScript 清除干扰元素：

```bash
./snir scan example.com --js "document.querySelectorAll('.popup, .notification').forEach(el => el.remove());"
```

---

以上示例涵盖了大多数常见的使用场景。您可以根据需要组合使用多个参数，满足具体的扫描需求。 