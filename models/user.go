package models

import (
	"errors"
	"strings"
)

type User struct {
	Username string `json: username`
	Password string `json: password`
	FullName string `json: fullname`
	Email    string `json: email`
}

// This should be populated from a database, but for prototyping, we'll define them here
var UserList = []User{
	{Username: "test1", Password: "pass1", FullName: "Test OneGuy", Email: "guy1@test.com"},
	{Username: "test2", Password: "pass2", FullName: "Test TwoGal", Email: "gal2@test.com"},
	{Username: "test3", Password: "pass3", FullName: "Test ThreeThey", Email: "they3@test.com"},
}

func RegisterNewUser(u User) (bool, error) {

	// Make sure password isn't empty.
	if strings.TrimSpace(u.Password) == "" {
		return false, errors.New("Passwords cannot be empty.")
	}

	if !isUsernameAvailable(u.Username) {
		return false, errors.New("Username is already taken.")
	}

	if !isEmailAvailable(u.Email) {
		return false, errors.New("There is already an account associated with this email address.")
	}

	UserList = append(UserList, u)

	// this function will eventually return a boolean representing
	// the outcome of actually inserting a user into the database

	return true, nil
}

// This will be replaced by a call to the users table of the database
func isUsernameAvailable(username string) bool {

	for _, u := range UserList {
		if u.Username == username {
			return false
		}
	}

	return true
}

// This will be replaced by a call to the users table of the database
func isEmailAvailable(email string) bool {

	for _, u := range UserList {
		if u.Email == email {
			return false
		}
	}

	return true
}

func isUserValid(username, password string) bool {
	return false
}
