package main

import "errors"

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

func registerNewUser(username, password string) (*user, error) {
	return nil, errors.New("User registration failed.")
}

func isUsernameAvailable(username string) bool {
	return false
}
