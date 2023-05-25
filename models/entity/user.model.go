package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Name     string `json:"name" `
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	// Todo      []TodoResponse `json:"todo"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name" `
	Email string `json:"email"`
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt gorm.DeletedAt `json:"-"`
}

func (UserResponse) TableName() string {
	return "todos"
}
