package wsc_apis

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	wsc_jsonstructs "wscmanager.com/jsonstructs"
	wsc_models "wscmanager.com/models"
)

func Signup(c *gin.Context) {
	var input wsc_jsonstructs.Signupjson
	//유효성 검사
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing id or password"})
		return
	}
	// key 체크
	godotenv.Load(".env")

	secret := os.Getenv("KEY")
	if key := input.Key; key != secret {
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid key value."})
		return
	}
	//check whitespace
	if strings.Contains(input.Id, " ") || strings.Contains(input.Password, " ") {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No Write Space"})
		return
	}

	//DB 저장
	user := wsc_models.User{}
	user.Id = input.Id
	user.Password = input.Password
	_, err := user.SaveUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{})

}
