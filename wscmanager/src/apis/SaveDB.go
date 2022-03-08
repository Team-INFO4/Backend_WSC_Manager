package wsc_apis

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	wsc_jsonstructs "wscmanager.com/jsonstructs"
)

func SaveDB(c *gin.Context) {
	AuthorizaionHeader := c.GetHeader("Authorizaion")

	if AuthorizaionHeader == "" || !strings.Contains(AuthorizaionHeader, "Bearer secret_") {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to get Authorization header.",
		})
		return
	}

	var Crawljson wsc_jsonstructs.SaveDBjson

	binderr := c.BindJSON(&Crawljson)
	if binderr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to get body.",
		})
		return
	}

	if Crawljson.StartDate == "" || Crawljson.EndDate == "" || Crawljson.Target == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Insufficient parameters.",
		})
		return
	}
}
