package wsc_models

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func ConnectDB() {
	dbconn := os.Getenv("DBCONN")
	db, err := gorm.Open("mysql", dbconn)
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.SingularTable(true)
	DB = db

}
