package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cyberspacesec/go-snir/pkg/log"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: log.Yellow("扫描并截图网站"),
	Long:  log.Yellow("扫描指定的URL、文件或网段，并对网站进行截图和信息收集"),
}

func init() {
	// 添加scan命令到根命令
	rootCmd.AddCommand(scanCmd)

	// 添加通用的截图选项
	scanCmd.PersistentFlags().StringVar(&opts.Scan.ScreenshotPath, "screenshot-path", "screenshots", log.Cyan("截图保存路径"))
	scanCmd.PersistentFlags().StringVar(&opts.Scan.ScreenshotFormat, "screenshot-format", "png", log.Cyan("截图格式 (png或jpeg)"))
	scanCmd.PersistentFlags().IntVar(&opts.Scan.ScreenshotQuality, "screenshot-quality", 90, log.Cyan("截图质量 (仅对jpeg格式有效)"))
	scanCmd.PersistentFlags().BoolVar(&opts.Scan.ScreenshotSkipSave, "skip-screenshot", false, log.Cyan("跳过保存截图"))
	scanCmd.PersistentFlags().BoolVar(&opts.Scan.SaveHTML, "save-html", false, log.Cyan("保存网页HTML内容"))
	scanCmd.PersistentFlags().BoolVar(&opts.Scan.SaveHeaders, "save-headers", false, log.Cyan("保存HTTP响应头"))
	scanCmd.PersistentFlags().BoolVar(&opts.Scan.SaveConsole, "save-console", false, log.Cyan("保存控制台日志"))
	scanCmd.PersistentFlags().BoolVar(&opts.Scan.SaveCookies, "save-cookies", false, log.Cyan("保存Cookie"))

	// Chrome相关选项
	scanCmd.PersistentFlags().StringVar(&opts.Chrome.Path, "chrome-path", "", log.Cyan("Chrome可执行文件路径"))
	scanCmd.PersistentFlags().StringVar(&opts.Chrome.UserAgent, "user-agent", "", log.Cyan("自定义User-Agent"))
	scanCmd.PersistentFlags().StringVar(&opts.Chrome.Proxy, "proxy", "", log.Cyan("代理服务器地址"))
	scanCmd.PersistentFlags().IntVar(&opts.Chrome.Timeout, "timeout", 30, log.Cyan("页面加载超时时间(秒)"))
	scanCmd.PersistentFlags().IntVar(&opts.Chrome.Delay, "delay", 0, log.Cyan("截图前等待时间(秒)"))
	scanCmd.PersistentFlags().IntVar(&opts.Chrome.WindowX, "resolution-x", 1280, log.Cyan("窗口宽度"))
	scanCmd.PersistentFlags().IntVar(&opts.Chrome.WindowY, "resolution-y", 800, log.Cyan("窗口高度"))
	scanCmd.PersistentFlags().BoolVar(&opts.Chrome.Headless, "headless", true, log.Cyan("使用无头模式"))

	// 扫描相关选项
	scanCmd.PersistentFlags().IntVar(&opts.Scan.Threads, "threads", 2, log.Cyan("并发线程数"))
	scanCmd.PersistentFlags().BoolVar(&opts.Scan.HTTP, "http", true, log.Cyan("使用HTTP协议"))
	scanCmd.PersistentFlags().BoolVar(&opts.Scan.HTTPS, "https", true, log.Cyan("使用HTTPS协议"))
	scanCmd.PersistentFlags().IntVar(&opts.Scan.MaxRetries, "max-retries", 1, log.Cyan("最大重试次数"))
	scanCmd.PersistentFlags().StringVar(&opts.Scan.JavaScript, "js", "", log.Cyan("要在页面上执行的JavaScript代码"))
	scanCmd.PersistentFlags().StringVar(&opts.Scan.JavaScriptFile, "js-file", "", log.Cyan("包含JavaScript代码的文件路径"))

	// 数据库相关选项
	scanCmd.PersistentFlags().BoolVar(&opts.DB.Enable, "db", false, log.Cyan("启用数据库存储"))
	scanCmd.PersistentFlags().StringVar(&opts.DB.Path, "db-path", "go-web-screenshot.db", log.Cyan("数据库文件路径"))

	// 输出相关选项
	scanCmd.PersistentFlags().BoolVar(&opts.Writer.Jsonl, "write-jsonl", false, log.Cyan("写入JSONL格式结果"))
	scanCmd.PersistentFlags().StringVar(&opts.Writer.JsonlFile, "jsonl-file", "results.jsonl", log.Cyan("JSONL结果文件路径"))
	scanCmd.PersistentFlags().BoolVar(&opts.Writer.Csv, "write-csv", false, log.Cyan("写入CSV格式结果"))
	scanCmd.PersistentFlags().StringVar(&opts.Writer.CsvFile, "csv-file", "results.csv", log.Cyan("CSV结果文件路径"))
	scanCmd.PersistentFlags().BoolVar(&opts.Writer.Stdout, "write-stdout", true, log.Cyan("输出结果到控制台"))

	// 添加黑名单相关选项
	scanCmd.PersistentFlags().BoolVar(&opts.Scan.EnableBlacklist, "enable-blacklist", true, log.Cyan("启用URL黑名单检查"))
	scanCmd.PersistentFlags().BoolVar(&opts.Scan.DefaultBlacklist, "default-blacklist", true, log.Cyan("使用默认黑名单规则"))
	scanCmd.PersistentFlags().StringSliceVar(&opts.Scan.BlacklistPatterns, "blacklist-pattern", []string{}, log.Cyan("添加自定义黑名单规则 (可多次使用)"))
	scanCmd.PersistentFlags().StringVar(&opts.Scan.BlacklistFile, "blacklist-file", "", log.Cyan("黑名单规则文件路径"))

	log.Debug(log.Green("已注册scan命令"))
}
