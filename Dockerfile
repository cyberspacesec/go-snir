# 第一阶段：构建环境
FROM golang:1.22-alpine3.19 AS build

# 设置工作目录
WORKDIR /app

# 安装构建依赖
RUN apk add --no-cache git make

# 设置环境变量
ENV GOTOOLCHAIN=auto

# 复制 Go 模块定义文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制所有源代码
COPY . .

# 执行构建
RUN make build

# 第二阶段：运行环境
FROM alpine:3.19

# 安装运行时依赖
RUN apk add --no-cache ca-certificates chromium

# 设置工作目录
WORKDIR /app

# 创建配置和数据目录
RUN mkdir -p /app/data

# 从构建阶段复制编译好的二进制文件
COPY --from=build /app/snir /app/snir

# 复制必要的资源文件
COPY --from=build /app/webpage /app/webpage

# 设置环境变量
ENV PATH="/app:${PATH}"
ENV CHROME_BIN="/usr/bin/chromium-browser"
ENV CHROME_PATH="/usr/lib/chromium/"

# 暴露 API 端口（如果应用提供 Web 服务）
EXPOSE 8080

# 设置容器启动命令
ENTRYPOINT ["/app/snir"]
CMD ["serve"] 