package entity

import (
	"time"

	"gorm.io/gorm"
)

type Banner struct {
	gorm.Model
	ID        uint64    `gorm:"primaryKey;unique;autoIncrement" json:"id"`
	Image     string    `gorm:"size:100" json:"image,omitempty"`
	Link      string    `gorm:"size:100" json:"link,omitempty"`
	Active    bool      `gorm:"default:false"`
	StartDate time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"start_date"`
	EndDate   time.Time `gorm:"not null" json:"end_date"`
}
