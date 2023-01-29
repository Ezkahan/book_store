package router

import (
	middleware "github.com/ezkahan/book_store/middlewares"
	authHandler "github.com/ezkahan/book_store/src/handlers/auth"
	homeHandler "github.com/ezkahan/book_store/src/handlers/home"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	routes := gin.Default()
	routes.SetTrustedProxies([]string{"10.10.11.130"})
	gin.SetMode(gin.ReleaseMode)

	routes.GET("/", middleware.AuthMiddleware, homeHandler.Index)
	routes.POST("/signin", authHandler.SignIn)
	routes.POST("/signup", authHandler.SignUp)
	routes.POST("/logout", authHandler.Logout)

	return routes
}
