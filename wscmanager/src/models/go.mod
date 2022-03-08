module wsc_models

go 1.17

require (
	github.com/jinzhu/gorm v1.9.16
	github.com/joho/godotenv v1.4.0
	wscmanager.com/jsonstructs v0.0.0
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
)

replace wscmanager.com/jsonstructs v0.0.0 => ../json
