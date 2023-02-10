package entity

import "time"

type Author struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;unique" json:"id"`
	Slug      string `gorm:"unique;size:255,not null" json:"slug"`
	Name      string `gorm:"size:25" json:"name"`
	Image     string `gorm:"string" json:"image"`
	Preview   int    `gorm:"default:0" json:"preview"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AuthorList []Author
