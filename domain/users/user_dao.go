package users

import (
	"fmt"

	"github.com/gvu0110/bookstore_users-api/datasources/mysql/user_db"
	"github.com/gvu0110/bookstore_users-api/utils/errors"
)

// Data Access Object: the logic to access to database.
// Only entry point from the application to interact with the database

const (
	queryInsertUser        = "INSERT INTO users (first_name, last_name, email, date_created, status, password) VALUES (?, ?, ?, ?, ?, ?);"
	queryGetUser           = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser        = "UPDATE users SET first_name=?, last_name=?, email=?, password=? WHERE id=?;"
	queryDeleteUser        = "DELETE FROM users WHERE id=?;"
	queryFindUsersByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

// Get function gets user from database
func (user *User) Get() *errors.RESTError {
	stmt, err := user_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewBadRequestRESTError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
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

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
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

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password, user.ID)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (user *User) Delete() *errors.RESTError {
	stmt, err := user_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewBadRequestRESTError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("Error when trying to delete userID %d: %s", user.ID, err.Error()))
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RESTError) {
	stmt, err := user_db.Client.Prepare(queryFindUsersByStatus)
	if err != nil {
		return nil, errors.NewBadRequestRESTError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewBadRequestRESTError(err.Error())
	}
	defer rows.Close()

	result := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, errors.NewInternalServerError(err.Error())
		}
		result = append(result, user)
	}

	if len(result) == 0 {
		return nil, errors.NewNotFoundRESTError(fmt.Sprintf("No user matching status %s", status))
	}
	return result, nil
}
