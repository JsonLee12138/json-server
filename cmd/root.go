package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// RootCmd 定义 CLI 的根命令
var RootCmd = &cobra.Command{
	Use:   "json-server",
	Short: "A server for web applications",
}

func init() {
}

// Execute 入口，加载所有命令
func Execute() {
	RootCmd.AddCommand(ServerCmd, GenerateCmd)
	if err := RootCmd.Execute(); err != nil {
		fmt.Printf("❌ CLI Error: %v\n", err)
	}
}
