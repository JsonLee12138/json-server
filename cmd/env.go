package cmd

import (
	"fmt"
	"github.com/JsonLee12138/json-server/pkg/core"
	"github.com/spf13/cobra"
)

func EnvSetup(cmd *cobra.Command) {
	var env string
	cmd.PersistentFlags().StringVarP(&env, "env", "e", "dev", "Set the application environment (dev, prod, test)")
	cobra.OnInitialize(func() {
		core.SetMode(core.ParseEnv(env))
		fmt.Printf("✅ Server is running in the %s environment\n", core.Mode())
	})
}

func ModeEnvHandler(e string) {
	core.SetMode(core.ParseEnv(e))
	fmt.Printf("✅ Server is running in the %s environment\n", core.Mode())
}
