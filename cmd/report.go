package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cyberspacesec/go-snir/pkg/log"
)

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "报告相关命令",
	Long:  "管理和查看扫描报告的相关命令",
}

func init() {
	// 添加report命令到根命令
	rootCmd.AddCommand(reportCmd)

	// 添加报告相关选项
	reportCmd.PersistentFlags().StringVar(&opts.Report.OutputPath, "output-path", "reports", "报告输出路径")
	reportCmd.PersistentFlags().StringVar(&opts.Report.Format, "format", "html", "报告格式 (html, json, csv)")
	reportCmd.PersistentFlags().StringVar(&opts.Report.Host, "host", "127.0.0.1", "Web服务器主机地址")
	reportCmd.PersistentFlags().IntVar(&opts.Report.Port, "port", 8080, "Web服务器端口")

	// 添加serve子命令
	reportCmd.AddCommand(serveCmd)

	log.Debug("已注册report命令")
}