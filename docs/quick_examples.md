# go-snir 快速使用示例

以下是最常见的 `go-snir` 使用场景，您可以直接复制这些命令：

## 单个网站扫描

**基本扫描**
```bash
./snir scan example.com
```

**增加超时时间和延迟（解决加载慢问题）**
```bash
./snir scan example.com --timeout 60 --delay 3
```

**使用代理**
```bash
./snir scan example.com --proxy http://127.0.0.1:8080
```

**截图 + 保存HTML**
```bash
./snir scan example.com --save-html
```

**截图 + 保存HTTP头信息**
```bash
./snir scan example.com --save-headers
```

**自定义截图保存路径**
```bash
./snir scan example.com --screenshot-path custom_folder
```

## 批量扫描

**从文件扫描**
```bash
./snir scan file -f urls.txt
```

**扫描网段**
```bash
./snir scan cidr 192.168.1.0/24
```

**批量扫描并发数调整**
```bash
./snir scan file -f urls.txt --threads 10
```

## 结果输出

**输出为CSV**
```bash
./snir scan example.com --write-csv
```

**输出为JSONL**
```bash
./snir scan example.com --write-jsonl
```

**保存到数据库**
```bash
./snir scan example.com --db
```

## 常见问题解决

**网站有弹窗干扰截图**
```bash
./snir scan example.com --js "document.querySelectorAll('.popup').forEach(el => el.remove());"
```

**高分辨率截图**
```bash
./snir scan example.com --resolution-x 1920 --resolution-y 1080
```

**显示浏览器界面（非无头模式）**
```bash
./snir scan example.com --headless=false
```

**多次重试失败的网站**
```bash
./snir scan example.com --max-retries 3
``` 