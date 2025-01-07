package controller

import (
	"github.com/JsonLee12138/json-server/internal/apps/example/service"
	"github.com/gofiber/fiber/v2"
)

type ExampleController struct {
	exampleService *service.ExampleService
}

func NewExampleController(exampleService *service.ExampleService) *ExampleController {
	return &ExampleController{exampleService: exampleService}
}

func (ctrl *ExampleController) Hello(c *fiber.Ctx) error {
	return c.SendString(ctrl.exampleService.Health())
}
