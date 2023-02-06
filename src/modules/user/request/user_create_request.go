package request

type UserCreateRequest struct {
	Name     string `validate:"required" binding:"required"`
	Phone    string `validate:"required|min_len:12"`
	Password string `validate:"required|min_len:8"`
}
