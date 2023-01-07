package controllers

import (
	"log"
	"net/http"

	"github.com/ezkahan/golab/configs"
	"github.com/ezkahan/golab/helpers"
	"github.com/ezkahan/golab/models"
	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "Invalid json data")
		return
	}

	result := configs.DB.First(&user)

	if result.Error != nil {
		log.Fatal("Error finding user")
	}

	// check user continue

	token, err := helpers.CreateToken(uint64(user.ID))

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}

func SignIn(ctx *gin.Context) {
	// var user models.User

	// return
}

func Logout(ctx *gin.Context) {
	// var user models.User

	// return
}
