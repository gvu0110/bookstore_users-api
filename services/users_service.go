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
	LoginRequest(users.LoginRequest) (*users.User, *errors.RESTError)
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
	user := &users.User{ID: userID}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *usersService) UpdateUser(user users.User) (*users.User, *errors.RESTError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	currentUser, err := UsersService.GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	currentUser.FirstName = user.FirstName
	currentUser.LastName = user.LastName
	currentUser.Email = user.Email
	currentUser.Password = user.Password
	if err := currentUser.Update(); err != nil {
		return nil, err
	}
	return currentUser, nil
}

func (s *usersService) DeleteUser(userID int64) *errors.RESTError {
	user := &users.User{ID: userID}
	if err := user.Delete(); err != nil {
		return err
	}
	return nil
}

func (s *usersService) FindUsersByStatus(status string) (users.Users, *errors.RESTError) {
	user := &users.User{}
	return user.FindByStatus(status)
}

func (s *usersService) LoginRequest(request users.LoginRequest) (*users.User, *errors.RESTError) {
	user := &users.User{
		Email:    request.Email,
		Password: request.Password,
	}
	if err := user.GetByEmailAndPassword(); err != nil {
		return nil, err
	}
	return user, nil
}
