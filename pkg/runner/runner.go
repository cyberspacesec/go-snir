package runner

import (
	"context"
	"errors"
	"log/slog"
	"net/url"
	"os"
	"sync"

	"github.com/cyberspacesec/go-web-screenshot/internal/islazy"
	"github.com/cyberspacesec/go-web-screenshot/pkg/models"
)

// Runner is a runner that probes web targets using a driver
type Runner struct {
	Driver Driver

	// options for the Runner to consider
	options Options
	// writers are the result writers to use
	writers []Writer
	// log handler
	log *slog.Logger

	// Targets to scan.
	Targets chan string

	// in case we need to bail
	ctx    context.Context
	cancel context.CancelFunc
}

// Writer is the interface result writers will implement
var defaultRunner *Runner

type Writer interface {
	Write(result *models.Result) error
	Close() error
}

// NewRunner gets a new Runner ready for probing.
// It's up to the caller to call Close() on the runner
func NewRunner(logger *slog.Logger, driver Driver, opts Options, writers []Writer) (*Runner, error) {
	if !opts.Scan.ScreenshotSkipSave {
		screenshotPath, err := islazy.CreateDir(opts.Scan.ScreenshotPath)
		if err != nil {
			return nil, err
		}
		opts.Scan.ScreenshotPath = screenshotPath
		logger.Debug("最终截图路径", "screenshot-path", opts.Scan.ScreenshotPath)
	} else {
		logger.Debug("不保存截图到磁盘")
	}

	// 检查截图格式
	if !islazy.SliceHasStr([]string{"jpeg", "png"}, opts.Scan.ScreenshotFormat) {
		return nil, errors.New("无效的截图格式")
	}

	// 包含JavaScript代码的文件
	// 读取文件内容并设置到Scan.JavaScript
	if opts.Scan.JavaScriptFile != "" {
		javascript, err := os.ReadFile(opts.Scan.JavaScriptFile)
		if err != nil {
			return nil, err
		}

		opts.Scan.JavaScript = string(javascript)
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &Runner{
		Driver:   driver,
		options:  opts,
		writers:  writers,
		Targets:  make(chan string),
		log:      logger,
		ctx:      ctx,
		cancel:   cancel,
	}, nil
}

// runWriters takes a result and passes it to writers
func (run *Runner) runWriters(result *models.Result) error {
	for _, writer := range run.writers {
		if err := writer.Write(result); err != nil {
			return err
		}
	}

	return nil
}

// checkUrl ensures a url is valid
func (run *Runner) checkUrl(target string) error {
	_, err := url.Parse(target)
	return err
}

// Run starts the runner, processing targets from the Targets channel
// Screenshot 执行单次截图操作
func Screenshot(target string) (*models.Result, error) {
	return defaultRunner.Driver.Witness(target, defaultRunner)
}

func (run *Runner) Run() error {
	var wg sync.WaitGroup

	// 创建工作线程池
	for i := 0; i < run.options.Scan.Threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				select {
				case target, ok := <-run.Targets:
					if !ok {
						return
					}

					if err := run.checkUrl(target); err != nil {
						run.log.Error("无效的URL", "url", target, "error", err)
						continue
					}

					result, err := run.Driver.Witness(target, run)
					if err != nil {
						run.log.Error("截图失败", "url", target, "error", err)
						continue
					}

					if err := run.runWriters(result); err != nil {
						run.log.Error("写入结果失败", "url", target, "error", err)
					}
				case <-run.ctx.Done():
					return
				}
			}
		}()
	}

	wg.Wait()
	return nil
}

// Close closes the runner and all writers
func (run *Runner) Close() error {
	run.cancel()

	// 关闭所有写入器
	for _, writer := range run.writers {
		if err := writer.Close(); err != nil {
			run.log.Error("关闭写入器失败", "error", err)
		}
	}

	return nil
}