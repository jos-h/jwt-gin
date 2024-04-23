package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jwt-gin/utils"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := utils.TokenValid(ctx)
		if err != nil {
			ctx.String(http.StatusUnauthorized, "Unauthorized")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
