module src

go 1.13

require (
	github.com/gin-gonic/gin v1.7.7
	wscmanager.com/apis v0.0.0
	wscmanager.com/jsonstructs v0.0.0
	wscmanager.com/middleware v0.0.0
	wscmanager.com/staticfiles v0.0.0
)

replace (
	wscmanager.com/apis v0.0.0 => ./apis
	wscmanager.com/jsonstructs v0.0.0 => ./json
	wscmanager.com/middleware v0.0.0 => ./middleware
	wscmanager.com/staticfiles v0.0.0 => ./static
)
