package cmd

import (
	"fmt"
	"github.com/JsonLee12138/json-server/core"
	"github.com/JsonLee12138/json-server/core/generate"
	"github.com/spf13/cobra"
	"os"
)

func GenerateSetup(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("module", "m", "", "Specify the module name")
	cmd.PersistentFlags().StringP("service", "s", "", "Specify the service name")
	cmd.PersistentFlags().StringP("controller", "c", "", "Specify the controller name")
	cmd.PersistentFlags().StringP("repository", "r", "", "Specify the controller name")
	cmd.PersistentFlags().StringP("entity", "e", "", "Specify the controller name")
	cmd.PersistentFlags().BoolP("override", "o", false, "Specify the output directory")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		moduleName, _ := cmd.Flags().GetString("module")
		serviceName, _ := cmd.Flags().GetString("service")
		controllerName, _ := cmd.Flags().GetString("controller")
		repositoryName, _ := cmd.Flags().GetString("repository")
		entityName, _ := cmd.Flags().GetString("entity")
		override, _ := cmd.Flags().GetBool("override")
		currentPath, _ := os.Getwd()
		if moduleName != "" {
			modulePath := fmt.Sprintf("%s/%s", currentPath, moduleName)
			core.RaiseVoid(generate.GenerateModule(moduleName, modulePath))
			os.Exit(0)
		}
		if serviceName != "" {
			core.RaiseVoid(generate.GenerateInjectService(serviceName, fmt.Sprintf("%s", currentPath), override, ""))
		}
		if controllerName != "" {
			core.RaiseVoid(generate.GenerateInjectController(controllerName, fmt.Sprintf("%s", currentPath), override, ""))
		}
		if repositoryName != "" {
			core.RaiseVoid(generate.GenerateInjectRepository(repositoryName, fmt.Sprintf("%s", currentPath), override))
		}
		if entityName != "" {
			core.RaiseVoid(generate.GenerateEntity(entityName, fmt.Sprintf("%s", currentPath), override))
		}
		if moduleName == "" && serviceName == "" && controllerName == "" && repositoryName == "" && entityName == "" {
			cmd.Println("❌ Please specify a module, service, controller, repository, or entity name using -m or --module, -s or --service, -c or --controller, -r or --repository, or -e or --entity.")
			os.Exit(0)
		}
		return nil
	}
}

var GenerateCmd = &cobra.Command{
	Use:   "gen",
	Short: "A generator for Json Server",
	Args:  cobra.NoArgs,
}

func init() {
	GenerateSetup(GenerateCmd)
}
