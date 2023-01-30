package entity

import (
	commentEntity "github.com/ezkahan/book_store/src/modules/comment/entity"
	favoriteEntity "github.com/ezkahan/book_store/src/modules/favorite/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint64                    `gorm:"unique;primaryKey;autoIncrement" json:"id"`
	Name      string                    `gorm:"size:25;not null;" json:"name"`
	Email     string                    `gorm:"size:50;not null;unique" json:"email"`
	Phone     string                    `gorm:"size:16;not null;unique" json:"phone"`
	Password  string                    `gorm:"size:50;no null" json:"-"`
	Image     string                    `gorm:"size:100" json:"image,omitempty"`
	Points    string                    `gorm:"not null;default:0" json:"points"`
	Role      string                    `gorm:"default:'user'"`
	Comments  []commentEntity.Comment   `gorm:"string"`
	Favorites []favoriteEntity.Favorite `gorm:"string"`
}

type UserList []User

func (u *User) BeforeCreate() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 16)

	if err != nil {
		return err
	}

	u.Password = string(hash)

	return nil
}

// hooks
// BeforeSave
// BeforeCreate
// AfterSave
// AfterCreate
// ...
