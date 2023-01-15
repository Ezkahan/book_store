package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        uint64 `gorm:"unique"`
	Name      string `gorm:"string"`
	Email     string `gorm:"string"`
	Phone     string `gorm:"string"`
	Password  string `gorm:"string"`
	Image     string `gorm:"string"`
	Points    string `gorm:"string"`
	Role      string `gorm:"string"`
	Comments  string `gorm:"string"`
	Favorites string `gorm:"string"`
}
