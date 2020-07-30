package users

import (
	"fmt"

	"github.com/gvu0110/bookstore_users-api/utils/errors"
)

// Data Access Object: the logic to access to database.
// Only entry point from the application to interact with the database

var (
	usersDB = make(map[int64]*User)
)

// Get function gets user from database
func (user *User) Get() *errors.RESTError {
	result := usersDB[user.ID]
	if result == nil {
		return errors.NewNotFoundRESTError(fmt.Sprintf("User ID %d not found", user.ID))
	}

	user.ID = result.ID
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreate = result.DateCreate
	return nil
}

// Save function saves user to database
func (user *User) Save() *errors.RESTError {
	current := usersDB[user.ID]
	if current != nil {
		return errors.NewBadRequestRESTError(fmt.Sprintf("User ID %d already exists", user.ID))
	}

	usersDB[user.ID] = user
	return nil
}
