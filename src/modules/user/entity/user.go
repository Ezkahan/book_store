package entity

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	commentEntity "github.com/ezkahan/book_store/src/modules/comment/entity"
	favoriteEntity "github.com/ezkahan/book_store/src/modules/favorite/entity"
)

type User struct {
	ID        uint64                    `gorm:"unique;primaryKey;autoIncrement" json:"id"`
	Name      string                    `gorm:"size:25;not null;" json:"name"`
	Phone     string                    `gorm:"size:16;not null;unique" json:"phone"`
	Password  string                    `gorm:"size:128;not null" json:"-"`
	Email     string                    `gorm:"size:50;" json:"email"`
	Image     string                    `gorm:"size:100" json:"image,omitempty"`
	Points    int16                     `gorm:"not null;default:0" json:"points"`
	Role      string                    `gorm:"default:'user'" json:"role"`
	Comments  []commentEntity.Comment   `gorm:"string" json:"comments"`
	Favorites []favoriteEntity.Favorite `gorm:"string" json:"favorites"`
	Token     string                    `gorm:"-" json:"token"`
	gorm.Model
}

type UserList []User

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 8)

	if err != nil {
		return errors.New("error hashing password")
	}

	u.Password = string(hash)

	return
}
