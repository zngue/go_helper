package model

import (
	"gorm.io/gorm"
	"time"
)

type ZngModel struct {
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at" form:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"deleted_at;index" form:"deleted_at;index"`
}
