package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/JsonLee12138/jsonix/pkg/core/generate"
	"github.com/JsonLee12138/jsonix/pkg/utils"
	"github.com/spf13/cobra"
)

func generateRun(cmd *cobra.Command, args []string) error {
	return utils.TryCatchVoid(func() {
		currentPath, _ := os.Getwd()
		moduleName, _ := cmd.Flags().GetString("module")
		if moduleName != "" {
			modulePath := fmt.Sprintf("%s/%s", currentPath, moduleName)
			utils.RaiseVoid(generate.GenerateModule(moduleName, modulePath))
			return
		}
		if len(args) == 0 {
			panic(errors.New("❌ Please specify a service, controller, repository, or entity name using -s or --service, -c or --controller, -r or --repository, or -e or --entity"))
		}

		name := args[0]

		serviceFlag, _ := cmd.Flags().GetBool("service")
		controllerFlag, _ := cmd.Flags().GetBool("controller")
		repositoryFlag, _ := cmd.Flags().GetBool("repository")
		entityFlag, _ := cmd.Flags().GetBool("entity")
		override, _ := cmd.Flags().GetBool("override")

		if serviceFlag {
			utils.RaiseVoid(generate.GenerateInjectService(name, currentPath, override, ""))
		}
		if controllerFlag {
			utils.RaiseVoid(generate.GenerateInjectController(name, currentPath, override, ""))
		}
		if repositoryFlag {
			utils.RaiseVoid(generate.GenerateInjectRepository(name, currentPath, override))
		}
		if entityFlag {
			utils.RaiseVoid(generate.GenerateEntity(name, currentPath, override))
		}
	}, utils.DefaultErrorHandler)
}

func generateSetup(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("module", "m", "", "Specify the module name")
	cmd.PersistentFlags().BoolP("service", "s", false, "Specify the service name")
	cmd.PersistentFlags().BoolP("controller", "c", false, "Specify the controller name")
	cmd.PersistentFlags().BoolP("repository", "r", false, "Specify the controller name")
	cmd.PersistentFlags().BoolP("entity", "e", false, "Specify the controller name")
	cmd.PersistentFlags().BoolP("override", "o", false, "Specify the output directory")
}

var GenerateCmd = &cobra.Command{
	Use:   "gen",
	Short: "A generator for Jsonix",
	Args:  cobra.ArbitraryArgs,
	RunE:  generateRun,
}

func init() {
	generateSetup(GenerateCmd)
}
