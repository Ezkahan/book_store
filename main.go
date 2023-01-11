package main

import (
	"fmt"
	"os"

	"github.com/ezkahan/golab/configs"
	controller "github.com/ezkahan/golab/controllers"
	"github.com/ezkahan/golab/middleware"
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

	router.GET("/", middleware.RequireAuth, controller.Home)
	router.POST("/signin", controller.SignIn)
	router.POST("/signup", controller.SignUp)
	router.POST("/logout", controller.Logout)

	router.Run()

	fmt.Println("Server running on port: ", port)
}
