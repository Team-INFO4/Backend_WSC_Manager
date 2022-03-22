package wsc_middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"wscmanager.com/utils/token"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-Content-Type-Options", "nosniff")

		c.Next()
	}
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := token.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{})
			c.Abort()
			return
		}

		c.Next()
	}
}
