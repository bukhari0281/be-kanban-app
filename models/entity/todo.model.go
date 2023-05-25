package entity

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID         uint    `json:"id" gorm:"primary_key"`
	Name       string  `json:"name" gorm:"not null"`
	Note       *string `json:"note" gorm:""`
	IsComplete bool    `json:"is_completed" gorm:"boolean;default:false;not null"`
	User       *User   `json:"user"`
	CategoryID uint    `json:"-"`
	// Category   CategoryRespon       `json:"category"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type TodoResponse struct {
	ID         uint    `json:"id"`
	Name       string  `json:"name"`
	Note       *string `json:"note"`
	IsComplete bool    `json:"is_completed"`
	CategoryID uint    `json:"-"`
}

func (TodoResponse) TableName() string {
	return "todos"
}
