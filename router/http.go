package router

import (
	authHandler "github.com/ezkahan/book_store/src/handler/auth"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	router := gin.Default()
	router.SetTrustedProxies([]string{"10.10.11.130", "192.168.1.100"})
	// gin.SetMode(gin.ReleaseMode)

	authHandler.Routes(router)

	return router
}
