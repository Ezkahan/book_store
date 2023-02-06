package request

type UserLoginRequest struct {
	Phone    string `validate:"required|min_len:12"`
	Password string `validate:"required|min_len:8"`
}
