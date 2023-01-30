package repository

import (
	"gorm.io/gorm"
)

type IBannerRepository interface {
	Add(data string) error
	Delete(id uint) error
}

type BannerRepository struct {
	db *gorm.DB
}

func NewBannerRepository(db *gorm.DB) *BannerRepository {
	return &BannerRepository{
		db: db,
	}
}

func (repo *BannerRepository) Add(banner string) {
	// return repo.db.Create(banner)
}

func (repo *BannerRepository) Delete(id uint) {
	//
}
