package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cyberspacesec/go-snir/pkg/log"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "扫描并截图网站",
	Long:  "扫描指定的URL、文件或网段，并对网站进行截图和信息收集",
}

func init() {
	// 添加scan命令到根命令
	rootCmd.AddCommand(scanCmd)

	// 添加通用的截图选项
	scanCmd.PersistentFlags().StringVar(&opts.Scan.ScreenshotPath, "screenshot-path", "screenshots", "截图保存路径")
	scanCmd.PersistentFlags().StringVar(&opts.Scan.ScreenshotFormat, "screenshot-format", "png", "截图格式 (png或jpeg)")
	scanCmd.PersistentFlags().IntVar(&opts.Scan.ScreenshotQuality, "screenshot-quality", 90, "截图质量 (仅对jpeg格式有效)")
	scanCmd.PersistentFlags().BoolVar(&opts.Scan.ScreenshotSkipSave, "skip-screenshot", false, "跳过保存截图")
	scanCmd.PersistentFlags().BoolVar(&opts.Scan.SaveHTML, "save-html", false, "保存网页HTML内容")
	scanCmd.PersistentFlags().BoolVar(&opts.Scan.SaveHeaders, "save-headers", false, "保存HTTP响应头")
	scanCmd.PersistentFlags().BoolVar(&opts.Scan.SaveConsole, "save-console", false, "保存控制台日志")
	scanCmd.PersistentFlags().BoolVar(&opts.Scan.SaveCookies, "save-cookies", false, "保存Cookie")

	// Chrome相关选项
	scanCmd.PersistentFlags().StringVar(&opts.Chrome.Path, "chrome-path", "", "Chrome可执行文件路径")
	scanCmd.PersistentFlags().StringVar(&opts.Chrome.UserAgent, "user-agent", "", "自定义User-Agent")
	scanCmd.PersistentFlags().StringVar(&opts.Chrome.Proxy, "proxy", "", "代理服务器地址")
	scanCmd.PersistentFlags().IntVar(&opts.Chrome.Timeout, "timeout", 30, "页面加载超时时间(秒)")
	scanCmd.PersistentFlags().IntVar(&opts.Chrome.Delay, "delay", 0, "截图前等待时间(秒)")
	scanCmd.PersistentFlags().IntVar(&opts.Chrome.WindowX, "resolution-x", 1280, "窗口宽度")
	scanCmd.PersistentFlags().IntVar(&opts.Chrome.WindowY, "resolution-y", 800, "窗口高度")
	scanCmd.PersistentFlags().BoolVar(&opts.Chrome.Headless, "headless", true, "使用无头模式")

	// 扫描相关选项
	scanCmd.PersistentFlags().IntVar(&opts.Scan.Threads, "threads", 2, "并发线程数")
	scanCmd.PersistentFlags().BoolVar(&opts.Scan.HTTP, "http", true, "使用HTTP协议")
	scanCmd.PersistentFlags().BoolVar(&opts.Scan.HTTPS, "https", true, "使用HTTPS协议")
	scanCmd.PersistentFlags().IntVar(&opts.Scan.MaxRetries, "max-retries", 1, "最大重试次数")
	scanCmd.PersistentFlags().StringVar(&opts.Scan.JavaScript, "js", "", "要在页面上执行的JavaScript代码")
	scanCmd.PersistentFlags().StringVar(&opts.Scan.JavaScriptFile, "js-file", "", "包含JavaScript代码的文件路径")

	// 数据库相关选项
	scanCmd.PersistentFlags().BoolVar(&opts.DB.Enable, "db", false, "启用数据库存储")
	scanCmd.PersistentFlags().StringVar(&opts.DB.Path, "db-path", "go-web-screenshot.db", "数据库文件路径")

	// 输出相关选项
	scanCmd.PersistentFlags().BoolVar(&opts.Writer.Jsonl, "write-jsonl", false, "写入JSONL格式结果")
	scanCmd.PersistentFlags().StringVar(&opts.Writer.JsonlFile, "jsonl-file", "results.jsonl", "JSONL结果文件路径")
	scanCmd.PersistentFlags().BoolVar(&opts.Writer.Csv, "write-csv", false, "写入CSV格式结果")
	scanCmd.PersistentFlags().StringVar(&opts.Writer.CsvFile, "csv-file", "results.csv", "CSV结果文件路径")
	scanCmd.PersistentFlags().BoolVar(&opts.Writer.Stdout, "write-stdout", true, "输出结果到控制台")

	// 添加黑名单相关选项
	scanCmd.PersistentFlags().BoolVar(&opts.Scan.EnableBlacklist, "enable-blacklist", true, "启用URL黑名单检查")
	scanCmd.PersistentFlags().BoolVar(&opts.Scan.DefaultBlacklist, "default-blacklist", true, "使用默认黑名单规则")
	scanCmd.PersistentFlags().StringSliceVar(&opts.Scan.BlacklistPatterns, "blacklist-pattern", []string{}, "添加自定义黑名单规则 (可多次使用)")
	scanCmd.PersistentFlags().StringVar(&opts.Scan.BlacklistFile, "blacklist-file", "", "黑名单规则文件路径")

	log.Debug("已注册scan命令")
}
