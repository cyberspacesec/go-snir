package api

import (
	"context"
	"fmt"
	"sync"

	"github.com/cyberspacesec/go-snir/pkg/log"
)

// ConcurrencyLimiter 用于控制并发请求数
type ConcurrencyLimiter struct {
	maxConcurrent int           // 最大并发数
	sema          chan struct{} // 信号量通道
	waitQueue     int           // 等待队列长度
	activeCount   int           // 当前活跃请求数
	waitCount     int           // 当前等待请求数
	mu            sync.Mutex    // 互斥锁，保护计数器
}

// NewConcurrencyLimiter 创建一个新的并发限制器
func NewConcurrencyLimiter(maxConcurrent, waitQueue int) *ConcurrencyLimiter {
	// 确保参数有效
	if maxConcurrent <= 0 {
		maxConcurrent = 10 // 默认值
		log.Warn("设置了无效的最大并发数，使用默认值", "default", maxConcurrent)
	}
	if waitQueue <= 0 {
		waitQueue = 100 // 默认值
		log.Warn("设置了无效的队列大小，使用默认值", "default", waitQueue)
	}

	return &ConcurrencyLimiter{
		maxConcurrent: maxConcurrent,
		sema:          make(chan struct{}, maxConcurrent),
		waitQueue:     waitQueue,
	}
}

// Acquire 尝试获取许可
func (cl *ConcurrencyLimiter) Acquire(ctx context.Context) error {
	cl.mu.Lock()
	// 如果等待队列已满，直接拒绝
	if cl.waitCount >= cl.waitQueue {
		cl.mu.Unlock()
		return fmt.Errorf("服务器繁忙，请求队列已满")
	}

	cl.waitCount++
	cl.mu.Unlock()

	// 尝试获取信号量
	select {
	case cl.sema <- struct{}{}: // 获取到信号量
		cl.mu.Lock()
		cl.waitCount--
		cl.activeCount++
		cl.mu.Unlock()
		return nil
	case <-ctx.Done(): // 请求被取消或超时
		cl.mu.Lock()
		cl.waitCount--
		cl.mu.Unlock()
		return ctx.Err()
	}
}

// Release 释放许可
func (cl *ConcurrencyLimiter) Release() {
	cl.mu.Lock()
	if cl.activeCount > 0 {
		cl.activeCount--
		<-cl.sema // 释放信号量
	}
	cl.mu.Unlock()
}

// Stats 获取当前状态
func (cl *ConcurrencyLimiter) Stats() (active, waiting, maxConcurrent, queueSize int) {
	cl.mu.Lock()
	active = cl.activeCount
	waiting = cl.waitCount
	maxConcurrent = cl.maxConcurrent
	queueSize = cl.waitQueue
	cl.mu.Unlock()
	return
}
