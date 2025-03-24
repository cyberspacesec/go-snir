package api

import (
	"encoding/json"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/cyberspacesec/go-snir/pkg/islazy"
)

// CreateScreenshotDir 创建截图目录
func CreateScreenshotDir(path string) (string, error) {
	return islazy.CreateDir(path)
}

// GetImageContentType 获取图像文件的内容类型
func GetImageContentType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".png":
		return "image/png"
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	default:
		return "application/octet-stream"
	}
}

// IsImageFile 检查文件是否为图像文件
func IsImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".png" || ext == ".jpg" || ext == ".jpeg" || ext == ".gif"
}

// SendJSONResponse 发送JSON响应
func SendJSONResponse(w http.ResponseWriter, statusCode int, response APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
