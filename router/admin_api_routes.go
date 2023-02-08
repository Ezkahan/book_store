package router

import (
	"github.com/ezkahan/book_store/src/handler/banner"
	"github.com/ezkahan/book_store/src/modules/banner/repository"
	"github.com/ezkahan/book_store/src/modules/banner/service"
	"github.com/gin-gonic/gin"
)

var (
	bannerRepository repository.BannerRepository = repository.NewBannerRepository(db)
	bannerService    service.BannerService       = service.NewBannerService(bannerRepository)
	bannerHandler    banner.BannerHandler        = banner.NewBannerHandler(bannerService)
)

func AdminApiRoutes(router *gin.Engine) {
	api := router.Group("/api/v1/admin")
	{
		api.GET("/", func(ctx *gin.Context) {})

		banner := api.Group("/banner")
		{
			banner.GET("/list", bannerHandler.GetBanners)
			banner.POST("/add", bannerHandler.Add)
			banner.DELETE("/delete/:id", bannerHandler.Delete)
		}
	}
}
