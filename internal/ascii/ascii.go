package ascii

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/glamour"
)

// Logo returns the ASCII art logo for go-web-screenshot
func Logo() string {
	return `
  ____         __        __   _       ____                                  _           _   
 / ___| ___    \ \      / /__| |__   / ___|  ___ _ __ ___  ___ _ __  _ __ | |__   ___ | |_ 
| |  _ / _ \    \ \ /\ / / _ \ '_ \  \___ \ / __| '__/ _ \/ _ \ '_ \| '_ \| '_ \ / _ \| __|
| |_| | (_) |    \ V  V /  __/ |_) |  ___) | (__| | |  __/  __/ | | | | | | | | | (_) | |_ 
 \____|\___/      \_/\_/ \___|_.__/  |____/ \___|_|  \___|\___|_| |_|_| |_|_| |_|\___/ \__|
                                                                                           

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