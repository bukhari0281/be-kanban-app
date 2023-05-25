package request

type CategoryCreateRequest struct {
	Category string `json:"category" form:"category"`
}
