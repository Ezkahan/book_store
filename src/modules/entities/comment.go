package entities

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID        uint64 `gorm:"unique"`
	UserID    uint64 `gorm:"unique"`
	Message   string `gorm:"string"`
	Confirmed string `gorm:"string"`
}
