package app

import (
	"github.com/gvu0110/bookstore_users-api/controllers/ping"
	"github.com/gvu0110/bookstore_users-api/controllers/users"
)

func mapURLs() {
	// All handlers have to implement (c *gin.Context) interface
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
