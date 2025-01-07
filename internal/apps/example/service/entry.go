package service

type ExampleService struct {
}

func NewExampleService() *ExampleService {
	return &ExampleService{}
}

func (service *ExampleService) Health() string {
	return "health"
}
