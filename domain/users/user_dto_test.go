package users

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateInvalidEmail(t *testing.T) {
	fakeUser := User{
		Email:    "adam.vugmail.com",
		Password: "password",
	}
	err := fakeUser.Validate()
	assert.EqualValues(t, http.StatusBadRequest, err.StatusCode())
	assert.EqualValues(t, "Invalid email address", err.Message())
	assert.EqualValues(t, "bad_request", err.Error())
}

func TestValidateInvalidPassword(t *testing.T) {
	fakeUser := User{
		Email:    "adam.vu@gmail.com",
		Password: "",
	}

	fakeUser.Validate()
	err := fakeUser.Validate()
	assert.EqualValues(t, http.StatusBadRequest, err.StatusCode())
	assert.EqualValues(t, "Invalid password", err.Message())
	assert.EqualValues(t, "bad_request", err.Error())
}

func TestValidateNoError(t *testing.T) {
	fakeUser := User{
		Email:    "adam.vu@gmail.com",
		Password: "password",
	}

	fakeUser.Validate()
	err := fakeUser.Validate()
	assert.Nil(t, err)
}
