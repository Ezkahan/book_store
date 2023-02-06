package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		})
	}
}
