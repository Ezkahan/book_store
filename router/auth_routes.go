package router

import (
	"github.com/ezkahan/book_store/database/mysql"
	"github.com/ezkahan/book_store/src/handler/auth"
	"github.com/ezkahan/book_store/src/modules/user/repository"
	"github.com/ezkahan/book_store/src/modules/user/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                  = mysql.Connection()
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
