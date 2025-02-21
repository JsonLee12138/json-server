package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// RootCmd 定义 CLI 的根命令
var RootCmd = &cobra.Command{
	Use:   "jsonix",
	Short: "A server for web applications",
}

func init() {
	RootCmd.AddCommand(ServerCmd, GenerateCmd, AutoMigrateCmd, InitCmd)
}

// Execute 入口，加载所有命令
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		RootCmd.PrintErrln(err)
		os.Exit(1)
	}
}
