package main

import (
	"github.com/gin-gonic/gin"
	wsc_apis "wscmanager.com/apis"
	wsc_middleware "wscmanager.com/middleware" // 미들웨어 모듈
	wsc_models "wscmanager.com/models"
)

func main() {
	engine := gin.Default()
	wsc_models.ConnectDB()
	engine.Use(wsc_middleware.Middleware()) // 미들웨어
	wsc_apis.APIs(engine)                   // APIs

	engine.Run(":65530")
}
