package router

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{"10.10.11.130", "192.168.1.100"})
	// gin.SetMode(gin.ReleaseMode)

	AuthRoutes(r)
	ApiRoutes(r)
	AdminRoutes(r)

	return r
}
