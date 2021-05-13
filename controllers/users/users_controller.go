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

func getUserID(userIDParams string) (int64, *errors.RESTError) {
	userID, err := strconv.ParseInt(userIDParams, 10, 64)
	if err != nil {
		return 0, errors.NewBadRequestRESTError("User ID should be a number")
	}
	return userID, nil
}

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
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

// GetUser function gets user's information
func GetUser(c *gin.Context) {
	userID, err := getUserID(c.Param("user_id"))
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	user, err := services.GetUser(userID)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}

func UpdateUser(c *gin.Context) {
	userID, err := getUserID(c.Param("user_id"))
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestRESTError("Invalid JSON Body")
		c.JSON(restErr.StatusCode, restErr)
		return
	}

	user.ID = userID
	result, err := services.UpdateUser(user)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func DeleteUser(c *gin.Context) {
	userID, err := getUserID(c.Param("user_id"))
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	if err := services.DeleteUser(userID); err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.FindUsersByStatus(status)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, users.Marshall((c.GetHeader("X-Public") == "true")))
}
