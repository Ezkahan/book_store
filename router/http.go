package router

import (
	middleware "github.com/ezkahan/golab/middlewares"
	"github.com/gin-gonic/gin"

	authHandler "github.com/ezkahan/golab/src/handlers/auth"
	homeHandler "github.com/ezkahan/golab/src/handlers/home"
)

func Routes() *gin.Engine {
	routes := gin.Default()

	routes.GET("/", middleware.AuthMiddleware, homeHandler.Index)
	routes.POST("/signin", authHandler.SignIn)
	routes.POST("/signup", authHandler.SignUp)
	routes.POST("/logout", authHandler.Logout)

	return routes

}
