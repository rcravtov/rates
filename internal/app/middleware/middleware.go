package middleware

import (
	"net/http"
	"rates/internal/app/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {

	return cors.Default()

}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := service.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
