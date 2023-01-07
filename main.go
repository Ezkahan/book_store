package main

import (
	"fmt"
	"os"

	"github.com/ezkahan/golab/configs"
	authController "github.com/ezkahan/golab/controllers"
	"github.com/gin-gonic/gin"
)

func init() {
	configs.LoadEnv()
	configs.ConnectToDB()
	configs.SyncDatabase()
}

func main() {
	port := os.Getenv("PORT")
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello",
		})
	})

	router.POST("/signin", authController.SignIn)
	router.POST("/signup", authController.SignUp)
	router.POST("/logout", authController.Logout)

	router.Run()

	fmt.Println("Server running on port: ", port)
}
