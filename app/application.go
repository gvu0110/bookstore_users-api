package app

import (
	"github.com/gin-gonic/gin"
)

// Only app and controller layers are defining and using HTTP framework, may be changed when deploying a new framework
var (
	router = gin.Default()
)

// StartApplication function start the web application
func StartApplication() {
	mapURLs()
	router.Run(":8080")
}
