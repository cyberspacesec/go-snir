package api

import (
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/cyberspacesec/go-snir/pkg/log"
)

// APIKeyMiddleware 验证API请求中的API密钥
func (s *Server) APIKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")
		if apiKey == "" {
			// 如果请求中没有API密钥，检查查询参数
			apiKey = r.URL.Query().Get("api_key")
		}

		// 验证API密钥
		if apiKey != s.Options.APIKey {
			SendJSONResponse(w, http.StatusUnauthorized, APIResponse{
				Success: false,
				Error:   "无效的API密钥",
			})
			return
		}

		// API密钥验证通过，调用下一个处理器
		next.ServeHTTP(w, r)
	})
}

// CreateAuthMiddleware 创建验证中间件函数
func (s *Server) CreateAuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// 跳过静态文件和根路径的验证
			if strings.HasPrefix(r.URL.Path, "/screenshots/") || r.URL.Path == "/" {
				next.ServeHTTP(w, r)
				return
			}

			// 验证API密钥
			apiKey := r.Header.Get("X-API-Key")
			if apiKey == "" {
				apiKey = r.URL.Query().Get("api_key")
			}

			if apiKey != s.Options.APIKey {
				SendJSONResponse(w, http.StatusUnauthorized, APIResponse{
					Success: false,
					Error:   "无效的API密钥",
				})
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

// ConcurrencyLimitMiddleware 限制并发请求数量
func (s *Server) ConcurrencyLimitMiddleware(next http.Handler) http.Handler {
	// 确保启用了并发限制
	if s.concurrencyLimit == nil {
		return next
	}

	limiter, ok := s.concurrencyLimit.(*ConcurrencyLimiter)
	if !ok {
		return next
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 跳过对状态检查和静态资源的限制
		if r.URL.Path == "/health" || r.URL.Path == "/stats" ||
			r.URL.Path == "/" || r.URL.Path == "/favicon.ico" ||
			strings.HasPrefix(r.URL.Path, "/screenshots/") {
			next.ServeHTTP(w, r)
			return
		}

		// 设置超时上下文
		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		// 尝试获取许可
		err := limiter.Acquire(ctx)
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
		defer limiter.Release()

		// 继续处理请求
		next.ServeHTTP(w, r)
	})
}
