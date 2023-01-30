package repository

import (
	"github.com/ezkahan/book_store/src/modules/user/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	GetUser(id uint64) entity.User
	DeleteUser(id uint64) string
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db,
	}
}

func (userRepo *userRepository) InsertUser(user entity.User) entity.User {
	userRepo.db.Create(user)

	return entity.User{}
}

func (userRepo *userRepository) UpdateUser(user entity.User) entity.User {
	return entity.User{}
}

func (userRepo *userRepository) GetUser(id uint64) entity.User {
	return entity.User{}
}

func (userRepo *userRepository) DeleteUser(id uint64) string {
	return ""
}
