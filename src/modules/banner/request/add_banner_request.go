package request

import "mime/multipart"

type AddBannerRequest struct {
	Image     *multipart.FileHeader `json:"image" form:"image" validate:"required,file"`
	Link      string                `json:"link" form:"link" validate:"min=12,max=256"`
	Active    bool                  `json:"active" form:"active" validate:"required"`
	StartDate string                `json:"start_date" form:"start_date" validate:"required,datetime=2006-01-02"`
	EndDate   string                `json:"end_date" form:"end_date" validate:"required,datetime=2006-01-02"`
}
