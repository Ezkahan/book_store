package controllers

import (
	"net/http"

	"github.com/ezkahan/golab/configs"
	"github.com/ezkahan/golab/helpers"
	"github.com/ezkahan/golab/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(ctx *gin.Context) {
	var body models.User

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 16)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})

		return
	}

	user := models.User{
		Email:    body.Email,
		Password: string(hash),
	}

	result := configs.DB.Create(&user)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})

		return
	}

	token, err := helpers.CreateToken(uint64(user.ID))

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

func SignIn(ctx *gin.Context) {
	var body models.User

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	var user models.User
	configs.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	token, err := helpers.CreateToken(uint64(user.ID))

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

	// return
}

func Logout(ctx *gin.Context) {
	// var user models.User

	// return
}
