package service

import (
	"github.com/ezkahan/book_store/src/modules/user/dto"
	"github.com/ezkahan/book_store/src/modules/user/entity"
	"github.com/ezkahan/book_store/src/modules/user/repository"
)

type UserService interface {
	CreateUser(user dto.UserDTO) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) CreateUser(user dto.UserDTO) entity.User {
	return entity.User{}
}
