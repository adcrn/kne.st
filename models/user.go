package models

import (
	"errors"
	"strings"
)

// User contains login credentials and details about their profile including
// subscription type, which will dictate certain capabilities
type User struct {
	Username         string `json:"username"`
	Password         string `json:"password"`
	FullName         string `json:"fullname"`
	Email            string `json:"email"`
	SubscriptionType int    `json:"sub_type"`
}

// UserStorage is the interface through which methods will access the database
// in order to operate on user objects.
type UserStorage interface {
	List(...UserFilter) ([]User, error)
	Get(...UserFilter) (User, error)
	Create(User) error
	Update(User) error
	Delete(...UserFilter) error
}

// UserFilter is the set of critera that will be used to select certain users
type UserFilter func(*UserFilterConfig) error

// UserFilterConfig is the struct that will be edited and then called by the
// UserFilter interface for searching.
type UserFilterConfig struct {
	Username         string
	FullName         string
	Email            string
	SubscriptionType int
}

// UserUsernameFilter sets the username field
func UserUsernameFilter(username string) UserFilter {
	return func(fc *UserFilterConfig) error {
		fc.Username = username
		return nil
	}
}

// UserFullNameFilter sets the full name field
func UserFullNameFilter(fullName string) UserFilter {
	return func(fc *UserFilterConfig) error {
		fc.FullName = fullName
		return nil
	}
}

// UserEmailFilter sets the email field
func UserEmailFilter(email string) UserFilter {
	return func(fc *UserFilterConfig) error {
		fc.Email = email
		return nil
	}
}

// UserSubscriptionTypeFilter sets the subscription type field
func UserSubscriptionTypeFilter(subType int) UserFilter {
	return func(fc *UserFilterConfig) error {
		fc.SubscriptionType = subType
		return nil
	}
}

// UserList should be populated from a database, but for prototyping, we'll define them here
var UserList = []User{
	{Username: "test1", Password: "pass1", FullName: "Test OneGuy", Email: "guy1@test.com", SubscriptionType: 1},
	{Username: "test2", Password: "pass2", FullName: "Test TwoGal", Email: "gal2@test.com", SubscriptionType: 2},
	{Username: "test3", Password: "pass3", FullName: "Test ThreeThey", Email: "they3@test.com", SubscriptionType: 3},
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
