package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping function confirm the web application is working
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
