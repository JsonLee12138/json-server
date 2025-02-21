package service

import (
	"github.com/JsonLee12138/jsonix/internal/apps/example/repository"
)

type TestService struct {
	repository *repository.ExampleRepository
}

//type ExampleServiceDeps struct {
//	Repository *repository.ExampleRepository
//}

//func NewExampleService(deps ExampleServiceDeps) *ExampleService {
//	return &ExampleService{
//		repository: deps.Repository,
//	}
//}

func NewTestService(repository *repository.ExampleRepository) *TestService {
	return &TestService{
		repository,
	}
}

func (s *TestService) Health() string {
	return "Hello Test!"
}
