package core

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	Validator *validator.Validate
	once      sync.Once
)

func ValidatorSetup() {
	once.Do(func() {
		Validator = validator.New()
	})
}
