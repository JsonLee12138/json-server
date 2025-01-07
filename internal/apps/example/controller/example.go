package controller

import (
	"github.com/gofiber/fiber/v2"
	"json-server/internal/apps/example/service"
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
