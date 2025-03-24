package api

import (
	"github.com/cyberspacesec/go-snir/pkg/models"
)

// Write 实现 runner.Writer 接口
func (w *MemoryWriter) Write(result *models.Result) error {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.Results = append(w.Results, result)
	return nil
}

// Close 实现 runner.Writer 接口
func (w *MemoryWriter) Close() error {
	return nil
}
