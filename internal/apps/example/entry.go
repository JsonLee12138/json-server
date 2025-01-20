package example

import (
	"github.com/JsonLee12138/json-server/internal/apps/example/controller"
	"github.com/JsonLee12138/json-server/internal/apps/example/repository"
	"github.com/JsonLee12138/json-server/internal/apps/example/service"
	"go.uber.org/dig"
)

// RegisterexampleModule 在 DI 容器中注册 example 模块
func ExampleModuleSetup(container *dig.Container) {
	container.Provide(controller.NewExampleController)
	container.Provide(service.NewExampleService)
	container.Provide(repository.NewExampleRepository)
}
