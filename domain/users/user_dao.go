package users

import (
	"database/sql"
	"fmt"

	"github.com/gvu0110/bookstore_users-api/datasources/mysql/user_db"
	"github.com/gvu0110/bookstore_utils-go/logger"
	"github.com/gvu0110/bookstore_utils-go/rest_errors"
)

// Data Access Object: the logic to access to database.
// Only entry point from the application to interact with the database

const (
	queryInsertUser                = "INSERT INTO users (first_name, last_name, email, date_created, status, password) VALUES (?, ?, ?, ?, ?, ?);"
	queryGetUser                   = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser                = "UPDATE users SET first_name=?, last_name=?, email=?, password=? WHERE id=?;"
	queryDeleteUser                = "DELETE FROM users WHERE id=?;"
	queryFindUsersByStatus         = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
	queryGetUserByEmailAndPassword = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=? AND status=?;"
)

// Get function gets user from database
func (user *User) Get() rest_errors.RESTError {
	stmt, err := user_db.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("Error when trying to prepare get user statement", err)
		return rest_errors.NewInternalServerRESTError("Database error", err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	err = result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status)
	switch {
	case err == sql.ErrNoRows:
		return rest_errors.NewNotFoundRESTError(fmt.Sprintf("UserID %d not found", user.ID))
	case err != nil:
		logger.Error(fmt.Sprintf("Error when trying to get userID %d: %s", user.ID, err.Error()), err)
		return rest_errors.NewInternalServerRESTError("Database error", err)
	default:
		return nil
	}
}

// Save function saves user to database
func (user *User) Save() rest_errors.RESTError {
	stmt, err := user_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("Error when trying to prepare save user statement", err)
		return rest_errors.NewInternalServerRESTError("Database error", err)
	}
	defer stmt.Close()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		logger.Error(fmt.Sprintf("Error when trying to save user: %s", err.Error()), err)
		return rest_errors.NewInternalServerRESTError("Database error", err)
	}

	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error(fmt.Sprintf("Error when trying to save user: %s", err.Error()), err)
		return rest_errors.NewInternalServerRESTError("Database error", err)
	}

	user.ID = userID
	return nil
}

func (user *User) Update() rest_errors.RESTError {
	stmt, err := user_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("Error when trying to prepare update user statement", err)
		return rest_errors.NewInternalServerRESTError("Database error", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password, user.ID)
	if err != nil {
		logger.Error(fmt.Sprintf("Error when trying to update user: %s", err.Error()), err)
		return rest_errors.NewInternalServerRESTError("Database error", err)
	}
	return nil
}

func (user *User) Delete() rest_errors.RESTError {
	stmt, err := user_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("Error when trying to prepare delete user statement", err)
		return rest_errors.NewInternalServerRESTError("Database error", err)
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.ID); err != nil {
		logger.Error(fmt.Sprintf("Error when trying to delete userID %d: %s", user.ID, err.Error()), err)
		return rest_errors.NewInternalServerRESTError("Database error", err)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, rest_errors.RESTError) {
	stmt, err := user_db.Client.Prepare(queryFindUsersByStatus)
	if err != nil {
		logger.Error("Error when trying to prepare find users by status statement", err)
		return nil, rest_errors.NewInternalServerRESTError("Database error", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error(fmt.Sprintf("Error when trying to find users by status: %s", err.Error()), err)
		return nil, rest_errors.NewInternalServerRESTError("Database error", err)
	}
	defer rows.Close()

	result := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error(fmt.Sprintf("Error when trying to scan user row into user struct: %s", err.Error()), err)
			return nil, rest_errors.NewInternalServerRESTError("Database error", err)
		}
		result = append(result, user)
	}

	if len(result) == 0 {
		return nil, rest_errors.NewNotFoundRESTError(fmt.Sprintf("No user matching status %s", status))
	}
	return result, nil
}

func (user *User) GetByEmailAndPassword() rest_errors.RESTError {
	stmt, err := user_db.Client.Prepare(queryGetUserByEmailAndPassword)
	if err != nil {
		logger.Error("Error when trying to prepare get user by email and password statement", err)
		return rest_errors.NewInternalServerRESTError("Database error", err)
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Email, user.Password, StatusActive)
	err = result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status)
	switch {
	case err == sql.ErrNoRows:
		return rest_errors.NewNotFoundRESTError("Invalid credentials")
	case err != nil:
		logger.Error(fmt.Sprintf("Error when trying to get userID %d: %s", user.ID, err.Error()), err)
		return rest_errors.NewInternalServerRESTError("Database error", err)
	default:
		return nil
	}
}
