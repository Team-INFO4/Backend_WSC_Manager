package wsc_apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	wsc_jsonstructs "wscmanager.com/jsonstructs"
)

func Login(c *gin.Context) {
	var input wsc_jsonstructs.Loginjson

	//유효성 검사
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing id or password"})
		return
	}
}
