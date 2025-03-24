package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cyberspacesec/go-snir/pkg/log"
	"github.com/cyberspacesec/go-snir/pkg/report"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "启动Web服务器查看结果",
	Long:  "启动一个Web服务器，用于查看截图和扫描结果",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 创建服务器配置
		serverOptions := report.ServerOptions{
			Host:           opts.Report.Host,
			Port:           opts.Report.Port,
			ScreenshotPath: opts.Scan.ScreenshotPath,
			ReportPath:     opts.Report.OutputPath,
		}

		// 创建服务器
		server := report.NewServer(serverOptions)

		// 启动服务器
		log.Info("启动Web服务器", "host", opts.Report.Host, "port", opts.Report.Port)
		return server.Run()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// 添加服务器选项
	serveCmd.Flags().StringVar(&opts.Report.Host, "host", "0.0.0.0", "Web服务器监听地址")
	serveCmd.Flags().IntVar(&opts.Report.Port, "port", 8080, "Web服务器监听端口")

	log.Debug("已注册serve命令")
}
