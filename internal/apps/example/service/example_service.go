package service

import (
	"github.com/JsonLee12138/jsonix/internal/apps/example/repository"
)

type ExampleService struct {
	repository *repository.ExampleRepository
}

type ExampleServiceDeps struct {
	Repository *repository.ExampleRepository
}

//func NewExampleService(deps ExampleServiceDeps) *ExampleService {
//	return &ExampleService{
//		repository: deps.Repository,
//	}
//}

func NewExampleService(r *repository.ExampleRepository) *ExampleService {
	return &ExampleService{
		repository: r,
	}
}

func (s *ExampleService) Health() string {
	return "Hello World!"
}
