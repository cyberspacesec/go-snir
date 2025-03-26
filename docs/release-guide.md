# Go-SNIR 发布指南

本文档描述了如何使用GoReleaser发布Go-SNIR的新版本。

## 先决条件

1. 安装GoReleaser：

```bash
# 使用Homebrew (MacOS)
brew install goreleaser

# 使用Go安装
go install github.com/goreleaser/goreleaser@latest
```

2. 确保你有适当的GitHub权限来推送新标签并创建Releases。

## 发布流程

### 1. 更新版本号

首先，更新`Makefile`中的`VERSION`变量：

```makefile
VERSION := v1.0.0  # 将此更改为新版本号
```

### 2. 更新CHANGELOG

如果你维护一个更改日志，确保更新CHANGELOG.md文件，添加新版本的变更内容。

### 3. 提交所有更改

```bash
git add .
git commit -m "准备发布v1.0.0"
```

### 4. 创建新的Git标签

```bash
git tag -a v1.0.0 -m "发布v1.0.0版本"
```

### 5. 推送标签到GitHub

```bash
git push origin v1.0.0
```

这会触发GitHub Actions工作流，自动使用GoReleaser构建和发布新版本。

### 6. 手动发布（可选）

如果你想手动发布而不通过GitHub Actions，可以运行：

```bash
# 测试发布过程（不会实际发布）
make release-test

# 实际发布
make release
```

## 验证发布

发布完成后，查看GitHub Releases页面，确认：

1. 所有平台的二进制文件都已上传
2. 更改日志已正确显示
3. 下载链接工作正常

## 发布后任务

1. 更新文档中的版本号引用
2. 通知用户有新版本可用
3. 在Homebrew tap中更新版本（如果适用）

## 故障排除

如果遇到发布问题：

1. 检查GitHub Actions日志中的错误
2. 确保`.goreleaser.yml`配置正确
3. 验证你有正确的GitHub权限
4. 确保`GITHUB_TOKEN`有足够的权限

如需更多帮助，请参考[GoReleaser文档](https://goreleaser.com/) 