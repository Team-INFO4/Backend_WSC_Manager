package wsc_apis

import (
	"github.com/gin-gonic/gin"
	wsc_middleware "wscmanager.com/middleware"
)

func APIs(routers *gin.Engine) {
	routers.POST("/api/auth/login", Login)
	routers.POST("/api/auth/signup", Signup)

	data := routers.Group("/api/data")
	data.Use(wsc_middleware.JwtAuthMiddleware())
	{
		data.POST("/crawl", NotionCrawl)
		data.POST("/save", SaveDB)
		data.POST("/write", WriteReport)
	}

}
