package wsc_middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	wsc_models "wscmanager.com/models"
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
		id, err := token.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err})
			c.Abort()
			return
		}
		_, err = wsc_models.GetUserByID(id)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "incorrect content token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
