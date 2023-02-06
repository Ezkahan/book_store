package entity

import "gorm.io/gorm"

type Comment struct {
	ID        uint64 `gorm:"primaryKey;unique;autoIncrement"`
	UserID    uint64
	Message   string `gorm:"text"`
	Confirmed string `gorm:"not null;default:0" json:"confirmed"`
	gorm.Model
}
