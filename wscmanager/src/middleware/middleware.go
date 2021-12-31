package wsc_middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
func Test() {
	fmt.Println("Test_of_middleware")
}
*/

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-Content-Type-Options", "nosniff")

		if c.Request.Method == "PUT" || c.Request.Method == "DELETE" {
			c.AbortWithStatus(http.StatusMethodNotAllowed)
			return
		}

		c.Next()
	}
}
