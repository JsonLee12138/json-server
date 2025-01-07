package controller

import (
	"go.uber.org/dig"
	"json-server/internal/apps/example/service"
)

func ControllerSetup(container *dig.Container) error {
	container.Provide(service.NewExampleService)
	container.Provide(NewExampleController)
	return nil
}
