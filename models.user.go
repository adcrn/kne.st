package main

import (
	"errors"
	"fmt"
	"strings"
)

type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
	FullName string `json:"fullname"`
	Email    string `json:"email"`
}

// This should be populated from a database, but for prototyping, we'll define them here
var userList = []user{
	{Username: "test1", Password: "pass1", FullName: "Test OneGuy", Email: "guy1@test.com"},
	{Username: "test2", Password: "pass2", FullName: "Test TwoGal", Email: "gal2@test.com"},
	{Username: "test3", Password: "pass3", FullName: "Test ThreeThey", Email: "they3@test.com"},
}

func registerNewUser(username, password, fullname, email string) (*user, error) {

	// Make sure password isn't empty.
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("Passwords cannot be empty.")
	}

	if !isUsernameAvailable(username) {
		return nil, errors.New("Username is already taken.")
	}

	if !isEmailAvailable(email) {
		return nil, errors.New("There is already an account associated with this email address.")
	}

	u := user{Username: username, Password: password, FullName: fullname, Email: email}

	userList = append(userList, u)

	return &u, nil
}

// This will be replaced by a call to the users table of the database
func isUsernameAvailable(username string) bool {

	for _, u := range userList {
		if u.Username == username {
			return false
		}
	}

	return true
}

// This will be replaced by a call to the users table of the database
func isEmailAvailable(email string) bool {

	for _, u := range userList {
		if u.Email == email {
			return false
		}
	}

	return true
}
