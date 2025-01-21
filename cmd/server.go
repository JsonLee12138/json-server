package cmd

import (
	"github.com/JsonLee12138/json-server/core"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

// var (
// 	watchable bool
// )

// RootCmd 定义 CLI 根命令
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "A flexible Go configuration management tool",
	Run: func(cmd *cobra.Command, args []string) {
		EnvSetup(cmd)
		runApp(cmd)
	},
}

// 初始化 CLI 标志
func init() {
	// ServerCmd.PersistentFlags().BoolVar(&watchable, "watch", false, "Enable config file watching")
	ServerCmd.PersistentFlags().BoolP("watch", "w", false, "Enable config file watching")
}

// 核心运行逻辑
func runApp(mainCmd *cobra.Command) {
	watchable := core.Raise(mainCmd.Flags().GetBool("watch"))
	var cmd *exec.Cmd
	if watchable {
		// TODO: 监听配置文件变化
		cmd = exec.Command("air")
	} else {
		// TODO: 不监听配置文件变化
		cmd = exec.Command("go", "run", "main.go")
	}
	cmd.Dir = "."
	// 设置进程的标准输出和标准错误输出
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	core.RaiseVoid(cmd.Run())
}
