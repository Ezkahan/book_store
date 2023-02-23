package request

import "mime/multipart"

type CategoryRequest struct {
	Icon        *multipart.FileHeader `json:"icon" form:"icon" validate:"required,file"`
	Title       string                `json:"title" form:"title" validate:"min=3,max=256"`
	Description string                `json:"description" form:"description" validate:"required"`
}
