package entities

import (
	"gorm.io/gorm"
)

type Favorite struct {
	gorm.Model
	ID     uint64 `gorm:"unique"`
	UserID uint64 `gorm:"unique"`
	BookID uint64 `gorm:"unique"`
}
