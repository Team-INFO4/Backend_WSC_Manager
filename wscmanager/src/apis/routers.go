package wsc_apis

import (
	"github.com/gin-gonic/gin"
)

func APIs(routers *gin.Engine) {
	routers.POST("/api/auth/login", Login)
	routers.POST("/api/auth/signup", Signup)
	routers.POST("/api/data/crawl", NotionCrawl)
	routers.POST("/api/data/save", SaveDB)
	routers.POST("/api/data/write", WriteReport)
}
