package controller

import (
	"github.com/JsonLee12138/json-server/internal/apps/example/service"
	"go.uber.org/dig"
)

func ControllerSetup(container *dig.Container) error {
	container.Provide(service.NewExampleService)
	container.Provide(NewExampleController)
	return nil
}
