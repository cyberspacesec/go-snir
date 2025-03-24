package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/cyberspacesec/go-snir/pkg/log"
	"github.com/gorilla/mux"
)

// 全局并发限制器
var (
	limiterMu          sync.Mutex
	activeRequests     int           // 当前活跃请求数
	waitingRequests    int           // 等待中的请求数
	maxConcurrent      = 10          // 默认最大并发数
	maxQueueSize       = 100         // 默认等待队列大小
	concurrencySemaCh  chan struct{} // 信号量通道
	limiterInitialized bool
	startTime          = time.Now()
)

// 初始化全局并发限制器
func initConcurrencyLimiter(max, queueSize int) {
	limiterMu.Lock()
	defer limiterMu.Unlock()

	if limiterInitialized {
		return
	}

	if max <= 0 {
		max = 10
	}

	if queueSize <= 0 {
		queueSize = 100
	}

	maxConcurrent = max
	maxQueueSize = queueSize
	concurrencySemaCh = make(chan struct{}, max)
	limiterInitialized = true

	log.Info("初始化并发限制器", "max_concurrent", max, "queue_size", queueSize)
}

// 尝试获取并发许可
func acquireConcurrencyPermit(ctx context.Context) error {
	if !limiterInitialized {
		return nil // 未初始化，直接通过
	}

	limiterMu.Lock()
	// 检查等待队列是否已满
	if waitingRequests >= maxQueueSize {
		limiterMu.Unlock()
		return fmt.Errorf("服务器繁忙，请求队列已满")
	}

	waitingRequests++
	limiterMu.Unlock()

	// 尝试获取信号量
	select {
	case concurrencySemaCh <- struct{}{}:
		limiterMu.Lock()
		waitingRequests--
		activeRequests++
		limiterMu.Unlock()
		return nil
	case <-ctx.Done():
		limiterMu.Lock()
		waitingRequests--
		limiterMu.Unlock()
		return ctx.Err()
	}
}

// 释放并发许可
func releaseConcurrencyPermit() {
	if !limiterInitialized {
		return
	}

	limiterMu.Lock()
	if activeRequests > 0 {
		activeRequests--
		<-concurrencySemaCh
	}
	limiterMu.Unlock()
}

// 获取并发限制器状态
func getConcurrencyStats() (active, waiting, max, queue int, uptime time.Duration) {
	limiterMu.Lock()
	active = activeRequests
	waiting = waitingRequests
	max = maxConcurrent
	queue = maxQueueSize
	limiterMu.Unlock()
	uptime = time.Since(startTime)
	return
}

// 创建并发限制中间件
func createConcurrencyLimitMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 跳过对状态检查和静态资源的限制
			if r.URL.Path == "/health" || r.URL.Path == "/stats" ||
				r.URL.Path == "/" || r.URL.Path == "/favicon.ico" ||
				r.URL.Path == "/favicon.png" ||
				r.Method == http.MethodOptions ||
				strings.HasPrefix(r.URL.Path, "/screenshots/") {
				next.ServeHTTP(w, r)
				return
			}

			// 设置超时上下文
			ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
			defer cancel()

			// 尝试获取许可
			err := acquireConcurrencyPermit(ctx)
			if err != nil {
				if err == context.DeadlineExceeded {
					log.Warn("请求等待超时", "path", r.URL.Path, "method", r.Method)
					SendJSONResponse(w, http.StatusServiceUnavailable, APIResponse{
						Success: false,
						Error:   "服务器繁忙，请稍后重试",
					})
				} else {
					log.Warn("请求被拒绝，队列已满", "path", r.URL.Path, "method", r.Method)
					SendJSONResponse(w, http.StatusTooManyRequests, APIResponse{
						Success: false,
						Error:   "服务器繁忙，请求队列已满，请稍后重试",
					})
				}
				return
			}

			// 请求完成后释放许可
			defer releaseConcurrencyPermit()

			// 继续处理请求
			next.ServeHTTP(w, r)
		})
	}
}

// Stats处理器 - 获取服务器状态
func handleStats(w http.ResponseWriter, r *http.Request) {
	active, waiting, max, queue, uptime := getConcurrencyStats()

	SendJSONResponse(w, http.StatusOK, APIResponse{
		Success: true,
		Data: map[string]interface{}{
			"active_requests":  active,
			"waiting_requests": waiting,
			"max_concurrent":   max,
			"queue_size":       queue,
			"uptime":           uptime.String(),
			"started_at":       startTime.Format(time.RFC3339),
		},
	})
}

// Health检查处理器
func handleHealth(w http.ResponseWriter, r *http.Request) {
	SendJSONResponse(w, http.StatusOK, APIResponse{
		Success: true,
		Message: "服务正常运行",
	})
}

// NewServer 创建一个新的API服务器
func NewServer(options Options) *Server {
	router := mux.NewRouter()

	// 初始化并发限制器
	initConcurrencyLimiter(options.MaxConcurrentRequests, options.RequestQueueSize)

	return &Server{
		Options: options,
		Router:  router,
	}
}

// SetupRoutes 设置API路由
func (s *Server) SetupRoutes() {
	// 添加API密钥验证中间件到所有API请求
	apiAuth := s.CreateAuthMiddleware()

	// 添加并发限制中间件
	s.Router.Use(createConcurrencyLimitMiddleware())

	// 应用认证中间件
	s.Router.Use(apiAuth)

	// 设置API端点
	s.Router.HandleFunc("/screenshot", s.HandleScreenshot).Methods("POST")
	s.Router.HandleFunc("/batch", s.HandleBatchScreenshot).Methods("POST")
	s.Router.HandleFunc("/screenshots_list", s.HandleListScreenshots).Methods("GET")
	s.Router.HandleFunc("/get_screenshot/{filename}", s.HandleGetScreenshot).Methods("GET")

	// 设置静态文件服务
	s.Router.PathPrefix("/screenshots/").Handler(http.StripPrefix("/screenshots/", http.FileServer(http.Dir(s.Options.ScreenshotPath))))

	// 添加一个不需要认证的路由，显示API信息
	s.Router.HandleFunc("/", s.HandleRoot).Methods("GET")

	// 添加状态监控和健康检查端点
	s.Router.HandleFunc("/stats", handleStats).Methods("GET")
	s.Router.HandleFunc("/health", handleHealth).Methods("GET")
}

// Run 启动API服务器
func (s *Server) Run() error {
	addr := fmt.Sprintf("%s:%d", s.Options.Host, s.Options.Port)
	log.Info("启动API服务器", "address", addr)

	// 输出配置信息
	active, waiting, max, queue, _ := getConcurrencyStats()
	log.Info("服务器并发设置",
		"active", active,
		"waiting", waiting,
		"max_concurrent", max,
		"queue_size", queue,
	)

	// 创建HTTP服务器
	server := &http.Server{
		Addr:         addr,
		Handler:      s.Router,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	return server.ListenAndServe()
}
