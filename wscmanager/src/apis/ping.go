package wsc_apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(api *gin.Engine) {
	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
