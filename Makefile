.PHONY: build clean install test

# 默认目标
all: build

# 版本信息
VERSION := v0.0.1
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_DATE := $(shell date +%Y-%m-%d)
BUILD_TIME := $(shell date +%H:%M:%S)
LDFLAGS := -ldflags "-X github.com/cyberspacesec/go-snir/pkg/ascii.version=$(VERSION) -X github.com/cyberspacesec/go-snir/pkg/ascii.commit=$(COMMIT) -X github.com/cyberspacesec/go-snir/pkg/ascii.buildDate=$(BUILD_DATE) -X github.com/cyberspacesec/go-snir/pkg/ascii.buildTime=$(BUILD_TIME)"

# 构建可执行文件
build:
	@echo "正在构建 snir..."
	@go build $(LDFLAGS) -o snir

# 安装到系统
install:
	@echo "正在安装 snir..."
	@go install $(LDFLAGS)

# 清理构建结果
clean:
	@echo "正在清理..."
	@rm -f snir
	@rm -f go-snir

# 运行测试
test:
	@echo "正在运行测试..."
	@go test ./...

# 运行测试并生成覆盖率报告
coverage:
	@echo "正在生成测试覆盖率报告..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "覆盖率报告已生成到 coverage.html"

# 帮助信息
help:
	@echo "可用的命令:"
	@echo "  make build      - 构建可执行文件"
	@echo "  make install    - 安装到系统"
	@echo "  make clean      - 清理构建结果"
	@echo "  make test       - 运行测试"
	@echo "  make coverage   - 生成测试覆盖率报告"
	@echo "  make help       - 显示帮助信息" 