package wsc_apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	wsc_jsonstructs "wscmanager.com/jsonstructs"
	wsc_models "wscmanager.com/models"
)

func Login(c *gin.Context) {
	var input wsc_jsonstructs.Loginjson

	//유효성 검사
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing id or password"})
		return
	}

	//id, pw
	u := wsc_models.User{}

	u.Id = input.Id
	u.Password = input.Password
	//토큰 발급
	token, err := wsc_models.LoginCheck(u.Id, u.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"accessToken": token})
}
