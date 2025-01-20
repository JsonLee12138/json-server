package core

import (
	"github.com/rs/xid"
	"gorm.io/gorm"
	"time"
)

type BaseEntity struct {
	CreatedAt time.Time      `json:"createdAt" gorm:"index"`
	UpdatedAt time.Time      `json:"updatedAt"  gorm:"index"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedBy string         `json:"createdBy"`
	UpdatedBy string         `json:"updatedBy"`
	DeletedBy string         `json:"deletedBy"`
}

type BaseEntityWithID struct {
	ID uint `json:"id" gorm:"primarykey"`
	BaseEntity
}

type BaseEntityWithUuid struct {
	ID string `json:"id" gorm:"primarykey;type:char(20)" json:"id"`
	BaseEntity
}

func (e *BaseEntityWithUuid) BeforeCreate() error {
	if e.ID == "" {
		e.ID = GenerateUUID()
	}
	return nil
}

func GenerateUUID() string {
	return xid.New().String()
}
