package entities

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID          uint64 `gorm:"unique"`
	Slug        string `gorm:"string"`
	Title       string `gorm:"string"`
	Description string `gorm:"string"`
	Rating      string `gorm:"string"`
	Language    string `gorm:"string"`
	Year        string `gorm:"string"`
	Pages       string `gorm:"string"`
	Preview     string `gorm:"string"`
	CategoryId  string `gorm:"string"`
	AuthorId    string `gorm:"string"`
	Images      string `gorm:"string"`
}
