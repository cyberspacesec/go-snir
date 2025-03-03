package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/cyberspacesec/go-web-screenshot/internal/ascii"
	"github.com/cyberspacesec/go-web-screenshot/pkg/log"
	"github.com/cyberspacesec/go-web-screenshot/pkg/runner"
)

var (
	opts = &runner.Options{}
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var rootCmd = &cobra.Command{
	Use:   "go-web-screenshot",
	Short: "一个网页截图和信息收集工具",
	Long:  ascii.Logo(),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if opts.Logging.Silence {
			log.EnableSilence()
		}

		if opts.Logging.Debug && !opts.Logging.Silence {
			log.EnableDebug()
			log.Debug("调试日志已启用")
		}

		return nil
	},
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.SilenceErrors = true
	err := rootCmd.Execute()
	if err != nil {
		var cmd string
		c, _, cerr := rootCmd.Find(os.Args[1:])
		if cerr == nil {
			cmd = c.Name()
		}

		v := "\n"

		if cmd != "" {
			v += fmt.Sprintf("运行 `%s` 命令时发生错误\n", cmd)
		} else {
			v += "发生了一个错误。 "
		}

		v += "错误信息为:\n\n" + fmt.Sprintf("```%s```", err)
		fmt.Println(ascii.Markdown(v))

		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&opts.Logging.Debug, "debug-log", "D", false, "启用调试日志")
	rootCmd.PersistentFlags().BoolVarP(&opts.Logging.Silence, "quiet", "q", false, "静默（几乎所有）日志")
}
