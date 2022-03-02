package wsc_apis

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func NotionCrawl(c *gin.Context) {
	AuthorizaionHeader := c.GetHeader("Authorizaion")

	if AuthorizaionHeader == "" || !strings.Contains(AuthorizaionHeader, "Bearer ") {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to get Authorization header.",
		})
		return
	}

	var Crawljson NotionCrawljson

	binderr := c.BindJSON(&Crawljson)
	if binderr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to get body.",
		})
		return
	}
}
