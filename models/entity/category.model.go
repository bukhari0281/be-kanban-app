package entity

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	Category  string         `json:"category" gorm:"not null"`
	Todo      []TodoResponse `json:"todo"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

// type CategoryRespon struct {
// 	ID       uint   `json:"id"`
// 	Category string `json:"category" `
// 	// CreatedAt time.Time `json:"created_at"`
// 	// UpdatedAt time.Time `json:"updated_at"`
// 	// DeletedAt gorm.DeletedAt `json:"-"`
// }

// func (CategoryRespon) TableName() string {
// 	return "categories"
// }
