package entity

import "time"

type Category struct {
	ID          uint64 `gorm:"primaryKey;unique;autoIncrement"`
	Slug        string `gorm:"unique;size:255;not null" json:"slug"`
	Title       string `gorm:"size:255;not null" json:"title"`
	Description string `gorm:"text" json:"description"`
	Icon        string `gorm:"size:255" json:"icon,omitempty"`
	Preview     int    `gorm:"default:0" json:"preview"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type CategoryList []Category
