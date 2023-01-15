package entities

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	ID     uint64 `gorm:"unique"`
	BookID uint64 `gorm:"unique"`
	Image  string `gorm:"string"`
	Type   string `gorm:"string"`
}
