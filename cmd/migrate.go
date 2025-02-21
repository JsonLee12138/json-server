package cmd

import (
	"github.com/JsonLee12138/jsonix/pkg/core"
	"github.com/spf13/cobra"
)

func AutoMigrateRun(cmd *cobra.Command, args []string) error {
	root, _ := cmd.Flags().GetString("root")
	dest, _ := cmd.Flags().GetString("dest")
	err := core.AggregateEntities(root, dest)
	if err != nil {
		return err
	}
	return nil
}

func AutoMigrateSetup(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("root", "r", "./", "Specify the root directory")
	cmd.PersistentFlags().StringP("dest", "d", "./auto_migrate", "Specify the database type")
}

var AutoMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A auto_migrate for Jsonix",
	Args:  cobra.NoArgs,
	RunE:  AutoMigrateRun,
}

func init() {
	AutoMigrateSetup(AutoMigrateCmd)
}
