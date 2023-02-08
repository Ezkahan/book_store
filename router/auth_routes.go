package router

import (
	"github.com/ezkahan/book_store/src/handler/auth"
	"github.com/ezkahan/book_store/src/modules/user/repository"
	"github.com/ezkahan/book_store/src/modules/user/service"
	"github.com/gin-gonic/gin"
)

var (
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	userService    service.UserService       = service.NewUserService(userRepository)
	authHandler    auth.AuthHandler          = auth.NewAuthHandler(userService)
)

func AuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/signin", authHandler.SignIn)
		auth.POST("/signup", authHandler.SignUp)
	}
}
