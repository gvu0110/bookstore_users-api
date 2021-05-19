package app

import (
	"github.com/gin-gonic/gin"
	"github.com/gvu0110/bookstore_users-api/logger"
)

// Only app and controller layers are defining and using HTTP framework, may be changed when deploying a new framework
var (
	router = gin.Default()
)

// StartApplication function start the web application
func StartApplication() {
	mapURLs()
	logger.Info("Starting the application ...")
	router.Run(":8081")
}
