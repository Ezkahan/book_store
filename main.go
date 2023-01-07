package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ezkahan/golab/configs"
	"github.com/ezkahan/golab/models"
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
		user := models.User{Email: "ezkahan.dev@gmail.com", Password: "test123"}

		result := configs.DB.Create(&user)

		if result.Error != nil {
			log.Fatal("Error: ", result.Error)
		}

		ctx.JSON(200, gin.H{
			"message": "Hello",
		})
	})

	router.Run()

	fmt.Println("Server running on port: ", port)
}
