package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	fmt.Println("I'm from middleware")

	ctx.Next()
}
