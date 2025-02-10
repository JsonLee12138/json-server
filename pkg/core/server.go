package core

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/JsonLee12138/json-server/internal/global"
	"github.com/JsonLee12138/json-server/pkg/utils"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

func NewApp() *fiber.App {
	cnf := utils.Raise(NewConfig())
	utils.RaiseVoid(cnf.Bind(&global.Config))
	ValidatorSetup()
	app := fiber.New(fiber.Config{
		AppName:                  global.Config.System.AppName,
		DisableStartupMessage:    Mode() == ProMode,
		EnableIPValidation:       global.Config.System.IPValidationAble,
		EnablePrintRoutes:        global.Config.System.RoutesPrintAble,
		EnableSplittingOnParsers: global.Config.System.QuerySplitAble,
		EnableTrustedProxyCheck:  global.Config.System.ProxyCheckAble,
		ReadTimeout:              30 * time.Second,
		WriteTimeout:             30 * time.Second,
		JSONEncoder:              MarshalForFiber,
		JSONDecoder:              UnmarshalForFiber,
	})

	cmd := exec.Command("go", "generate", "./...")
	cmd.Dir = "./"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	utils.RaiseVoid(cmd.Run())

	if Mode() != ProMode {
		app.Use(swagger.New(swagger.Config{
			BasePath: "/",
			FilePath: "./docs/swagger.json",
			Path:     "swagger",
		}))
	}
	return app
}

func StartApp(app *fiber.App) {
	utils.RaiseVoid(app.Listen(fmt.Sprintf(":%s", global.Config.System.Port)))
}
