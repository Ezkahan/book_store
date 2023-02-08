package request

type UserCreateRequest struct {
	Name     string `validate:"required,min=3"`
	Phone    string `validate:"required,min=11"`
	Password string `validate:"required,min=8"`
}
