package entity

import (
	"gorm.io/gorm"
)

type Author struct {
	ID      uint64 `gorm:"unique"`
	Slug    string `gorm:"string"`
	Name    string `gorm:"string"`
	Image   string `gorm:"string"`
	Preview string `gorm:"string"`
	gorm.Model
}
