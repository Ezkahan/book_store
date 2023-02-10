package request

import "mime/multipart"

type AddAuthorRequest struct {
	Slug  string                `validate:"required,min=5,max=255"`
	Name  string                `validate:"required,min=5,max=255"`
	Image *multipart.FileHeader `validate:"required,file" json:"image"`
}
