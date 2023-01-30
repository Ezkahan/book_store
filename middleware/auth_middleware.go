package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	// authorization := ctx.Request.Header.Get("Authorization")
	authorization := ctx.GetHeader("Authorization")
	token := strings.Split(authorization, "Bearer")

	if len(token) > 0 && token[1] == "" {
		ctx.Abort()
		return
	}

	ctx.Next()
}
