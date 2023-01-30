package dto

type UserDTO struct {
	Email string `json:"email" binding:"required"`
}
