package core

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/JsonLee12138/json-server/internal/global"
	"github.com/JsonLee12138/json-server/pkg/configs"
	"github.com/JsonLee12138/json-server/pkg/utils"
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
)

var (
	apifox = new(Apifox)
)

func goGenerateCmd() error {
	cmd := exec.Command("go", "generate", "./...")
	cmd.Dir = "./"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

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

	utils.RaiseVoid(goGenerateCmd())

	if global.Config.System.SwaggerAble || global.Config.System.ApifoxAble || global.Config.System.OpenApiAble {
		utils.RaiseVoid(utils.SwaggerInitCmd())
	}

	if global.Config.System.SwaggerAble {
		app.Use(swagger.New(swagger.Config{
			BasePath: global.Config.Swagger.Get("BasePath").(string),
			FilePath: global.Config.Swagger.Get("FilePath").(string),
			Path:     global.Config.Swagger.Get("Path").(string),
			Title:    global.Config.Swagger.Title,
			CacheAge: global.Config.Swagger.Get("CacheAge").(int),
		}))
	}

	if global.Config.System.OpenApiAble {
		app.Static("/docs", "./docs")
	}

	if global.Config.System.ApifoxAble {
		docs := utils.Raise(os.ReadFile(utils.DefaultIfEmpty(global.Config.Apifox.FilePath, configs.DefaultFilePath)))
		err := apifox.Import(docs)
		if err != nil {
			panic(err)
		}
		fmt.Println("import apifox success")
	}

	return app
}

func StartApp(app *fiber.App) {
	utils.RaiseVoid(app.Listen(fmt.Sprintf(":%s", global.Config.System.Port)))
}
