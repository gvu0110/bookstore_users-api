package users

import (
	"fmt"

	"github.com/gvu0110/bookstore_users-api/datasources/mysql/user_db"
	"github.com/gvu0110/bookstore_users-api/utils/date"
	"github.com/gvu0110/bookstore_users-api/utils/errors"
)

// Data Access Object: the logic to access to database.
// Only entry point from the application to interact with the database

const (
	queryInsertUser = "INSERT INTO users (first_name, last_name, email, date_created) VALUES (?, ?, ?, ?);"
	queryGetUser    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
)

// Get function gets user from database
func (user *User) Get() *errors.RESTError {
	stmt, err := user_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewBadRequestRESTError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to get userID %d: %s", user.ID, err.Error()))
	}
	return nil
}

// Save function saves user to database
func (user *User) Save() *errors.RESTError {
	stmt, err := user_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewBadRequestRESTError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to save user: %s", err.Error()))
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to save user: %s", err.Error()))
	}

	user.ID = userID
	return nil
}

func (user *User) Update() *errors.RESTError {
	stmt, err := user_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewBadRequestRESTError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}
