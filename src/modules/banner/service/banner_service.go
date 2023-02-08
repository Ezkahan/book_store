package service

import (
	"log"
	"path/filepath"
	"time"

	"github.com/ezkahan/book_store/src/modules/banner/entity"
	"github.com/ezkahan/book_store/src/modules/banner/repository"
	"github.com/ezkahan/book_store/src/modules/banner/request"
	"github.com/gin-gonic/gin"
)

type BannerService interface {
	List() entity.BannerList
	Add(ctx *gin.Context, request request.AddBannerRequest) entity.Banner
	Delete(id uint64) error
}

type bannerService struct {
	bannerRepository repository.BannerRepository
}

func NewBannerService(bannerRepo repository.BannerRepository) BannerService {
	return &bannerService{
		bannerRepository: bannerRepo,
	}
}

func (service *bannerService) List() entity.BannerList {
	res := service.bannerRepository.List()
	return res
}

func (service *bannerService) Add(ctx *gin.Context, request request.AddBannerRequest) entity.Banner {
	image_url := service.SaveImage(ctx)

	start, err := time.Parse("2006-01-02", request.StartDate)
	if err != nil {
		log.Fatal(err)
	}

	end, err := time.Parse("2006-01-02", request.EndDate)
	if err != nil {
		log.Fatal(err)
	}

	banner := entity.Banner{
		Image:     image_url,
		Link:      request.Link,
		Active:    request.Active,
		StartDate: start,
		EndDate:   end,
	}

	return service.bannerRepository.Add(banner)
}

func (service *bannerService) SaveImage(ctx *gin.Context) string {
	file, _ := ctx.FormFile("image")
	name := filepath.Base(file.Filename)
	ctx.SaveUploadedFile(file, "assets/images/banners/"+name)

	return file.Filename
}

func (service *bannerService) Delete(id uint64) error {
	err := service.bannerRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
