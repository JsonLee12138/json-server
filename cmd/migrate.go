package cmd

import (
	"github.com/JsonLee12138/json-server/pkg/core"
	"github.com/spf13/cobra"
)

func AutoMigrateSetup(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("root", "r", "./", "Specify the root directory")
	cmd.PersistentFlags().StringP("dest", "d", "./auto_migrate_local", "Specify the database type")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		root, _ := cmd.Flags().GetString("root")
		dest, _ := cmd.Flags().GetString("dest")
		err := core.AggregateEntities(root, dest)
		if err != nil {
			return err
		}
		return nil
	}
}

var AutoMigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "A auto_migrate for Json Server",
	Args:  cobra.NoArgs,
}

func init() {
	AutoMigrateSetup(AutoMigrateCmd)
}
