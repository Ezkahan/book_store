package auth

import (
	"github.com/ezkahan/book_store/database/mysql"
	"github.com/ezkahan/book_store/src/modules/user/repository"
	"github.com/ezkahan/book_store/src/modules/user/service"
	"github.com/gin-gonic/gin"
)

var (
	userRepository repository.UserRepository = repository.NewUserRepository(mysql.DB)
	userService    service.UserService       = service.NewUserService(userRepository)
	handler        AuthHandler               = NewAuthHandler(userService)
)

func Routes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/signin", handler.SignIn)
		auth.POST("/signup", handler.SignUp)
		auth.POST("/logout", handler.Logout)
	}
}
