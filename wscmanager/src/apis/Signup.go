package wsc_apis

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	wsc_jsonstructs "wscmanager.com/jsonstructs"
)

func Signup(c *gin.Context) {
	var input wsc_jsonstructs.Signupjson
	//유효성 검사
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing id or password"})
		return
	}
	// key 체크
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secret := os.Getenv("KEY")
	if key := input.Key; key != secret {
		c.JSON(http.StatusForbidden, gin.H{"message": "Invalid key value."})
		return
	}
}
