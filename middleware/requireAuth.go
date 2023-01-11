package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func RequireAuth(ctx *gin.Context) {
	fmt.Println("I'm from middleware")

	ctx.Next()
}
