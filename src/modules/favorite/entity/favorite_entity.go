package entity

import "gorm.io/gorm"

type Favorite struct {
	ID     uint64 `gorm:"primaryKey;unique;autoIncrement"`
	UserID uint64
	BookID uint64
	gorm.Model
}
