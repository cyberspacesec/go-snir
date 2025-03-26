package ascii

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/fatih/color"
)

// 编译时注入的版本信息
var (
	version   = "v0.0.1"  // 默认版本
	commit    = "unknown" // Git提交哈希
	buildDate = "unknown" // 构建日期
	buildTime = "unknown" // 构建时间
)

// Logo returns the ASCII art logo for go-snir
func Logo() string {
	blue := color.New(color.FgBlue).SprintFunc()
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	logo := `
  ____           ____       _      
 / ___| ___     / ___|  ___(_)_ __ 
| |  _ / _ \____\___ \ / __| | '__|
| |_| | (_)|_____|__) | (__| | |   
 \____|\___/    |____/ \___|_|_|   
                                   
`
	coloredLogo := blue(logo)
	info := fmt.Sprintf("\n%s\n%s: %s\n%s: %s\n",
		yellow("一个强大的网页截图和信息收集工具"),
		cyan("版本"),
		green(version),
		cyan("项目地址"),
		"https://github.com/cyberspacesec/go-snir")

	return coloredLogo + info
}

// VersionInfo 返回详细的版本信息
func VersionInfo() string {
	cyan := color.New(color.FgCyan).SprintFunc()
	return fmt.Sprintf("版本: %s\n提交: %s\n构建时间: %s %s\n项目地址: %s\n",
		version, commit, buildDate, buildTime,
		cyan("https://github.com/cyberspacesec/go-snir"))
}

// Markdown renders markdown to the terminal
func Markdown(markdown string) string {
	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(100),
	)

	out, err := r.Render(markdown)
	if err != nil {
		return fmt.Sprintf("渲染Markdown时出错: %s\n%s", err, markdown)
	}

	return strings.TrimSuffix(out, "\n")
}
