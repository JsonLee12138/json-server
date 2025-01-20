package controller

import (
	"fmt"
	"github.com/JsonLee12138/json-server/internal/apps/example/service"
	"github.com/gofiber/fiber/v2"
)

type ExampleController struct {
	service *service.ExampleService
}

type ExampleControllerDeps struct {
	Service *service.ExampleService
}

func NewExampleController(deps ExampleControllerDeps) *ExampleController {
	return &ExampleController{
		service: deps.Service,
	}
}

func (c *ExampleController) Health(ctx *fiber.Ctx) error {
	fmt.Println(1)
	return ctx.SendString(c.service.Health())
}
