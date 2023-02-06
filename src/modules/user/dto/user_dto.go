package dto

type UserDTO struct {
	ID     uint64 `json:"id"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Points int16  `json:"points"`
	Role   string `json:"role"`
	Token  string `json:"token"`
}
