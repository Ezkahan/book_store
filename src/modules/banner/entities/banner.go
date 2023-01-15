package entities

import (
	"gorm.io/gorm"
)

type Banner struct {
	gorm.Model
	ID        uint64 `gorm:"unique"`
	Image     string `gorm:"string"`
	Link      string `gorm:"string"`
	Active    string `gorm:"string"`
	StartDate string `gorm:"string"`
	EndDate   string `gorm:"string"`
}
