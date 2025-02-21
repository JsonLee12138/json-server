package cmd

import (
	"fmt"

	"github.com/JsonLee12138/jsonix/pkg/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

func initRun(cmd *cobra.Command, args []string) error {
	return utils.TryCatchVoid(func() {
		prompt := promptui.Select{
			Label: "Please select the template source (github/gitee)",
			Items: []string{"github", "gitee"},
		}
		_, result, err := prompt.Run()
		if err != nil {
			panic(fmt.Errorf("❌ Error selecting template source: %s", err))
		}
		cmd.Println("✅ 选择模板来源: ", result)
	}, utils.DefaultErrorHandler)
}

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "A init for Jsonix",
	Args:  cobra.NoArgs,
	RunE:  initRun,
}
