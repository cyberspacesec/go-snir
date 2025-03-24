package cmd

import (
	"fmt"
	"net"

	"github.com/spf13/cobra"

	"github.com/cyberspacesec/go-snir/pkg/log"
	"github.com/cyberspacesec/go-snir/pkg/scan"
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

		// 创建扫描配置
		config := &scan.Config{
			Targets: ips,
			Options: opts,
		}

		// 创建扫描器
		scanner, err := scan.NewScanner(config)
		if err != nil {
			return fmt.Errorf("创建扫描器失败: %v", err)
		}
		defer scanner.Close()

		// 执行扫描
		log.Info("开始扫描网段", "cidr", cidr, "ip_count", len(ips))
		err = scanner.ScanMulti(ips)
		if err != nil {
			return fmt.Errorf("扫描网段失败: %v", err)
		}

		log.Info("网段扫描完成", "cidr", cidr)
		return nil
	},
}

// inc 递增IP地址
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
	log.Debug("已注册cidr命令")
}
