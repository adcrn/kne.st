package models

import (
	"errors"
	"strings"
)

// User contains login credentials and details about their profile including
// subscription type, which will dictate certain capabilities
type User struct {
	ID               int    `json:"id"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	FullName         string `json:"fullname"`
	Email            string `json:"email"`
	SubscriptionType int    `json:"sub_type"`
}

// UserList should be populated from a database, but for prototyping, we'll define them here
var UserList = []User{
	{ID: 1, Username: "test1", Password: "pass1", FullName: "Test OneGuy", Email: "guy1@test.com", SubscriptionType: 1},
	{ID: 2, Username: "test2", Password: "pass2", FullName: "Test TwoGal", Email: "gal2@test.com", SubscriptionType: 2},
	{ID: 3, Username: "test3", Password: "pass3", FullName: "Test ThreeThey", Email: "they3@test.com", SubscriptionType: 3},
}

// RegisterNewUser attempts to insert a new user into the database
func RegisterNewUser(u User) error {

	// Make sure password isn't empty.
	if strings.TrimSpace(u.Password) == "" {
		return errors.New("Passwords cannot be empty.")
	}

	if !isUsernameAvailable(u.Username) {
		return errors.New("Username is already taken.")
	}

	if !isEmailAvailable(u.Email) {
		return errors.New("There is already an account associated with this email address.")
	}

	UserList = append(UserList, u)

	// this function will eventually return a boolean representing
	// the outcome of actually inserting a user into the database

	return nil
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

/*
func isUserValid(username, password string) bool {
	return false
}*/
