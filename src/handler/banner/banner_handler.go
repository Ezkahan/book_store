package banner

import (
	"github.com/ezkahan/book_store/src/modules/banner/repository"
)

type IBannerHandler interface {
	Add(data string) error
	Delete(id uint) error
}

type BannerHandler struct {
	repo *repository.BannerRepository
}

func NewBannerHandler(bannerRepository *repository.BannerRepository) *BannerHandler {
	return &BannerHandler{
		repo: bannerRepository,
	}
}

func (b *BannerHandler) Add(data string) {
	b.repo.Add(data)
}

func (b *BannerHandler) Delete(id uint) {
	b.repo.Delete(id)
}
