package request

type UserRegisterRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"note" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserUpdateRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Note     string `json:"note" form:"note"`
	Password string `json:"password"`
}
