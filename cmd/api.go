package cmd

import (
	"crypto/rand"
	"encoding/hex"

	"github.com/spf13/cobra"

	"github.com/cyberspacesec/go-snir/pkg/api"
	"github.com/cyberspacesec/go-snir/pkg/log"
)

// 生成随机API密钥
func generateRandomAPIKey(length int) string {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		log.Error("生成API密钥失败", "error", err)
		return "go-snir-random-api-key"
	}
	return hex.EncodeToString(bytes)
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "启动API服务",
	Long:  "启动一个RESTful API服务，用于进行网页截图和信息收集",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 如果未指定API密钥，则生成一个随机密钥
		if opts.API.APIKey == "" {
			opts.API.APIKey = generateRandomAPIKey(32)
			log.Info("已生成随机API密钥", "api_key", opts.API.APIKey)
		}

		// 创建API服务配置
		apiOptions := api.Options{
			Port:                  opts.API.Port,
			Host:                  opts.API.Host,
			ScreenshotPath:        opts.Scan.ScreenshotPath,
			APIKey:                opts.API.APIKey,
			EnableBlacklist:       opts.Scan.EnableBlacklist,
			DefaultBlacklist:      opts.Scan.DefaultBlacklist,
			BlacklistPatterns:     opts.Scan.BlacklistPatterns,
			BlacklistFile:         opts.Scan.BlacklistFile,
			MaxConcurrentRequests: opts.API.MaxConcurrent,
			RequestQueueSize:      opts.API.QueueSize,
		}

		// 创建API服务
		server := api.NewServer(apiOptions)

		// 设置路由
		server.SetupRoutes()

		// 启动服务
		log.Info("启动API服务", "host", opts.API.Host, "port", opts.API.Port)
		return server.Run()
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	// 添加API相关选项
	apiCmd.Flags().StringVar(&opts.API.Host, "host", "0.0.0.0", "API服务监听地址")
	apiCmd.Flags().IntVar(&opts.API.Port, "port", 8080, "API服务监听端口")
	apiCmd.Flags().StringVar(&opts.API.APIKey, "api-key", "", "API密钥，用于API鉴权，如不指定则自动生成")

	// 添加黑名单相关选项
	apiCmd.Flags().BoolVar(&opts.Scan.EnableBlacklist, "enable-blacklist", true, "启用URL黑名单检查")
	apiCmd.Flags().BoolVar(&opts.Scan.DefaultBlacklist, "default-blacklist", true, "使用默认黑名单规则")
	apiCmd.Flags().StringSliceVar(&opts.Scan.BlacklistPatterns, "blacklist-pattern", []string{}, "添加自定义黑名单规则 (可多次使用)")
	apiCmd.Flags().StringVar(&opts.Scan.BlacklistFile, "blacklist-file", "", "黑名单规则文件路径")

	// 添加并发控制相关选项
	apiCmd.Flags().IntVar(&opts.API.MaxConcurrent, "max-concurrent", 10, "最大并发请求数")
	apiCmd.Flags().IntVar(&opts.API.QueueSize, "queue-size", 100, "请求队列大小")

	log.Debug("已注册api命令")
}
