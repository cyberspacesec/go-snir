package cmd

import (
	"fmt"

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
