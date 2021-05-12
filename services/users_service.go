package services

import (
	"github.com/gvu0110/bookstore_users-api/domain/users"
	"github.com/gvu0110/bookstore_users-api/utils/errors"
)

// Core entire business logic, shouldn't be changed

// CreateUser function creates a new user
func CreateUser(user users.User) (*users.User, *errors.RESTError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUser function gets user's information
func GetUser(userID int64) (*users.User, *errors.RESTError) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateUser(user users.User) (*users.User, *errors.RESTError) {
	current, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	current.FirstName = user.FirstName
	current.LastName = user.LastName
	current.Email = user.Email
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}
