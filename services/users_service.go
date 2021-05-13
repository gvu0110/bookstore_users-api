package services

import (
	"github.com/gvu0110/bookstore_users-api/domain/users"
	"github.com/gvu0110/bookstore_users-api/utils/date"
	"github.com/gvu0110/bookstore_users-api/utils/encryption"
	"github.com/gvu0110/bookstore_users-api/utils/errors"
)

// Core entire business logic, shouldn't be changed

var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.RESTError)
	GetUser(int64) (*users.User, *errors.RESTError)
	UpdateUser(users.User) (*users.User, *errors.RESTError)
	DeleteUser(int64) *errors.RESTError
	FindUsersByStatus(string) (users.Users, *errors.RESTError)
}

// CreateUser function creates a new user
func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RESTError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Status = users.StatusActive
	user.DateCreated = date.GetNowDBFormat()
	user.Password = encryption.GetMD5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUser function gets user's information
func (s *usersService) GetUser(userID int64) (*users.User, *errors.RESTError) {
	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

func (s *usersService) UpdateUser(user users.User) (*users.User, *errors.RESTError) {
	current, err := UsersService.GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	current.FirstName = user.FirstName
	current.LastName = user.LastName
	current.Email = user.Email
	current.Password = user.Password
	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

func (s *usersService) DeleteUser(userID int64) *errors.RESTError {
	result := &users.User{ID: userID}
	if err := result.Delete(); err != nil {
		return err
	}
	return nil
}

func (s *usersService) FindUsersByStatus(status string) (users.Users, *errors.RESTError) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}
