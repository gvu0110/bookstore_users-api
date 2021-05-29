package users

import (
	"regexp"
	"strings"

	"github.com/gvu0110/bookstore_utils-go/rest_errors"
)

// Data Transfer Object: object transferred between the database and the application.
// Core domain microservice API

const (
	StatusActive = "active"
)

// User struct provides and exposes user entity
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

type Users []User

// Validate function validates a User struct
func (user *User) Validate() rest_errors.RESTError {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if (len(user.Email) < 3 && len(user.Email) > 254) || !emailRegex.MatchString(user.Email) {
		return rest_errors.NewBadRequestRESTError("Invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return rest_errors.NewBadRequestRESTError("Invalid password")
	}
	return nil
}
