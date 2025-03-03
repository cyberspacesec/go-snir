package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/spf13/cobra"

	"github.com/cyberspacesec/go-web-screenshot/pkg/log"
	"github.com/cyberspacesec/go-web-screenshot/pkg/models"
	"github.com/cyberspacesec/go-web-screenshot/pkg/runner"
)

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "从文件批量扫描URL",
	Long:  "从文件中读取URL列表进行批量扫描和截图",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 检查文件路径是否提供
		if opts.Scan.FilePath == "" {
			return fmt.Errorf("请使用 -f 或 --file 参数指定URL文件路径")
		}

		// 打开文件
		file, err := os.Open(opts.Scan.FilePath)
		if err != nil {
			return fmt.Errorf("无法打开文件: %v", err)
		}
		defer file.Close()

		// 读取URL列表
		var urls []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			url := strings.TrimSpace(scanner.Text())
			if url != "" && !strings.HasPrefix(url, "#") {
				urls = append(urls, url)
			}
		}

		if err := scanner.Err(); err != nil {
			return fmt.Errorf("读取文件时出错: %v", err)
		}

		if len(urls) == 0 {
			return fmt.Errorf("文件中没有有效的URL")
		}

		log.Info("从文件中读取URL", "count", len(urls), "file", opts.Scan.FilePath)

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

		// 使用goroutine池并发处理URL
		var wg sync.WaitGroup
		semaphore := make(chan struct{}, opts.Scan.Threads)
		resultsChan := make(chan *models.Result, opts.Scan.Threads)

		// 启动结果处理goroutine
		go func() {
			for result := range resultsChan {
				printResult(result)
			}
		}()

		// 处理每个URL
		for _, url := range urls {
			wg.Add(1)
			semaphore <- struct{}{}

			go func(target string) {
				defer wg.Done()
				defer func() { <-semaphore }()

				// 确保URL格式正确
				if !strings.HasPrefix(target, "http://") && !strings.HasPrefix(target, "https://") {
					// 尝试HTTPS
					if opts.Scan.HTTPS {
						target = "https://" + target
					} else if opts.Scan.HTTP {
						target = "http://" + target
					} else {
						log.Error("URL缺少协议前缀且未启用HTTP或HTTPS选项", "url", target)
						return
					}
				}

				log.Info("开始扫描URL", "url", target)

				// 执行截图
				result, err := driver.Witness(target, r)
				if err != nil {
					log.Error("截图失败", "url", target, "error", err)
					return
				}

				// 发送结果到通道
				resultsChan <- result
			}(url)
		}

		// 等待所有goroutine完成
		wg.Wait()
		close(semaphore)
		close(resultsChan)

		log.Info("批量扫描完成", "total", len(urls))
		return nil
	},
}

func init() {
	scanCmd.AddCommand(fileCmd)
	
	// 添加文件相关选项
	fileCmd.Flags().StringVarP(&opts.Scan.FilePath, "file", "f", "", "包含URL列表的文件路径")
	fileCmd.MarkFlagRequired("file")
	
	log.Debug("已注册file命令")
}