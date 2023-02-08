package router

import (
	"github.com/ezkahan/book_store/database/mysql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = mysql.Connection()
)

func InitRoutes() *gin.Engine {
	routes := gin.Default()
	routes.SetTrustedProxies([]string{"10.10.11.130", "192.168.1.100"})
	// gin.SetMode(gin.ReleaseMode)

	AuthRoutes(routes)
	ApiRoutes(routes)
	AdminApiRoutes(routes)

	return routes
}
