package router

import (
	"github.com/ezkahan/book_store/src/handler/author"
	"github.com/ezkahan/book_store/src/handler/banner"
	authorRepo "github.com/ezkahan/book_store/src/modules/author/repository"
	authorServ "github.com/ezkahan/book_store/src/modules/author/service"
	bannerRepo "github.com/ezkahan/book_store/src/modules/banner/repository"
	bannerServ "github.com/ezkahan/book_store/src/modules/banner/service"
	"github.com/gin-gonic/gin"
)

var (
	bannerRepository bannerRepo.BannerRepository = bannerRepo.NewBannerRepository(db)
	bannerService    bannerServ.BannerService    = bannerServ.NewBannerService(bannerRepository)
	bannerHandler    banner.BannerHandler        = banner.NewBannerHandler(bannerService)

	authorRepository authorRepo.AuthorRepository = authorRepo.NewAuthorRepository(db)
	authorService    authorServ.AuthorService    = authorServ.NewAuthorService(authorRepository)
	authorHandler    author.AuthorHandler        = author.NewAuthorHandler(authorService)
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

		author := api.Group("/author")
		{
			author.POST("/add", authorHandler.Add)
			author.POST("/update/:slug", authorHandler.Update)
			author.GET("/:slug", authorHandler.Get)
			author.GET("/list", authorHandler.List)
			author.DELETE("/delete/:slug", authorHandler.Delete)
		}
	}
}
