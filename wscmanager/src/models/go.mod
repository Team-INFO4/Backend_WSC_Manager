module wsc_models

go 1.17

require (
	github.com/jinzhu/gorm v1.9.16
	golang.org/x/crypto v0.0.0-20191205180655-e7c4368fe9dd
	wscmanager.com/utils v0.0.0
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
)

replace wscmanager.com/jsonstructs v0.0.0 => ../json

replace wscmanager.com/utils v0.0.0 => ../utils
