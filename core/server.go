package core

import (
	"fmt"
	"github.com/JsonLee12138/json-server/internal/global"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"time"
)

func NewApp() *fiber.App {
	cnf := Raise(NewConfig())
	RaiseVoid(cnf.Bind(&global.Config))
	app := fiber.New(fiber.Config{
		AppName:                  global.Config.System.AppName,
		DisableStartupMessage:    Mode() == ProMode,
		EnableIPValidation:       global.Config.System.IPValidationAble,
		EnablePrintRoutes:        global.Config.System.RoutesPrintAble,
		EnableSplittingOnParsers: global.Config.System.QuerySplitAble,
		EnableTrustedProxyCheck:  global.Config.System.ProxyCheckAble,
		ReadTimeout:              30 * time.Second,
		WriteTimeout:             30 * time.Second,
	})
	if Mode() != ProMode {
		app.Get("/", swagger.HandlerDefault)
	}
	return app
}

func StartApp(app *fiber.App) {
	fmt.Println(global.Config.System.Port, "port")
	RaiseVoid(app.Listen(fmt.Sprintf(":%s", global.Config.System.Port)))
}
