package example

import (
	"github.com/JsonLee12138/jsonix/internal/apps/example/controller"
	"github.com/JsonLee12138/jsonix/internal/apps/example/repository"
	"github.com/JsonLee12138/jsonix/internal/apps/example/service"
	"github.com/JsonLee12138/jsonix/pkg/utils"
	"go.uber.org/dig"
)

// RegisterexampleModule 在 DI 容器中注册 example 模块
func ExampleModuleSetup(container *dig.Container) error {
	return utils.TryCatchVoid(func() {
		utils.RaiseVoid(container.Provide(controller.NewExampleController))
		utils.RaiseVoid(container.Provide(service.NewExampleService))
		utils.RaiseVoid(container.Provide(service.NewTestService))
		utils.RaiseVoid(container.Provide(repository.NewExampleRepository))
	}, utils.DefaultErrorHandler)
}

//func provideControllers(container *dig.Container) error {
//
//}
