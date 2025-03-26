package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cyberspacesec/go-snir/pkg/log"
	"github.com/cyberspacesec/go-snir/pkg/scan"
)

var singleCmd = &cobra.Command{
	Use:   "single [url]",
	Short: log.Yellow("扫描单个URL"),
	Long:  log.Yellow("扫描单个URL并进行截图"),
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		target := args[0]

		// 创建扫描配置
		config := &scan.Config{
			Target:  target,
			Options: opts,
		}

		// 创建扫描器
		scanner, err := scan.NewScanner(config)
		if err != nil {
			return fmt.Errorf("创建扫描器失败: %v", err)
		}
		defer scanner.Close()

		// 执行扫描
		log.CommandTitle("扫描URL")
		log.Info("开始扫描", "url", log.Cyan(target))
		result, err := scanner.ScanSingle(target)
		if err != nil {
			// 美化错误消息
			errMsg := err.Error()

			// 处理常见的ChromeDP错误
			if strings.Contains(errMsg, "Could not find node with given id") {
				return fmt.Errorf("扫描过程中发生错误: 无法找到页面上的某个元素。这可能是因为:\n" +
					"1. 网站加载较慢，请尝试增加超时时间 (--timeout 选项)\n" +
					"2. 网站可能有反爬虫措施\n" +
					"3. 网站结构与预期不符\n" +
					"建议尝试增加延迟: --delay 3")
			} else if strings.Contains(errMsg, "timeout") {
				return fmt.Errorf("扫描超时: 无法在指定时间内完成页面加载。请尝试:\n" +
					"1. 增加超时时间: --timeout 60\n" +
					"2. 检查网络连接\n" +
					"3. 检查目标站点是否可访问")
			} else if strings.Contains(errMsg, "net::ERR_") {
				return fmt.Errorf("网络错误: 无法连接到目标网站。请检查:\n" +
					"1. 目标URL是否正确\n" +
					"2. 您的网络连接\n" +
					"3. 目标站点是否在线")
			}

			return fmt.Errorf("扫描失败: %v", err)
		}

		// 打印结果
		printResult(result)

		return nil
	},
}

// printResult 打印扫描结果
func printResult(result interface{}) {
	log.Success("扫描完成")
}

func init() {
	scanCmd.AddCommand(singleCmd)
	log.Debug(log.Green("已注册single命令"))
}
