package repository

import (
	"github.com/ezkahan/book_store/src/modules/user/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	VerifyCredential(email string, password string) interface{}
	GetUser(id uint64) entity.User
	IsDuplicatePhone(phone string) *gorm.DB
	DeleteUser(id uint64) *gorm.DB
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
	userRepo.db.Create(&user)
	return user
}

func (userRepo *userRepository) UpdateUser(user entity.User) entity.User {
	userRepo.db.Save(&user)
	return user
}

func (userRepo *userRepository) GetUser(id uint64) entity.User {
	var user entity.User
	userRepo.db.First(&user, id)
	return user
}

func (userRepo *userRepository) VerifyCredential(phone string, password string) interface{} {
	var user entity.User
	res := userRepo.db.Where("phone = ?", phone).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (userRepo *userRepository) IsDuplicatePhone(phone string) *gorm.DB {
	var user entity.User
	return userRepo.db.Where("phone = ?", phone).Take(&user)
}

func (userRepo *userRepository) DeleteUser(id uint64) *gorm.DB {
	return userRepo.db.Delete(&entity.User{}, id)
}
