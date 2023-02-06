package service

import (
	"github.com/ezkahan/book_store/src/modules/user/entity"
	"github.com/ezkahan/book_store/src/modules/user/repository"
	"github.com/ezkahan/book_store/src/modules/user/request"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(req request.UserCreateRequest) entity.User
	IsDuplicatePhone(email string) bool
	VerifyCredential(phone string, password string) interface{}
	GetUser(id uint64) entity.User
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) CreateUser(req request.UserCreateRequest) entity.User {
	user := entity.User{
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
	}
	return service.userRepository.InsertUser(user)
}

func (service *userService) IsDuplicatePhone(phone string) bool {
	res := service.userRepository.IsDuplicatePhone(phone)
	return !(res.Error == nil)
}

func (service *userService) GetUser(id uint64) entity.User {
	return service.userRepository.GetUser(id)
}

func (service *userService) VerifyCredential(phone string, password string) interface{} {
	res := service.userRepository.VerifyCredential(phone, password)
	if v, ok := res.(entity.User); ok {
		comparedPass := service.CompareHashAndPassword(v.Password, []byte(password))
		if v.Phone == phone && comparedPass {
			return res
		}
		return false
	}
	return false
}

func (service *userService) CompareHashAndPassword(hashed string, plain []byte) bool {
	byteHash := []byte(hashed)

	if err := bcrypt.CompareHashAndPassword(byteHash, plain); err != nil {
		return false
	}

	return true
}
