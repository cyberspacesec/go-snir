package cmd

import (
	"fmt"
	"net"
	"sync"

	"github.com/spf13/cobra"

	"github.com/cyberspacesec/go-web-screenshot/pkg/log"
	"github.com/cyberspacesec/go-web-screenshot/pkg/models"
	"github.com/cyberspacesec/go-web-screenshot/pkg/runner"
)

var cidrCmd = &cobra.Command{
	Use:   "cidr [cidr]",
	Short: "扫描网段",
	Long:  "扫描指定CIDR网段中的所有IP地址并进行截图",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cidr := args[0]

		// 解析CIDR
		ip, ipnet, err := net.ParseCIDR(cidr)
		if err != nil {
			return fmt.Errorf("无效的CIDR格式: %v", err)
		}

		// 获取网段中的所有IP
		var ips []string
		for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
			ips = append(ips, ip.String())
		}

		// 移除网络地址和广播地址
		if len(ips) > 2 {
			ips = ips[1 : len(ips)-1]
		}

		if len(ips) == 0 {
			return fmt.Errorf("网段中没有有效的IP地址")
		}

		log.Info("从CIDR中解析IP", "count", len(ips), "cidr", cidr)

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

		// 使用goroutine池并发处理IP
		var wg sync.WaitGroup
		semaphore := make(chan struct{}, opts.Scan.Threads)
		resultsChan := make(chan *models.Result, opts.Scan.Threads)

		// 启动结果处理goroutine
		go func() {
			for result := range resultsChan {
				printResult(result)
			}
		}()

		// 处理每个IP
		for _, ip := range ips {
			// 如果指定了端口，则为每个端口创建一个任务
			ports := opts.Scan.Ports
			if len(ports) == 0 {
				ports = []int{80, 443} // 默认扫描80和443端口
			}

			for _, port := range ports {
				wg.Add(1)
				semaphore <- struct{}{}

				go func(ip string, port int) {
					defer wg.Done()
					defer func() { <-semaphore }()

					// 构建目标URL
					var target string
					if port == 80 && opts.Scan.HTTP {
						target = fmt.Sprintf("http://%s", ip)
					} else if port == 443 && opts.Scan.HTTPS {
						target = fmt.Sprintf("https://%s", ip)
					} else if port == 80 {
						target = fmt.Sprintf("http://%s", ip)
					} else if port == 443 {
						target = fmt.Sprintf("https://%s", ip)
					} else if opts.Scan.HTTPS {
						target = fmt.Sprintf("https://%s:%d", ip, port)
					} else if opts.Scan.HTTP {
						target = fmt.Sprintf("http://%s:%d", ip, port)
					} else {
						log.Error("未启用HTTP或HTTPS选项", "ip", ip, "port", port)
						return
					}

					log.Info("开始扫描IP", "url", target)

					// 执行截图
					result, err := driver.Witness(target, r)
					if err != nil {
						log.Error("截图失败", "url", target, "error", err)
						return
					}

					// 发送结果到通道
					resultsChan <- result
				}(ip, port)
			}
		}

		// 等待所有goroutine完成
		wg.Wait()
		close(semaphore)
		close(resultsChan)

		log.Info("网段扫描完成", "total", len(ips), "cidr", cidr)
		return nil
	},
}

// inc 函数用于递增IP地址
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func init() {
	scanCmd.AddCommand(cidrCmd)
	
	// 添加网段扫描相关选项
	cidrCmd.Flags().IntSliceVar(&opts.Scan.Ports, "ports", []int{80, 443}, "要扫描的端口列表")
	
	log.Debug("已注册cidr命令")
}