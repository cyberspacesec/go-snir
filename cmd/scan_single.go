package cmd

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cyberspacesec/go-web-screenshot/pkg/log"
	"github.com/cyberspacesec/go-web-screenshot/pkg/models"
	"github.com/cyberspacesec/go-web-screenshot/pkg/runner"
)

var singleCmd = &cobra.Command{
	Use:   "single [url]",
	Short: "扫描单个URL",
	Long:  "扫描单个URL并进行截图",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		target := args[0]

		// 确保URL格式正确
		if !strings.HasPrefix(target, "http://") && !strings.HasPrefix(target, "https://") {
			// 根据配置添加协议前缀
			if opts.Scan.HTTPS {
				target = "https://" + target
			} else if opts.Scan.HTTP {
				target = "http://" + target
			} else {
				return fmt.Errorf("未指定协议，且未启用HTTP或HTTPS选项")
			}
		}

		// 验证URL格式
		_, err := url.Parse(target)
		if err != nil {
			return fmt.Errorf("无效的URL: %v", err)
		}

		log.Info("开始扫描单个URL", "url", target)

		// 创建Chrome驱动
		driver, err := createDriver()
		if err != nil {
			return fmt.Errorf("创建浏览器驱动失败: %v", err)
		}
		defer driver.Close()

		// 创建结果写入器
		writers, err := createWriters()
		if err != nil {
			return fmt.Errorf("创建结果写入器失败: %v", err)
		}

		// 创建运行器
		r, err := runner.NewRunner(log.SlogHandler, driver, *opts, writers)
		if err != nil {
			return fmt.Errorf("创建运行器失败: %v", err)
		}
		defer r.Close()

		// 执行截图
		result, err := driver.Witness(target, r)
		if err != nil {
			return fmt.Errorf("截图失败: %v", err)
		}

		// 输出结果
		printResult(result)

		return nil
	},
}

// 创建浏览器驱动
func createDriver() (runner.Driver, error) {
	// 创建ChromeDP驱动
	driver, err := runner.NewChromeDP(opts)
	if err != nil {
		return nil, err
	}
	return driver, nil
}

// 创建结果写入器
func createWriters() ([]runner.Writer, error) {
	return runner.CreateWriters(opts)
}

// 打印扫描结果
func printResult(result *models.Result) {
	if result == nil {
		log.Error("无结果")
		return
	}

	log.Info("扫描完成", "url", result.URL)
	log.Info("截图保存路径", "path", result.Filename)
	log.Info("页面标题", "title", result.Title)
	log.Info("响应状态码", "code", result.ResponseCode)
}

func init() {
	scanCmd.AddCommand(singleCmd)
	log.Debug("已注册single命令")
}