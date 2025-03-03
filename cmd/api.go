package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	"github.com/cyberspacesec/go-web-screenshot/pkg/log"
	"github.com/cyberspacesec/go-web-screenshot/pkg/models"
	"github.com/cyberspacesec/go-web-screenshot/pkg/runner"
)

// API响应结构
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// 截图请求结构
type ScreenshotRequest struct {
	URL      string `json:"url"`
	HTTPS    bool   `json:"https,omitempty"`
	HTTP     bool   `json:"http,omitempty"`
	UserAgent string `json:"user_agent,omitempty"`
	Proxy    string `json:"proxy,omitempty"`
	Timeout  int    `json:"timeout,omitempty"`
	Delay    int    `json:"delay,omitempty"`
}

// 批量截图请求结构
type BatchScreenshotRequest struct {
	URLs     []string `json:"urls"`
	HTTPS    bool     `json:"https,omitempty"`
	HTTP     bool     `json:"http,omitempty"`
	UserAgent string   `json:"user_agent,omitempty"`
	Proxy    string   `json:"proxy,omitempty"`
	Timeout  int      `json:"timeout,omitempty"`
	Delay    int      `json:"delay,omitempty"`
	Threads  int      `json:"threads,omitempty"`
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "启动HTTP API服务",
	Long:  "启动一个HTTP API服务，提供网页截图功能的RESTful接口",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 确保截图目录存在
		screenshotPath, err := createScreenshotDir()
		if err != nil {
			return fmt.Errorf("创建截图目录失败: %v", err)
		}

		// 创建路由器
		router := mux.NewRouter()

		// API路由
		apiRouter := router.PathPrefix("/api").Subrouter()
		apiRouter.HandleFunc("/screenshot", handleScreenshot).Methods("POST")
		apiRouter.HandleFunc("/batch", handleBatchScreenshot).Methods("POST")
		apiRouter.HandleFunc("/screenshots", handleListScreenshots).Methods("GET")
		apiRouter.HandleFunc("/screenshots/{filename}", handleGetScreenshot).Methods("GET")

		// 静态文件服务
		router.PathPrefix("/screenshots/").Handler(
			http.StripPrefix("/screenshots/", http.FileServer(http.Dir(screenshotPath))),
		)

		// 启动服务器
		addr := fmt.Sprintf("%s:%d", opts.API.Host, opts.API.Port)
		log.Info("启动API服务器", "address", addr)
		log.Info(fmt.Sprintf("API服务器已启动，访问 http://%s/api", addr))

		return http.ListenAndServe(addr, router)
	},
}

// 创建截图目录
func createScreenshotDir() (string, error) {
	// 确保截图目录存在
	screenshotPath := opts.Scan.ScreenshotPath
	if _, err := os.Stat(screenshotPath); os.IsNotExist(err) {
		err := os.MkdirAll(screenshotPath, 0755)
		if err != nil {
			return "", err
		}
	}
	return screenshotPath, nil
}

// 处理单个URL截图请求
func handleScreenshot(w http.ResponseWriter, r *http.Request) {
	// 解析请求
	var req ScreenshotRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendJSONResponse(w, http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "无效的请求格式: " + err.Error(),
		})
		return
	}

	// 验证URL
	if req.URL == "" {
		sendJSONResponse(w, http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "URL不能为空",
		})
		return
	}

	// 应用请求中的配置到选项
	applyScreenshotOptions(&req)

	// 确保URL格式正确
	target := req.URL
	if !strings.HasPrefix(target, "http://") && !strings.HasPrefix(target, "https://") {
		// 根据配置添加协议前缀
		if opts.Scan.HTTPS {
			target = "https://" + target
		} else if opts.Scan.HTTP {
			target = "http://" + target
		} else {
			sendJSONResponse(w, http.StatusBadRequest, APIResponse{
				Success: false,
				Error:   "未指定协议，且未启用HTTP或HTTPS选项",
			})
			return
		}
	}

	// 执行截图
	result, err := runner.Screenshot(target)
	if err != nil {
		sendJSONResponse(w, http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "截图失败: " + err.Error(),
		})
		return
	}

	// 发送成功响应
	sendJSONResponse(w, http.StatusOK, APIResponse{
		Success: true,
		Message: "截图成功",
		Data: map[string]interface{}{
			"url":      target,
			"filename": result.Filename,
			"path":     result.Path,
		},
	})
}

// 处理批量截图请求
func handleBatchScreenshot(w http.ResponseWriter, r *http.Request) {
	// 解析请求
	var req BatchScreenshotRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		sendJSONResponse(w, http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "无效的请求格式: " + err.Error(),
		})
		return
	}

	// 验证URLs
	if len(req.URLs) == 0 {
		sendJSONResponse(w, http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "URLs列表不能为空",
		})
		return
	}

	// 应用请求中的配置到选项
	applyBatchScreenshotOptions(&req)

	// 创建Chrome驱动
	driver, err := runner.NewChromeDP(opts)
	if err != nil {
		sendJSONResponse(w, http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "创建浏览器驱动失败: " + err.Error(),
		})
		return
	}
	defer driver.Close()

	// 创建结果写入器
	writers, err := runner.CreateWriters(opts)
	if err != nil {
		sendJSONResponse(w, http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "创建结果写入器失败: " + err.Error(),
		})
		return
	}

	// 创建运行器
	runnerInstance, err := runner.NewRunner(log.SlogHandler, driver, *opts, writers)
	if err != nil {
		sendJSONResponse(w, http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "创建运行器失败: " + err.Error(),
		})
		return
	}
	defer runnerInstance.Close()

	// 使用goroutine池并发处理URL
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, opts.Scan.Threads)
	resultsChan := make(chan *models.Result, opts.Scan.Threads)
	results := make([]*models.Result, 0)

	// 启动结果处理goroutine
	go func() {
		for result := range resultsChan {
			results = append(results, result)
		}
	}()

	// 处理每个URL
	for _, url := range req.URLs {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(target string) {
			defer wg.Done()
			defer func() { <-semaphore }()

			// 确保URL格式正确
			if !strings.HasPrefix(target, "http://") && !strings.HasPrefix(target, "https://") {
				// 尝试HTTPS
				if opts.Scan.HTTPS {
					target = "https://" + target
				} else if opts.Scan.HTTP {
					target = "http://" + target
				} else {
					log.Error("URL缺少协议前缀且未启用HTTP或HTTPS选项", "url", target)
					return
				}
			}

			log.Info("开始扫描URL", "url", target)

			// 执行截图
			result, err := driver.Witness(target, runnerInstance)
			if err != nil {
				log.Error("截图失败", "url", target, "error", err)
				return
			}

			// 发送结果到通道
			resultsChan <- result
		}(url)
	}

	// 等待所有goroutine完成
	wg.Wait()
	close(semaphore)
	close(resultsChan)

	// 构建响应数据
	responseData := make([]map[string]interface{}, 0)
	for _, result := range results {
		responseData = append(responseData, map[string]interface{}{
			"url":           result.URL,
			"title":         result.Title,
			"response_code": result.ResponseCode,
			"screenshot":    "/screenshots/" + filepath.Base(result.Filename),
			"probed_at":     result.ProbedAt,
		})
	}

	// 发送响应
	sendJSONResponse(w, http.StatusOK, APIResponse{
		Success: true,
		Message: fmt.Sprintf("成功处理 %d 个URL", len(responseData)),
		Data:    responseData,
	})
}

// 处理获取单个截图请求
func handleGetScreenshot(w http.ResponseWriter, r *http.Request) {
	// 获取文件名参数
	vars := mux.Vars(r)
	filename := vars["filename"]

	// 验证文件名
	if filename == "" {
		sendJSONResponse(w, http.StatusBadRequest, APIResponse{
			Success: false,
			Error:   "文件名不能为空",
		})
		return
	}

	// 构建文件路径
	filePath := filepath.Join(opts.Scan.ScreenshotPath, filename)

	// 检查文件
	// 检查文件是否存在
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		sendJSONResponse(w, http.StatusNotFound, APIResponse{
			Success: false,
			Error:   "找不到指定的截图文件",
		})
		return
	}

	// 读取文件内容
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		sendJSONResponse(w, http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "读取文件失败: " + err.Error(),
		})
		return
	}

	// 设置响应头
	w.Header().Set("Content-Type", "image/"+getImageContentType(filename))
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(fileContent)))

	// 发送文件内容
	if _, err := w.Write(fileContent); err != nil {
		log.Error("发送文件失败", "error", err)
	}
}

// 获取图片内容类型
func getImageContentType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".png":
		return "png"
	case ".jpg", ".jpeg":
		return "jpeg"
	default:
		return "png"
	}
}

// 处理获取截图列表请求
func handleListScreenshots(w http.ResponseWriter, r *http.Request) {
	// 获取截图目录中的所有文件
	files, err := ioutil.ReadDir(opts.Scan.ScreenshotPath)
	if err != nil {
		sendJSONResponse(w, http.StatusInternalServerError, APIResponse{
			Success: false,
			Error:   "读取截图目录失败: " + err.Error(),
		})
		return
	}

	// 构建截图列表
	screenshots := make([]map[string]interface{}, 0)
	for _, file := range files {
		// 只处理图片文件
		if !file.IsDir() && isImageFile(file.Name()) {
			screenshots = append(screenshots, map[string]interface{}{
				"filename":  file.Name(),
				"size":      file.Size(),
				"modified":  file.ModTime().Format(time.RFC3339),
				"url":       "/screenshots/" + file.Name(),
			})
		}
	}

	// 发送响应
	sendJSONResponse(w, http.StatusOK, APIResponse{
		Success: true,
		Message: fmt.Sprintf("找到 %d 个截图", len(screenshots)),
		Data:    screenshots,
	})
}

// 判断是否为图片文件
func isImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".png" || ext == ".jpg" || ext == ".jpeg"
}

// 发送JSON响应
func sendJSONResponse(w http.ResponseWriter, statusCode int, response APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}