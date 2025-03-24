package api

import (
	"sync"
	"time"

	"github.com/cyberspacesec/go-snir/pkg/models"
	"github.com/gorilla/mux"
)

// APIResponse 表示API响应结构
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// InteractionAction 表示交互操作
type InteractionAction struct {
	Type        string `json:"type"`         // click, scroll, type, wait, hover
	Selector    string `json:"selector"`     // CSS选择器
	XPath       string `json:"xpath"`        // XPath
	Value       string `json:"value"`        // 用于输入的值或滚动距离
	WaitTime    int    `json:"wait_time"`    // 等待时间(毫秒)
	WaitVisible bool   `json:"wait_visible"` // 等待元素可见
}

// FormField 表示表单字段
type FormField struct {
	Selector string `json:"selector"` // CSS选择器
	XPath    string `json:"xpath"`    // XPath
	Value    string `json:"value"`    // 填充的值
	Type     string `json:"type"`     // input, select, checkbox, radio
}

// Form 表示表单配置
type Form struct {
	Fields          []FormField `json:"fields"`            // 表单字段
	SubmitSelector  string      `json:"submit_selector"`   // 提交按钮选择器
	SubmitXPath     string      `json:"submit_xpath"`      // 提交按钮XPath
	WaitAfterSubmit int         `json:"wait_after_submit"` // 提交后等待时间(毫秒)
}

// CustomCookie 表示自定义Cookie
type CustomCookie struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Domain   string `json:"domain,omitempty"`
	Path     string `json:"path,omitempty"`
	Secure   bool   `json:"secure,omitempty"`
	HttpOnly bool   `json:"http_only,omitempty"`
}

// BrowserFingerprint 表示浏览器指纹
type BrowserFingerprint struct {
	UserAgent       string            `json:"user_agent,omitempty"`
	AcceptLanguage  string            `json:"accept_language,omitempty"`
	Platform        string            `json:"platform,omitempty"`
	Plugins         []string          `json:"plugins,omitempty"`
	Vendor          string            `json:"vendor,omitempty"`
	WebGLVendor     string            `json:"webgl_vendor,omitempty"`
	WebGLRenderer   string            `json:"webgl_renderer,omitempty"`
	CustomHeaders   map[string]string `json:"custom_headers,omitempty"`
	DisableWebRTC   bool              `json:"disable_webrtc,omitempty"`
	SpoofScreenSize bool              `json:"spoof_screen_size,omitempty"`
	ScreenWidth     int               `json:"screen_width,omitempty"`
	ScreenHeight    int               `json:"screen_height,omitempty"`
}

// ScreenshotRequest 表示截图请求结构
type ScreenshotRequest struct {
	URL              string `json:"url"`
	HTTPS            bool   `json:"https,omitempty"`
	HTTP             bool   `json:"http,omitempty"`
	UserAgent        string `json:"user_agent,omitempty"`
	Proxy            string `json:"proxy,omitempty"`
	Timeout          int    `json:"timeout,omitempty"`
	Delay            int    `json:"delay,omitempty"`
	IgnoreCertErrors bool   `json:"ignore_cert_errors,omitempty"`

	// 高级浏览器控制
	JavaScript     string             `json:"javascript,omitempty"`      // 注入的JS代码
	JavaScriptFile string             `json:"javascript_file,omitempty"` // JS文件路径
	RunJSBefore    bool               `json:"run_js_before,omitempty"`   // 在页面加载前执行
	RunJSAfter     bool               `json:"run_js_after,omitempty"`    // 在页面加载后执行
	Fingerprint    BrowserFingerprint `json:"fingerprint,omitempty"`     // 浏览器指纹配置
	Cookies        []CustomCookie     `json:"cookies,omitempty"`         // 自定义Cookie

	// 高级元素选择和交互
	Selector        string              `json:"selector,omitempty"`          // CSS选择器
	XPath           string              `json:"xpath,omitempty"`             // XPath
	CaptureFullPage bool                `json:"capture_full_page,omitempty"` // 是否捕获整个页面
	Actions         []InteractionAction `json:"actions,omitempty"`           // 交互操作列表
	Form            Form                `json:"form,omitempty"`              // 表单配置
}

// BatchScreenshotRequest 表示批量截图请求结构
type BatchScreenshotRequest struct {
	URLs             []string `json:"urls"`
	HTTPS            bool     `json:"https,omitempty"`
	HTTP             bool     `json:"http,omitempty"`
	UserAgent        string   `json:"user_agent,omitempty"`
	Proxy            string   `json:"proxy,omitempty"`
	Timeout          int      `json:"timeout,omitempty"`
	Delay            int      `json:"delay,omitempty"`
	Threads          int      `json:"threads,omitempty"`
	IgnoreCertErrors bool     `json:"ignore_cert_errors,omitempty"`

	// 高级浏览器控制
	JavaScript     string             `json:"javascript,omitempty"`      // 注入的JS代码
	JavaScriptFile string             `json:"javascript_file,omitempty"` // JS文件路径
	RunJSBefore    bool               `json:"run_js_before,omitempty"`   // 在页面加载前执行
	RunJSAfter     bool               `json:"run_js_after,omitempty"`    // 在页面加载后执行
	Fingerprint    BrowserFingerprint `json:"fingerprint,omitempty"`     // 浏览器指纹配置
	Cookies        []CustomCookie     `json:"cookies,omitempty"`         // 自定义Cookie

	// 高级元素选择和交互
	Selector        string              `json:"selector,omitempty"`          // CSS选择器
	XPath           string              `json:"xpath,omitempty"`             // XPath
	CaptureFullPage bool                `json:"capture_full_page,omitempty"` // 是否捕获整个页面
	Actions         []InteractionAction `json:"actions,omitempty"`           // 交互操作列表
	Form            Form                `json:"form,omitempty"`              // 表单配置
}

// Options 包含API服务的配置选项
type Options struct {
	Port                  int
	Host                  string
	ScreenshotPath        string
	APIKey                string   // API密钥，用于鉴权
	EnableBlacklist       bool     // 是否启用URL黑名单
	DefaultBlacklist      bool     // 是否使用默认黑名单
	BlacklistPatterns     []string // 自定义黑名单规则
	BlacklistFile         string   // 黑名单文件路径
	MaxConcurrentRequests int      // 最大并发请求数
	RequestQueueSize      int      // 请求队列大小
}

// Server 表示API服务器
type Server struct {
	Options          Options
	Router           *mux.Router
	concurrencyLimit interface{}   // 并发限制器
	shutdownCh       chan struct{} // 关闭通道
	serverStartTime  time.Time     // 服务器启动时间
}

// MemoryWriter 内存写入器实现 runner.Writer 接口
type MemoryWriter struct {
	Results []*models.Result
	mu      sync.Mutex
}
