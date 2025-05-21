package models

import (
	"time"

	"gorm.io/gorm"
)

// GSModel is the base model for all models in the application
type GSModel struct {
	ID        uint64         `json:"id" gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
} 