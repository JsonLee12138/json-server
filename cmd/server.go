package cmd

import (
	"fmt"
	"github.com/JsonLee12138/json-server/core"
	"github.com/JsonLee12138/json-server/internal/apps/example"
	"github.com/JsonLee12138/json-server/internal/apps/example/controller"
	"github.com/JsonLee12138/json-server/internal/global"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
	"time"
)

var (
	watchable bool
)

// RootCmd 定义 CLI 根命令
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "A flexible Go configuration management tool",
	Run: func(cmd *cobra.Command, args []string) {
		EnvSetup(cmd)
		runApp()
	},
}

// 初始化 CLI 标志
func init() {
	ServerCmd.PersistentFlags().BoolVar(&watchable, "watch", false, "Enable config file watching")
}

// 执行 CLI
//func Execute() {
//	RootCmd.AddCommand(GenerateCmd)
//	if err := RootCmd.Execute(); err != nil {
//		fmt.Printf("❌ CLI Error: %v\n", err)
//	}
//}

// 核心运行逻辑
func runApp() {
	if cnf, err := core.NewConfig(); err != nil {
		panic(err)
	} else {
		err = cnf.Bind(&global.Config)
		if err != nil {
			panic(err)
		}
	}
	container := dig.New()
	example.ExampleModuleSetup(container)
	app := fiber.New(fiber.Config{
		AppName:                  global.Config.System.AppName,
		DisableStartupMessage:    core.Mode() == core.ProMode,
		EnableIPValidation:       global.Config.System.IPValidationAble,
		EnablePrintRoutes:        global.Config.System.RoutesPrintAble,
		EnableSplittingOnParsers: global.Config.System.QuerySplitAble,
		EnableTrustedProxyCheck:  global.Config.System.ProxyCheckAble,
		ReadTimeout:              30 * time.Second,
		WriteTimeout:             30 * time.Second,
	})
	// TODO: 路由目前是错误的, 后面也不需要改, 直接重构成使用注解来处理
	container.Invoke(func(ctrl *controller.ExampleController) {
		app.Get("/", ctrl.Health)
	})
	if err := app.Listen(fmt.Sprintf(":%s", global.Config.System.Port)); err != nil {
		panic(err)
	}
}
