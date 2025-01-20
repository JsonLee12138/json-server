package service

import (
	"github.com/JsonLee12138/json-server/internal/apps/example/repository"
)

type ExampleService struct {
	repository *repository.ExampleRepository
}

type ExampleServiceDeps struct {
	Repository *repository.ExampleRepository
}

func NewExampleService(deps ExampleServiceDeps) *ExampleService {
	return &ExampleService{
		repository: deps.Repository,
	}
}

func (s *ExampleService) Health() string {
	return "Hello World!"
}
