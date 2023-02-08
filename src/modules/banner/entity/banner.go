package entity

import (
	"database/sql"
	"time"
)

type Banner struct {
	ID        uint64       `gorm:"primaryKey;unique;autoIncrement" json:"id"`
	Image     string       `gorm:"size:100" json:"image,omitempty"`
	Link      string       `gorm:"size:100" json:"link,omitempty"`
	Active    bool         `gorm:"default:false" json:"active"`
	StartDate time.Time    `gorm:"not null;autoCreateTime" json:"start_date"`
	EndDate   time.Time    `gorm:"not null" json:"end_date"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt sql.NullTime `gorm:"index" json:"-"`
}

type BannerList []Banner
