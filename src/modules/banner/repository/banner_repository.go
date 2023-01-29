package repository

import (
	"github.com/ezkahan/book_store/src/modules/banner/entity"
	"gorm.io/gorm"
)

type IBannerRepository interface {
	Add(data string) entity.Banner
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

func (repo *BannerRepository) Add(data string) {
	//
}

func (repo *BannerRepository) Delete(id uint) {
	//
}
