package auth

import (
	"net/http"

	"github.com/ezkahan/book_store/src/modules/user/entity"
	"github.com/ezkahan/book_store/src/modules/user/request"
	"github.com/ezkahan/book_store/src/modules/user/service"
	"github.com/ezkahan/book_store/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type AuthHandler interface {
	SignUp(ctx *gin.Context)
	SignIn(ctx *gin.Context)
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
	var (
		request  request.UserCreateRequest
		validate *validator.Validate = validator.New()
	)

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
	}

	err := validate.Struct(request)
	if err != nil {
		errors := utils.ParseValidationError(err)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})
		return
	}

	if !h.userService.IsDuplicatePhone(request.Phone) {
		ctx.JSON(http.StatusOK, gin.H{
			"error": "Duplicate phone",
		})
	} else {
		createdUser := h.userService.CreateUser(request)
		token, err := utils.GenerateToken(createdUser.ID)
		createdUser.Token = token

		if err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": "Failed to generate token",
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"user": createdUser,
		})
	}
}

func (h *authHandler) SignIn(ctx *gin.Context) {
	var (
		request  request.UserLoginRequest = request.UserLoginRequest{}
		validate *validator.Validate      = validator.New()
	)

	if err := ctx.ShouldBindBodyWith(&request, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	err := validate.Struct(request)
	if err != nil {
		errors := utils.ParseValidationError(err)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": errors,
		})
		return
	}

	authResult := h.userService.VerifyCredential(request.Phone, request.Password)
	if user, ok := authResult.(entity.User); ok {
		generatedToken, err := utils.GenerateToken(user.ID)
		user.Token = generatedToken

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"user": "Failed to generate token",
			})
		}

		ctx.JSON(http.StatusOK, gin.H{
			"user": user,
		})
		return
	}

	ctx.JSON(http.StatusBadRequest, gin.H{
		"message": "Phone or password invalid",
	})
}
