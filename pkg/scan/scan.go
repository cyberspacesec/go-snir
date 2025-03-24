package scan

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/cyberspacesec/go-snir/pkg/log"
	"github.com/cyberspacesec/go-snir/pkg/models"
	"github.com/cyberspacesec/go-snir/pkg/runner"
)

// Config 表示扫描配置
type Config struct {
	Target     string   // 扫描目标
	TargetFile string   // 包含目标的文件
	Targets    []string // 目标列表
	Options    *runner.Options
}

// Scanner 表示扫描器
type Scanner struct {
	Config  *Config
	Driver  runner.Driver
	Writers []runner.Writer
	Runner  *runner.Runner
}

// NewScanner 创建一个新的扫描器
func NewScanner(config *Config) (*Scanner, error) {
	// 验证配置
	if config == nil || config.Options == nil {
		return nil, fmt.Errorf("扫描配置不能为空")
	}

	// 创建驱动
	driver, err := createDriver(config.Options)
	if err != nil {
		return nil, fmt.Errorf("创建浏览器驱动失败: %v", err)
	}

	// 创建结果写入器
	writers, err := createWriters(config.Options)
	if err != nil {
		return nil, fmt.Errorf("创建结果写入器失败: %v", err)
	}

	scanner := &Scanner{
		Config:  config,
		Driver:  driver,
		Writers: writers,
	}

	return scanner, nil
}

// createDriver 创建浏览器驱动
func createDriver(options *runner.Options) (runner.Driver, error) {
	// 使用完整的ChromeDP实现而不是简化版的ChromeDriver
	return runner.NewChromeDP(options)
}

// createWriters 创建结果写入器
func createWriters(options *runner.Options) ([]runner.Writer, error) {
	return runner.CreateWriters(options)
}

// ScanSingle 扫描单个URL
func (s *Scanner) ScanSingle(target string) (*models.Result, error) {
	// 确保URL格式正确
	if !strings.HasPrefix(target, "http://") && !strings.HasPrefix(target, "https://") {
		// 根据配置添加协议前缀
		if s.Config.Options.Scan.HTTPS {
			target = "https://" + target
		} else if s.Config.Options.Scan.HTTP {
			target = "http://" + target
		} else {
			// 默认使用HTTPS
			target = "https://" + target
		}
	}

	// 验证URL格式
	_, err := url.Parse(target)
	if err != nil {
		return nil, fmt.Errorf("无效的URL: %v", err)
	}

	log.Info("开始扫描单个URL", "url", target)

	// 创建Runner（如果尚未创建）
	if s.Runner == nil {
		runner, err := runner.NewRunner(log.GetLogger(), s.Driver, *s.Config.Options, s.Writers)
		if err != nil {
			return nil, fmt.Errorf("创建扫描运行器失败: %v", err)
		}
		s.Runner = runner
	}

	// 执行扫描
	result, err := s.Driver.Witness(target, s.Runner)
	if err != nil {
		return nil, fmt.Errorf("扫描失败: %v", err)
	}

	// 运行写入器
	for _, writer := range s.Writers {
		if err := writer.Write(result); err != nil {
			log.Error("写入结果失败", "error", err)
		}
	}

	return result, nil
}

// ScanMulti 扫描多个URL
func (s *Scanner) ScanMulti(targets []string) error {
	// 创建Runner（如果尚未创建）
	if s.Runner == nil {
		runner, err := runner.NewRunner(log.GetLogger(), s.Driver, *s.Config.Options, s.Writers)
		if err != nil {
			return fmt.Errorf("创建扫描运行器失败: %v", err)
		}
		s.Runner = runner
	}

	// 启动扫描
	go func() {
		for _, target := range targets {
			// 确保URL格式正确
			if !strings.HasPrefix(target, "http://") && !strings.HasPrefix(target, "https://") {
				// 根据配置添加协议前缀
				if s.Config.Options.Scan.HTTPS {
					target = "https://" + target
				} else if s.Config.Options.Scan.HTTP {
					target = "http://" + target
				} else {
					// 默认使用HTTPS
					target = "https://" + target
				}
			}
			s.Runner.Targets <- target
		}
		close(s.Runner.Targets)
	}()

	// 执行扫描
	return s.Runner.Run()
}

// Close 关闭扫描器
func (s *Scanner) Close() error {
	var err error
	// 关闭Runner
	if s.Runner != nil {
		err = s.Runner.Close()
	} else {
		// 关闭驱动
		s.Driver.Close()

		// 关闭写入器
		for _, writer := range s.Writers {
			if writerErr := writer.Close(); writerErr != nil {
				log.Error("关闭写入器失败", "error", writerErr)
				if err == nil {
					err = writerErr
				}
			}
		}
	}
	return err
}
