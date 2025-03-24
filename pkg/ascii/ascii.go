package ascii

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/glamour"
)

// Logo returns the ASCII art logo for go-snir
func Logo() string {
	return `
  ____           ____       _      
 / ___| ___     / ___|  ___(_)_ __ 
| |  _ / _ \____\___ \ / __| | '__|
| |_| | (_)|_____|__) | (__| | |   
 \____|\___/    |____/ \___|_|_|   
                                   

一个强大的网页截图和信息收集工具
版本: v0.1.0
`
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
