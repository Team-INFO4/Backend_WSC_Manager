package wsc_models

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDB() {
	//.env 파일 읽기
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbconn := os.Getenv("DBCONN")
	db, err := gorm.Open("mysql", dbconn)
	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = db

	defer DB.Close()
}
