package controller

import (
	"github.com/JsonLee12138/jsonix/internal/apps/example/service"
	"github.com/gofiber/fiber/v2"
)

type ExampleController struct {
	service     *service.ExampleService
	testService *service.TestService
}

type ExampleControllerDeps struct {
	Service *service.ExampleService
}

func NewExampleController(t *service.TestService, s *service.ExampleService) *ExampleController {
	return &ExampleController{
		service:     s,
		testService: t,
	}
}

// Health godoc
// @Summary Show the status of server.
// @Router /health [get]
func (c *ExampleController) Health(ctx *fiber.Ctx) error {
	return ctx.SendString(c.service.Health())
}

func (c *ExampleController) Test(ctx *fiber.Ctx) error {
	return ctx.SendString(c.testService.Health())
}
