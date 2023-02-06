package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID         uint           `json:"id" gorm:"primary_key"`
	Name       string         `json:"name" gorm:"not null"`
	Note       *string        `json:"note" gorm:""`
	IsComplete bool           `json:"is_completed" gorm:"boolean;default:false;not null"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-"`
}