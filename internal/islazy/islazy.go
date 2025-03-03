package islazy

import (
	"os"
	"path/filepath"
	"strings"
)

// CreateDir creates a directory if it does not exist
func CreateDir(path string) (string, error) {
	// 如果路径为空，使用默认路径
	if path == "" {
		path = "./screenshots"
	}

	// 获取绝对路径
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	// 检查目录是否存在，如果不存在则创建
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		err = os.MkdirAll(absPath, 0755)
		if err != nil {
			return "", err
		}
	}

	return absPath, nil
}

// FileExists checks if a file exists and is not a directory
func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// DirExists checks if a directory exists
func DirExists(dirname string) bool {
	info, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return false
	}
	return info.IsDir()
}

// SliceHasStr checks if a string slice contains a specific string
func SliceHasStr(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// SanitizeFilename sanitizes a string to be used as a filename
func SanitizeFilename(filename string) string {
	// 替换不安全的字符
	unsafe := []string{"\\", "/", ":", "*", "?", "\"", "<", ">", "|", "%"}
	result := filename

	for _, char := range unsafe {
		result = strings.ReplaceAll(result, char, "_")
	}

	// 移除前导和尾随空格
	result = strings.TrimSpace(result)

	// 确保文件名不为空
	if result == "" {
		result = "unnamed"
	}

	return result
}