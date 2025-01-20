package repository

import (
	"gorm.io/gorm"
)

type ExampleRepository struct {
	db *gorm.DB
}

type ExampleRepositoryDeps struct {
	DB *gorm.DB
}

func NewExampleRepository(deps ExampleRepositoryDeps) *ExampleRepository {
	return &ExampleRepository{
		db: deps.DB,
	}
}
