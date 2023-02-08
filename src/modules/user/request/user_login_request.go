package request

type UserLoginRequest struct {
	Phone    string `json:"phone" form:"phone" validate:"required,min=11"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}
