package middleware

import (
	"net/http"
	"rates/internal/app/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {

	config := cors.DefaultConfig()
	config.AddAllowHeaders("Authorization")
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	return cors.New(config)

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
