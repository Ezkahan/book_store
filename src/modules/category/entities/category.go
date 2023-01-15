package entities

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID          uint64 `gorm:"unique"`
	Slug        string `gorm:"string"`
	Title       string `gorm:"string"`
	Description string `gorm:"string"`
	Icon        string `gorm:"string"`
	Preview     string `gorm:"string"`
}
