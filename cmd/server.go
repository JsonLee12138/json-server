package cmd

import (
	"errors"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/JsonLee12138/jsonix/pkg/utils"
	"github.com/spf13/cobra"
)

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "A flexible Go configuration management tool",
	RunE:  serverRun,
}

func init() {
	serverSetup(ServerCmd)
}

func serverSetup(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("env", "e", "dev", "Set the application environment (dev, prod, test)")
	cmd.PersistentFlags().BoolP("watch", "w", false, "Enable config file watching")
	cmd.PersistentFlags().StringP("show", "s", "", "Show port")
	cmd.PersistentFlags().StringP("kill", "k", "", "Kill port")
}

func serverRun(cmd *cobra.Command, args []string) error {
	return utils.TryCatchVoid(func() {
		env, _ := cmd.Flags().GetString("env")
		watchable, _ := cmd.Flags().GetBool("watch")
		showPort, _ := cmd.Flags().GetString("show")
		killPort, _ := cmd.Flags().GetString("kill")
		if showPort != "" {
			pid := utils.Raise(utils.FindPIDByPort(showPort))
			if pid == "" {
				panic(errors.New("No process found on port"))
			}
			cmd.Println("Process running on port", showPort, "is", pid)
			return
		}
		if killPort != "" {
			pid := utils.Raise(utils.FindPIDByPort(killPort))
			if pid == "" {
				panic(errors.New("No process found on port"))
			}
			utils.RaiseVoid(utils.KillProcess(pid))
			cmd.Println("Process killed on port", killPort)
			return
		}
		var execCmd *exec.Cmd
		if watchable {
			execCmd = exec.Command("air")
		} else {
			execCmd = exec.Command("go", "run", "main.go")
		}
		execCmd.Dir = "."
		execCmd.Stdout = os.Stdout
		execCmd.Stderr = os.Stderr
		ModeEnvHandler(env)
		utils.RaiseVoid(execCmd.Start())
		go func() {
			err := execCmd.Wait()
			if err != nil {
				panic(err)
			} else {
				cmd.Println("Server started successfully!")
			}
		}()

		signalChannel := make(chan os.Signal, 1)
		signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

		<-signalChannel

		utils.RaiseVoid(execCmd.Process.Kill())
	}, utils.DefaultErrorHandler)
}
