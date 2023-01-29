package entity

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	ID     uint64 `gorm:"primaryKey;unique;autoIncrement"`
	BookID uint64
	Image  string `gorm:"size:255" json:"image,omitempty"`
	Type   string `gorm:"size:10" json:"type"`
}
