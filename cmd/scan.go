package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cyberspacesec/go-snir/pkg/log"
	"github.com/cyberspacesec/go-snir/pkg/scan"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: log.Yellow("扫描并截图网站"),
	Long:  log.Yellow("扫描指定的URL、文件或网段，并对网站进行截图和信息收集"),
	Example: `  # 基本扫描单个网站
  ./snir scan example.com
  
  # 扫描单个网站并增加超时和延迟
  ./snir scan example.com --timeout 60 --delay 3
  
  # 从文件批量扫描
  ./snir scan file -f urls.txt
  
  # 扫描网段
  ./snir scan cidr 192.168.1.0/24
  
  # 保存HTML内容和HTTP头
  ./snir scan example.com --save-html --save-headers
  
  # 高分辨率截图
  ./snir scan example.com --resolution-x 1920 --resolution-y 1080
  
  # 使用代理
  ./snir scan example.com --proxy http://127.0.0.1:8080
  
  # 更多示例请查看 docs/usage_examples.md`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 如果直接提供了URL参数，则视为单URL扫描模式
		if len(args) == 1 {
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
		}

		// 如果没有提供参数，则显示帮助信息
		return cmd.Help()
	},
}

func init() {
	// 添加scan命令到根命令
	rootCmd.AddCommand(scanCmd)

	// 自定义帮助输出，为示例部分添加颜色
	defaultHelpFunc := scanCmd.HelpFunc()
	scanCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		// 保存原始示例
		originalExample := cmd.Example

		// 为示例添加颜色
		coloredExample := ""
		lines := strings.Split(originalExample, "\n")
		for _, line := range lines {
			// 为示例添加颜色
			if strings.HasPrefix(line, "  #") {
				coloredExample += log.Cyan(line) + "\n"
			} else if strings.HasPrefix(line, "  ./snir") {
				coloredExample += log.Yellow(line) + "\n"
			} else {
				coloredExample += line + "\n"
			}
		}
		cmd.Example = coloredExample

		// 调用默认帮助函数
		defaultHelpFunc(cmd, args)

		// 恢复原始示例
		cmd.Example = originalExample
	})

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
