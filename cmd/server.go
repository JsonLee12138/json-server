package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/JsonLee12138/json-server/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	watchable bool
	showPort  string
	killPort  string
	env       string
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "A flexible Go configuration management tool",
	Run: func(cmd *cobra.Command, args []string) {
		runApp()
	},
}

// 初始化 CLI 标志
func init() {
	ServerCmd.PersistentFlags().StringVarP(&env, "env", "e", "dev", "Set the application environment (dev, prod, test)")
	ServerCmd.PersistentFlags().BoolVarP(&watchable, "watch", "w", false, "Enable config file watching")
	ServerCmd.PersistentFlags().StringVarP(&showPort, "show", "s", "", "Show port")
	ServerCmd.PersistentFlags().StringVarP(&killPort, "kill", "k", "", "Kill port")
}

// 核心运行逻辑
func runApp() {
	if showPort != "" {
		pid, err := utils.FindPIDByPort(showPort)
		if err != nil {
			panic(err)
		}
		if pid == "" {
			fmt.Println("No process found on port")
		} else {
			fmt.Println("Process running on port", showPort, "is", pid)
		}
		os.Exit(0)
	}
	if killPort != "" {
		pid, err := utils.FindPIDByPort(killPort)
		if err != nil {
			panic(err)
		}
		if pid == "" {
			fmt.Println("No process found on port")
		} else {
			fmt.Println("Process running on port", killPort, "is", pid)
			err = utils.KillProcess(pid)
			if err != nil {
				panic(err)
			}
			fmt.Println("Process killed")
		}
		os.Exit(0)
	}
	var cmd *exec.Cmd
	if watchable {
		cmd = exec.Command("air")
	} else {
		cmd = exec.Command("go", "run", "main.go")
	}
	cmd.Dir = "."
	// 设置进程的标准输出和标准错误输出
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	ModeEnvHandler(env)
	// 执行命令
	utils.RaiseVoid(cmd.Start())
	go func() {
		err := cmd.Wait()
		if err != nil {
			fmt.Printf("Error occurred during the execution: %v\n", err)
		} else {
			fmt.Println("Command executed successfully!")
		}
	}()

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	<-signalChannel

	utils.RaiseVoid(cmd.Process.Kill())
}
