package cmd

import (
	"fmt"
	"github.com/JsonLee12138/json-server/core"
	"github.com/JsonLee12138/json-server/embed"
	"github.com/spf13/cobra"
	"os"
	"path"
)

func GenerateSetup(cmd *cobra.Command) {
	//cmd.PersistentFlags().BoolP("generate", "g", false, "Generate module")
	cmd.PersistentFlags().StringP("module", "m", "", "Specify the module name")

	// ✅ 使用 RunE，确保在 Flag 解析后执行逻辑
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		//generate, _ := cmd.Flags().GetBool("generate")
		moduleName, _ := cmd.Flags().GetString("module")

		//if !generate {
		//	fmt.Println("❌ Please use -g or --generate to create a new module.")
		//	return nil
		//}

		if moduleName == "" {
			fmt.Println("❌ Please specify a module name using -m or --module.")
			return nil
		}

		// ✅ 生成模块
		generateModule(moduleName)
		return nil
	}
}

func generateModule(moduleName string) {
	currentDir, _ := os.Getwd()
	basePath := path.Join(currentDir, moduleName)
	dirs := []string{"controller", "service", "model", "middleware", "router"}
	fmt.Println(basePath)
	// ✅ 创建模块目录
	for _, dir := range dirs {
		path := fmt.Sprintf("%s/%s", basePath, dir)
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			fmt.Printf("❌ Failed to create directory: %v\n", err)
			return
		}
	}
	SupperModuleName := core.UpperCamelCase(moduleName)
	pkgName, err := core.GetModuleName()
	if err != nil {
		fmt.Println(err)
	}
	if pkgName == "" {
		fmt.Println("❌ Could not determine module path. Ensure go.mod exists.")
		return
	}
	pkgPath, err := core.GetModuleFullPath(moduleName)
	if err != nil {
		fmt.Println(err)
		return
	}
	// ✅ 使用模板生成代码
	tmplFile, err := embed.TemplatesPath.ReadFile("templates/module.tmpl")
	if err != nil {
		fmt.Println(err, ">>>")
	}
	core.GenerateFromTemplateFile(string(tmplFile), fmt.Sprintf("%s/entry.go", basePath), map[string]string{
		"ModuleName":       moduleName,
		"SupperModuleName": SupperModuleName,
		"PkgPath":          pkgPath,
	})
	//generateFromTemplate("embed/controller.tmpl", fmt.Sprintf("%s/controller/%s_controller.go", basePath, moduleName), moduleName)
	//generateFromTemplate("embed/service.tmpl", fmt.Sprintf("%s/service/%s_service.go", basePath, moduleName), moduleName)
	//generateFromTemplate("embed/router.tmpl", fmt.Sprintf("%s/router/router.go", basePath), moduleName)

	fmt.Printf("✅ Module '%s' has been successfully generated!\n", moduleName)
}

var GenerateCmd = &cobra.Command{
	Use:   "gen",
	Short: "A generator for Json Server",
}

func init() {
	GenerateSetup(GenerateCmd)
}

//func Execute() {
//	if err := GenerateCmd.Execute(); err != nil {
//		fmt.Printf("❌ CLI Error: %v\n", err)
//	}
//}
