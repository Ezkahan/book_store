package repository

import (
	"errors"

	"github.com/ezkahan/book_store/src/modules/banner/entity"
	"gorm.io/gorm"
)

type BannerRepository interface {
	List() entity.BannerList
	Add(banner entity.Banner) entity.Banner
	Delete(id uint64) error
}

type bannerRepository struct {
	db *gorm.DB
}

func NewBannerRepository(db *gorm.DB) BannerRepository {
	return &bannerRepository{
		db,
	}
}

func (repo *bannerRepository) List() entity.BannerList {
	var banners entity.BannerList
	repo.db.Find(&banners)
	return banners
}

func (repo *bannerRepository) Add(banner entity.Banner) entity.Banner {
	repo.db.Create(&banner)
	return banner
}

func (repo *bannerRepository) Delete(id uint64) error {
	res := repo.db.Delete(&entity.Banner{}, id)
	if res.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}
