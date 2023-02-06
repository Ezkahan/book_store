package entity

import (
	"time"

	"gorm.io/gorm"
)

type Banner struct {
	ID        uint64    `gorm:"primaryKey;unique;autoIncrement" json:"id"`
	Image     string    `gorm:"size:100" json:"image,omitempty"`
	Link      string    `gorm:"size:100" json:"link,omitempty"`
	Active    bool      `gorm:"default:false"`
	StartDate time.Time `gorm:"not null;autoCreateTime" json:"start_date"`
	EndDate   time.Time `gorm:"not null" json:"end_date"`
	gorm.Model
}
