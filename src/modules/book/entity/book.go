package entity

import (
	imageEntity "github.com/ezkahan/book_store/src/modules/image/entity"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Slug        string  `gorm:"unique;size:255;not null" json:"slug"`
	Title       string  `gorm:"not null; size:255" json:"title"`
	Description string  `gorm:"text" json:"description"`
	Rating      float32 `gorm:"default:0.0" json:"rating"`
	Language    string  `gorm:"size:10;not null;default:'tm'" json:"language"`
	Year        int8    `gorm:"not null" json:"year"`
	TotalPages  int8    `gorm:"not null" json:"total_pages"`
	Preview     int32   `gorm:"default:0" json:"preview"`
	CategoryId  uint64
	AuthorId    uint64
	Images      []imageEntity.Image
}
