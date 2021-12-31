package wsc_apis

import "github.com/gin-gonic/gin"

func APIs(api *gin.Engine) {
	Ping(api)
}
