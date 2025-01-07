package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"json-server/core"
)

func EnvSetup(cmd *cobra.Command) {
	var env string
	cmd.PersistentFlags().StringVarP(&env, "env", "e", "dev", "Set the application environment (dev, prod, test)")
	cobra.OnInitialize(func() {
		core.SetMode(core.ParseEnv(env))
		fmt.Printf("✅ Server is running in the %s environment\n", core.Mode())
	})
}
