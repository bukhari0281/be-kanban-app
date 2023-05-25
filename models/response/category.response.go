package response

type CategoryRespon struct {
	ID       uint   `json:"id"`
	Category string `json:"category" `
	// CreatedAt time.Time `json:"created_at"`
	// UpdatedAt time.Time `json:"updated_at"`
	// DeletedAt gorm.DeletedAt `json:"-"`
}

func (CategoryRespon) TableName() string {
	return "categories"
}
