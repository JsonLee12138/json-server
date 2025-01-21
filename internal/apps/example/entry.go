package example

import (
	"github.com/JsonLee12138/json-server/core"
	"github.com/JsonLee12138/json-server/internal/apps/example/controller"
	"github.com/JsonLee12138/json-server/internal/apps/example/repository"
	"github.com/JsonLee12138/json-server/internal/apps/example/service"
	"go.uber.org/dig"
)

// RegisterexampleModule 在 DI 容器中注册 example 模块
func ExampleModuleSetup(container *dig.Container) error {
	return core.TryCatchVoid(func() {
		core.RaiseVoid(container.Provide(controller.NewExampleController))
		core.RaiseVoid(container.Provide(service.NewExampleService))
		core.RaiseVoid(container.Provide(service.NewTestService))
		core.RaiseVoid(container.Provide(repository.NewExampleRepository))
	}, core.DefaultErrorHandler)
}

//func provideControllers(container *dig.Container) error {
//
//}
