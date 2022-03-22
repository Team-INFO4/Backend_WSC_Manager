package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	wsc_apis "wscmanager.com/apis"
	wsc_middleware "wscmanager.com/middleware" // 미들웨어 모듈
	wsc_models "wscmanager.com/models"
)

func main() {
	//.env 파일 읽기
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	engine := gin.Default()
	wsc_models.ConnectDB()

	data := engine.Group("/api/data")
	engine.Use(wsc_middleware.Middleware())      // 미들웨어
	data.Use(wsc_middleware.JwtAuthMiddleware()) // /api/datt/[] 미들웨어
	wsc_apis.APIs(engine)                        // APIs

	engine.Run(":65530")

	defer wsc_models.DB.Close()
}
