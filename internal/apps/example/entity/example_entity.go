package entity

import (
	"github.com/JsonLee12138/json-server/core"
)

type ExampleEntity struct {
	core.BaseEntityWithUuid
}

func (e *ExampleEntity) BeforeCreate() error {
	return e.BaseEntityWithUuid.BeforeCreate()
}
