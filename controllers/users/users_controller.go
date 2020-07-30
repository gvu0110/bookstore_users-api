package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gvu0110/bookstore_users-api/domain/users"
	"github.com/gvu0110/bookstore_users-api/services"
	"github.com/gvu0110/bookstore_users-api/utils/errors"
)

// Entry points of the appication

// CreateUser function creates new users
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestRESTError("Invalid JSON Body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	result, err := services.CreateUser(user)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, result)
}

// GetUser function gets user's information
func GetUser(c *gin.Context) {
	userID, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestRESTError("Invalid User ID - User ID should be a number")
		c.JSON(err.StatusCode, err)
		return
	}

	user, err := services.GetUser(userID)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
