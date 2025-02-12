package core

import (
	"time"

	"github.com/rs/xid"
	"gorm.io/gorm"
)

//go:generate easyjson -all ./entity.go

// @Description 核心基础实体
type BaseEntity struct {
	CreatedAt time.Time      `json:"createdAt" gorm:"index"`                     // 创建时间
	UpdatedAt time.Time      `json:"updatedAt"  gorm:"index"`                    // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-,intern" swaggerignore:"true"` // 删除时间
	CreatedBy string         `json:"createdBy"`                                  // 创建人
	UpdatedBy string         `json:"updatedBy"`                                  // 更新人
	DeletedBy string         `json:"deletedBy,intern" swaggerignore:"true"`      // 删除人
}

// @Description 核心基础实体，带 ID
type BaseEntityWithID struct {
	ID uint `json:"id" gorm:"primarykey"` // 主键ID
	BaseEntity
}

// @Description 核心基础实体，带 UUID
type BaseEntityWithUuid struct {
	ID string `json:"id" gorm:"primarykey;type:char(20)"` // 主键ID
	BaseEntity
}

func (e *BaseEntityWithUuid) BeforeCreate(tx *gorm.DB) error {
	if e.ID == "" || e.ID == xid.NilID().String() {
		e.ID = GenerateUUID()
	}
	return nil
}

func GenerateUUID() string {
	return xid.New().String()
}
