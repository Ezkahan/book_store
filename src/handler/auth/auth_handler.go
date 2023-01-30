package auth

import (
	"net/http"

	// "github.com/ezkahan/book_store/database/mysql"
	"github.com/ezkahan/book_store/src/modules/user/entity"
	"github.com/ezkahan/book_store/src/modules/user/service"
	"github.com/ezkahan/book_store/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler interface {
	SignUp(ctx *gin.Context)
	SignIn(ctx *gin.Context)
	Logout(ctx *gin.Context)
}

type authHandler struct {
	userService service.UserService
}

func NewAuthHandler(userService service.UserService) AuthHandler {
	return &authHandler{
		userService,
	}
}

func (h *authHandler) SignUp(ctx *gin.Context) {
	var body entity.User

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	user := entity.User{
		Email:    body.Email,
		Password: body.Password,
	}

	// result := mysql.DB.Create(&user)

	// if result.Error != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Failed to create user",
	// 	})

	// 	return
	// }

	token, err := utils.GenerateToken(uint64(user.ID))

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": "Failed to create token",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}

func (h *authHandler) SignIn(ctx *gin.Context) {
	var body entity.User

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	var user entity.User
	// mysql.DB.First(&user, "email = ?", body.Email)

	// if user.ID == 0 {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Invalid email or password",
	// 	})

	// 	return
	// }

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	token, err := utils.GenerateToken(uint64(user.ID))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}

func (h *authHandler) Logout(ctx *gin.Context) {
	// return
}
