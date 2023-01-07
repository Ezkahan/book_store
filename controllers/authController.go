package controllers

import (
	"net/http"

	"github.com/ezkahan/golab/helpers"
	"github.com/ezkahan/golab/models"
	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJson(&user); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, "Invalid json data")
		return
	}

	// check user

	token, err := helpers.CreateToken(uint64(user.ID))

	if err != nil {
		ctx.JSON(http.StatusunprocessableEntity, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": token,
	})
}

func SignIn(ctx *gin.Context) {
	var user models.User

	return
}
