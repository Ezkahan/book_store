package entity

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	ID        uint64 `gorm:"primaryKey;unique;autoIncrement"`
	UserID    uint64
	Message   string `gorm:"text"`
	Confirmed string `gorm:"not null;default:0" json:"confirmed"`
}
