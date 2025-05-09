# 并发控制机制

go-snir API服务包含了并发限制机制，用于防止在高负载情况下服务崩溃。这种机制确保系统资源得到合理分配，并提供平滑的服务体验。

## 工作原理

并发控制系统通过以下机制工作：

1. **最大并发限制**：限制同时处理的请求数量，防止系统资源耗尽
2. **请求队列**：当达到最大并发限制时，新请求会进入等待队列
3. **请求超时**：队列中的请求有等待超时时间，避免请求无限期等待
4. **资源管理**：动态追踪活跃请求数和等待请求数
5. **拒绝策略**：当队列已满时，新请求会被拒绝，返回适当的HTTP状态码

## 配置参数

在启动API服务时，可以通过命令行参数配置并发控制：

```bash
# 设置最大并发请求数（默认：10）
go-snir api --max-concurrent=20

# 设置请求队列大小（默认：100）
go-snir api --queue-size=200
```

## 工作流程

1. **请求到达**：当API请求到达时，首先通过并发控制中间件
2. **获取许可**：
   - 如果当前活跃请求数小于最大并发限制，请求获得处理许可并继续
   - 如果达到并发限制，请求进入等待队列（如果队列未满）
   - 如果队列已满，请求被拒绝，返回HTTP 429错误（Too Many Requests）
3. **请求处理**：一旦获得许可，请求被正常处理
4. **释放许可**：请求完成后，释放许可，允许队列中的下一个请求获得处理

## 状态监控

API服务提供了一个状态端点，用于监控并发控制系统的实时状态：

```
GET /stats
```

响应示例：

```json
{
  "success": true,
  "data": {
    "active_requests": 5,
    "waiting_requests": 2,
    "max_concurrent": 10,
    "queue_size": 100,
    "uptime": "3h5m10s",
    "started_at": "2023-06-01T08:30:00Z"
  }
}
```

## 适用场景

并发控制机制在以下场景特别有用：

1. **高流量生产环境**：防止服务在高负载下崩溃
2. **资源受限的环境**：在资源有限的服务器上运行时保护系统
3. **限制截图操作**：截图操作较为耗费资源，通过并发控制避免系统负担过重
4. **防止滥用**：防止API被意外或恶意滥用

## 自适应能力

当前实现提供了基本的并发控制功能。未来可能添加更高级的特性，如：

1. **动态调整**：根据系统负载自动调整最大并发数
2. **优先级队列**：支持不同优先级的请求
3. **客户端限制**：基于IP或API密钥的每客户端限制

## 边缘情况处理

1. **请求取消**：如果客户端取消请求，系统会检测到并释放等待中的许可
2. **超时处理**：等待队列中的请求有5秒超时时间，超时后返回HTTP 503错误
3. **重启恢复**：服务重启时，所有计数器重置，不会累积过期的请求

## 最佳实践

1. **合理设置参数**：基于服务器资源和预期负载设置适当的并发限制
2. **监控状态**：定期检查`/stats`端点，了解系统负载情况
3. **客户端重试**：客户端应该实现退避重试机制，处理429和503响应
4. **负载测试**：在部署前进行负载测试，确定最佳并发参数 