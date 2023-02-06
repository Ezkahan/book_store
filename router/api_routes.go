package router

import "github.com/gin-gonic/gin"

func ApiRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/books", func(ctx *gin.Context) {})
		api.GET("/book/:slug", func(ctx *gin.Context) {})
		api.GET("/authors", func(ctx *gin.Context) {})
		api.GET("/author/:slug", func(ctx *gin.Context) {})
		api.GET("/categories", func(ctx *gin.Context) {})
		api.GET("/category/:slug", func(ctx *gin.Context) {})
		api.GET("/banners", func(ctx *gin.Context) {})
		api.GET("/users", func(ctx *gin.Context) {})
	}
}
