package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/cyberspacesec/go-snir/pkg/islazy"
	"github.com/cyberspacesec/go-snir/pkg/log"
	"github.com/cyberspacesec/go-snir/pkg/models"
	"github.com/cyberspacesec/go-snir/pkg/runner"
	"github.com/gorilla/mux"
)

// HandleRoot 处理根路径请求，不需要认证
func (s *Server) HandleRoot(w http.ResponseWriter, r *http.Request) {
	SendJSONResponse(w, http.StatusOK, APIResponse{
		Success: true,
		Message: "Go Web Screenshot API",
		Data: map[string]interface{}{
			"version": "1.0.0",
			"paths": []string{
				"/screenshot - 截图单个URL (需要API密钥)",
				"/batch - 批量截图多个URL (需要API密钥)",
				"/screenshots_list - 列出所有截图 (需要API密钥)",
				"/get_screenshot/{filename} - 获取指定截图 (需要API密钥)",
				"/screenshots/ - 直接访问截图文件（无需认证）",
			},
			"auth_required": true,
			"auth_method":   "请在请求头中添加X-API-Key或URL参数中添加api_key",
			"api_key":       "当前API密钥: " + s.Options.APIKey,
		},
	})
}

// HandleScreenshot 处理单个URL截图请求
func (s *Server) HandleScreenshot(w http.ResponseWriter, r *http.Request) {
	var req ScreenshotRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendJSONResponse(w, http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "无效的请求体: " + err.Error(),
		})
		return
	}

	if req.URL == "" {
		SendJSONResponse(w, http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "URL不能为空",
		})
		return
	}

	// 确保URL格式正确
	if !strings.HasPrefix(req.URL, "http://") && !strings.HasPrefix(req.URL, "https://") {
		if req.HTTPS {
			req.URL = "https://" + req.URL
		} else if req.HTTP {
			req.URL = "http://" + req.URL
		} else {
			// 默认使用HTTPS
			req.URL = "https://" + req.URL
		}
	}

	// 准备Chrome选项
	opts := runner.Options{}

	// 基本选项
	opts.Chrome.Path = ""
	opts.Chrome.UserAgent = req.UserAgent
	opts.Chrome.Proxy = req.Proxy
	opts.Chrome.Timeout = req.Timeout
	opts.Chrome.Delay = req.Delay
	opts.Chrome.WindowX = 1280
	opts.Chrome.WindowY = 800
	opts.Chrome.Headless = true
	opts.Chrome.IgnoreCertErrors = req.IgnoreCertErrors

	// 截图选项
	opts.Scan.ScreenshotPath = s.Options.ScreenshotPath
	opts.Scan.ScreenshotFormat = "png"
	opts.Scan.ScreenshotQuality = 90
	opts.Scan.ScreenshotSkipSave = false
	opts.Scan.HTTP = req.HTTP
	opts.Scan.HTTPS = req.HTTPS

	// 添加服务器级别的黑名单配置
	opts.Scan.EnableBlacklist = s.Options.EnableBlacklist
	opts.Scan.DefaultBlacklist = s.Options.DefaultBlacklist
	opts.Scan.BlacklistPatterns = s.Options.BlacklistPatterns
	opts.Scan.BlacklistFile = s.Options.BlacklistFile

	// 高级浏览器控制选项
	if req.Fingerprint.UserAgent != "" {
		opts.Chrome.UserAgent = req.Fingerprint.UserAgent
	}
	opts.Chrome.AcceptLanguage = req.Fingerprint.AcceptLanguage
	opts.Chrome.Platform = req.Fingerprint.Platform
	opts.Chrome.Vendor = req.Fingerprint.Vendor
	opts.Chrome.Plugins = req.Fingerprint.Plugins
	opts.Chrome.WebGLVendor = req.Fingerprint.WebGLVendor
	opts.Chrome.WebGLRenderer = req.Fingerprint.WebGLRenderer
	opts.Chrome.CustomHeaders = req.Fingerprint.CustomHeaders
	opts.Chrome.DisableWebRTC = req.Fingerprint.DisableWebRTC
	opts.Chrome.SpoofScreenSize = req.Fingerprint.SpoofScreenSize

	// 如果请求指定了屏幕尺寸，则使用请求中的值
	if req.Fingerprint.ScreenWidth > 0 {
		opts.Chrome.ScreenWidth = req.Fingerprint.ScreenWidth
	}
	if req.Fingerprint.ScreenHeight > 0 {
		opts.Chrome.ScreenHeight = req.Fingerprint.ScreenHeight
	}

	// JavaScript选项
	opts.Scan.JavaScript = req.JavaScript
	opts.Scan.JavaScriptFile = req.JavaScriptFile
	opts.Scan.RunJSBefore = req.RunJSBefore
	opts.Scan.RunJSAfter = req.RunJSAfter

	// Cookie管理
	if len(req.Cookies) > 0 {
		for _, cookie := range req.Cookies {
			opts.Scan.Cookies = append(opts.Scan.Cookies, runner.CustomCookie{
				Name:     cookie.Name,
				Value:    cookie.Value,
				Domain:   cookie.Domain,
				Path:     cookie.Path,
				Secure:   cookie.Secure,
				HttpOnly: cookie.HttpOnly,
			})
		}
	}

	// 高级元素选择和交互
	opts.Scan.Selector = req.Selector
	opts.Scan.XPath = req.XPath
	opts.Scan.CaptureFullPage = req.CaptureFullPage

	// 交互操作
	if len(req.Actions) > 0 {
		for _, action := range req.Actions {
			opts.Scan.Actions = append(opts.Scan.Actions, runner.InteractionAction{
				Type:        action.Type,
				Selector:    action.Selector,
				XPath:       action.XPath,
				Value:       action.Value,
				WaitTime:    action.WaitTime,
				WaitVisible: action.WaitVisible,
			})
		}
	}

	// 表单填充
	if len(req.Form.Fields) > 0 {
		formFields := []runner.FormField{}
		for _, field := range req.Form.Fields {
			formFields = append(formFields, runner.FormField{
				Selector: field.Selector,
				XPath:    field.XPath,
				Value:    field.Value,
				Type:     field.Type,
			})
		}
		opts.Scan.Form = runner.Form{
			Fields:          formFields,
			SubmitSelector:  req.Form.SubmitSelector,
			SubmitXPath:     req.Form.SubmitXPath,
			WaitAfterSubmit: req.Form.WaitAfterSubmit,
		}
	}

	// 首先创建黑名单实例并检查URL是否在黑名单中
	blacklist, err := runner.NewURLBlacklist(&opts)
	if err != nil {
		SendJSONResponse(w, http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "创建URL黑名单失败: " + err.Error(),
		})
		return
	}

	// 在调用Witness之前首先检查URL是否在黑名单中
	if isBlacklisted, reason := blacklist.IsBlacklisted(req.URL); isBlacklisted {
		log.Warn("尝试访问黑名单URL", "url", req.URL, "reason", reason)
		SendJSONResponse(w, http.StatusForbidden, APIResponse{
			Success: false,
			Error:   "URL在黑名单中: " + reason,
		})
		return
	}

	// 创建Chrome驱动
	driver, err := runner.NewChromeDP(&opts)
	if err != nil {
		SendJSONResponse(w, http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "创建浏览器驱动失败: " + err.Error(),
		})
		return
	}
	defer driver.Close()

	runnerInstance, err := runner.NewRunner(log.GetLogger(), driver, opts, nil)
	if err != nil {
		SendJSONResponse(w, http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "创建截图运行器失败: " + err.Error(),
		})
		return
	}
	defer runnerInstance.Close()

	result, err := driver.Witness(req.URL, runnerInstance)
	if err != nil {
		SendJSONResponse(w, http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "截图失败: " + err.Error(),
		})
		return
	}

	// 如果结果标记为失败，返回失败信息
	if result.Failed {
		SendJSONResponse(w, http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   result.FailedReason,
		})
		return
	}

	// 返回成功结果
	SendJSONResponse(w, http.StatusOK, APIResponse{
		Success: true,
		Message: "截图成功",
		Data:    result,
	})
}

// HandleBatchScreenshot 处理批量URL截图请求
func (s *Server) HandleBatchScreenshot(w http.ResponseWriter, r *http.Request) {
	var req BatchScreenshotRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		SendJSONResponse(w, http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "无效的请求体: " + err.Error(),
		})
		return
	}

	if len(req.URLs) == 0 {
		SendJSONResponse(w, http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "URL列表不能为空",
		})
		return
	}

	// 使用默认线程数
	if req.Threads <= 0 {
		req.Threads = 2
	}

	// 准备Chrome选项
	opts := runner.Options{}

	// 基本选项
	opts.Chrome.Path = ""
	opts.Chrome.UserAgent = req.UserAgent
	opts.Chrome.Proxy = req.Proxy
	opts.Chrome.Timeout = req.Timeout
	opts.Chrome.Delay = req.Delay
	opts.Chrome.WindowX = 1280
	opts.Chrome.WindowY = 800
	opts.Chrome.Headless = true
	opts.Chrome.IgnoreCertErrors = req.IgnoreCertErrors

	// 截图选项
	opts.Scan.ScreenshotPath = s.Options.ScreenshotPath
	opts.Scan.ScreenshotFormat = "png"
	opts.Scan.ScreenshotQuality = 90
	opts.Scan.ScreenshotSkipSave = false
	opts.Scan.HTTP = req.HTTP
	opts.Scan.HTTPS = req.HTTPS
	opts.Scan.Threads = req.Threads

	// 添加服务器级别的黑名单配置
	opts.Scan.EnableBlacklist = s.Options.EnableBlacklist
	opts.Scan.DefaultBlacklist = s.Options.DefaultBlacklist
	opts.Scan.BlacklistPatterns = s.Options.BlacklistPatterns
	opts.Scan.BlacklistFile = s.Options.BlacklistFile

	// 高级浏览器控制选项
	if req.Fingerprint.UserAgent != "" {
		opts.Chrome.UserAgent = req.Fingerprint.UserAgent
	}
	opts.Chrome.AcceptLanguage = req.Fingerprint.AcceptLanguage
	opts.Chrome.Platform = req.Fingerprint.Platform
	opts.Chrome.Vendor = req.Fingerprint.Vendor
	opts.Chrome.Plugins = req.Fingerprint.Plugins
	opts.Chrome.WebGLVendor = req.Fingerprint.WebGLVendor
	opts.Chrome.WebGLRenderer = req.Fingerprint.WebGLRenderer
	opts.Chrome.CustomHeaders = req.Fingerprint.CustomHeaders
	opts.Chrome.DisableWebRTC = req.Fingerprint.DisableWebRTC
	opts.Chrome.SpoofScreenSize = req.Fingerprint.SpoofScreenSize

	// 如果请求指定了屏幕尺寸，则使用请求中的值
	if req.Fingerprint.ScreenWidth > 0 {
		opts.Chrome.ScreenWidth = req.Fingerprint.ScreenWidth
	}
	if req.Fingerprint.ScreenHeight > 0 {
		opts.Chrome.ScreenHeight = req.Fingerprint.ScreenHeight
	}

	// JavaScript选项
	opts.Scan.JavaScript = req.JavaScript
	opts.Scan.JavaScriptFile = req.JavaScriptFile
	opts.Scan.RunJSBefore = req.RunJSBefore
	opts.Scan.RunJSAfter = req.RunJSAfter

	// Cookie管理
	if len(req.Cookies) > 0 {
		for _, cookie := range req.Cookies {
			opts.Scan.Cookies = append(opts.Scan.Cookies, runner.CustomCookie{
				Name:     cookie.Name,
				Value:    cookie.Value,
				Domain:   cookie.Domain,
				Path:     cookie.Path,
				Secure:   cookie.Secure,
				HttpOnly: cookie.HttpOnly,
			})
		}
	}

	// 高级元素选择和交互
	opts.Scan.Selector = req.Selector
	opts.Scan.XPath = req.XPath
	opts.Scan.CaptureFullPage = req.CaptureFullPage

	// 交互操作
	if len(req.Actions) > 0 {
		for _, action := range req.Actions {
			opts.Scan.Actions = append(opts.Scan.Actions, runner.InteractionAction{
				Type:        action.Type,
				Selector:    action.Selector,
				XPath:       action.XPath,
				Value:       action.Value,
				WaitTime:    action.WaitTime,
				WaitVisible: action.WaitVisible,
			})
		}
	}

	// 表单填充
	if len(req.Form.Fields) > 0 {
		formFields := []runner.FormField{}
		for _, field := range req.Form.Fields {
			formFields = append(formFields, runner.FormField{
				Selector: field.Selector,
				XPath:    field.XPath,
				Value:    field.Value,
				Type:     field.Type,
			})
		}
		opts.Scan.Form = runner.Form{
			Fields:          formFields,
			SubmitSelector:  req.Form.SubmitSelector,
			SubmitXPath:     req.Form.SubmitXPath,
			WaitAfterSubmit: req.Form.WaitAfterSubmit,
		}
	}

	// 创建黑名单检查器
	blacklist, err := runner.NewURLBlacklist(&opts)
	if err != nil {
		SendJSONResponse(w, http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "创建URL黑名单失败: " + err.Error(),
		})
		return
	}

	// 预检查所有URL，过滤掉黑名单中的URL
	var filteredURLs []string
	var blacklistedURLs []map[string]string

	for _, urlStr := range req.URLs {
		// 确保URL格式正确
		if !strings.HasPrefix(urlStr, "http://") && !strings.HasPrefix(urlStr, "https://") {
			if req.HTTPS {
				urlStr = "https://" + urlStr
			} else if req.HTTP {
				urlStr = "http://" + urlStr
			} else {
				// 默认使用HTTPS
				urlStr = "https://" + urlStr
			}
		}

		// 检查是否在黑名单中
		if isBlacklisted, reason := blacklist.IsBlacklisted(urlStr); isBlacklisted {
			log.Warn("批量扫描中跳过黑名单URL", "url", urlStr, "reason", reason)
			blacklistedURLs = append(blacklistedURLs, map[string]string{
				"url":    urlStr,
				"reason": reason,
			})
		} else {
			filteredURLs = append(filteredURLs, urlStr)
		}
	}

	// 如果所有URL都在黑名单中，直接返回错误
	if len(filteredURLs) == 0 {
		SendJSONResponse(w, http.StatusForbidden, APIResponse{
			Success: false,
			Error:   "所有请求的URL都在黑名单中",
			Data:    blacklistedURLs,
		})
		return
	}

	// 创建Chrome驱动
	driver, err := runner.NewChromeDP(&opts)
	if err != nil {
		SendJSONResponse(w, http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "创建浏览器驱动失败: " + err.Error(),
		})
		return
	}
	defer driver.Close()

	// 创建内存写入器
	memWriter := &MemoryWriter{
		Results: []*models.Result{},
	}

	runnerInstance, err := runner.NewRunner(log.GetLogger(), driver, opts, []runner.Writer{memWriter})
	if err != nil {
		SendJSONResponse(w, http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "创建截图运行器失败: " + err.Error(),
		})
		return
	}

	// 启动处理线程
	go func() {
		err := runnerInstance.Run()
		if err != nil {
			log.Error("批量截图失败", "error", err)
		}
		runnerInstance.Close()
	}()

	// 添加URL到处理队列
	for _, urlStr := range filteredURLs {
		runnerInstance.Targets <- urlStr
	}
	close(runnerInstance.Targets)

	// 等待所有任务完成
	time.Sleep(1 * time.Second)

	// 返回成功结果
	SendJSONResponse(w, http.StatusOK, APIResponse{
		Success: true,
		Message: fmt.Sprintf("已提交%d个URL进行截图", len(filteredURLs)),
		Data: map[string]interface{}{
			"task_id":          time.Now().Unix(),
			"filtered_urls":    len(filteredURLs),
			"blacklisted_urls": blacklistedURLs,
		},
	})
}

// HandleGetScreenshot 处理获取截图请求
func (s *Server) HandleGetScreenshot(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]

	if filename == "" {
		SendJSONResponse(w, http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "文件名不能为空",
		})
		return
	}

	// 防止目录遍历攻击
	cleanFilename := filepath.Clean(filename)
	if strings.Contains(cleanFilename, "..") {
		SendJSONResponse(w, http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "无效的文件名",
		})
		return
	}

	// 构建文件路径
	filePath := filepath.Join(s.Options.ScreenshotPath, cleanFilename)

	// 检查文件是否存在
	if !islazy.FileExists(filePath) {
		SendJSONResponse(w, http.StatusNotFound, APIResponse{
			Success: false,
			Error:   "文件不存在",
		})
		return
	}

	// 获取文件内容类型
	contentType := GetImageContentType(filePath)

	// 读取文件
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		SendJSONResponse(w, http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "读取文件失败: " + err.Error(),
		})
		return
	}

	// 设置响应头
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", cleanFilename))
	w.Write(data)
}

// HandleListScreenshots 处理列出截图请求
func (s *Server) HandleListScreenshots(w http.ResponseWriter, r *http.Request) {
	// 检查截图目录是否存在
	if !islazy.DirExists(s.Options.ScreenshotPath) {
		SendJSONResponse(w, http.StatusOK, APIResponse{
			Success: true,
			Data:    []string{},
		})
		return
	}

	var screenshots []map[string]interface{}

	// 遍历截图目录
	err := filepath.Walk(s.Options.ScreenshotPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 只处理文件
		if !info.IsDir() && IsImageFile(info.Name()) {
			relPath, err := filepath.Rel(s.Options.ScreenshotPath, path)
			if err != nil {
				relPath = path
			}

			screenshots = append(screenshots, map[string]interface{}{
				"filename": info.Name(),
				"path":     relPath,
				"size":     info.Size(),
				"time":     info.ModTime().Format(time.RFC3339),
				"url":      fmt.Sprintf("/screenshots/%s", relPath),
			})
		}
		return nil
	})

	if err != nil {
		SendJSONResponse(w, http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "读取截图目录失败: " + err.Error(),
		})
		return
	}

	SendJSONResponse(w, http.StatusOK, APIResponse{
		Success: true,
		Data:    screenshots,
	})
}
