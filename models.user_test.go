package main

import "testing"

func TestUsernameAvailability(t *testing.T) {
	saveLists()

	// This user is definitely not in the list
	if !isUsernameAvailable("newone") {
		t.Fail()
	}

	// This user is definitely a part of the list
	if isUsernameAvailable("test1") {
		t.Fail()
	}

	// Add user to list of users
	registerNewUser("newone", "easypass", "New Person", "new@one.com")

	// Check to see that registration worked
	if isUsernameAvailable("newone") {
		t.Fail()
	}

	restoreLists()
}

func TestEmailAvailability(t *testing.T) {
	saveLists()

	// Email not in list.
	if !isEmailAvailable("not@here.com") {
		t.Fail()
	}

	// Email in list.
	if isEmailAvailable("guy1@test.com") {
		t.Fail()
	}

	registerNewUser("newone", "easypass", "New Person", "new@one.com")

	// Email now in list.
	if isEmailAvailable("new@one.com") {
		t.Fail()
	}

	restoreLists()
}

func TestValidUserRegistration(t *testing.T) {
	saveLists()

	// Attempt to register new user
	u, err := registerNewUser("newone", "easypass", "New Person", "new@one.com")

	// Check for possible bad outcomes
	if err != nil || u.Username == "" {
		t.Fail()
	}

	restoreLists()

}

func TestInvalidUserRegistration(t *testing.T) {
	saveLists()

	// Attempt to register user already in database
	u, err := registerNewUser("test1", "pass1", "Test OneGuy", "guy1@test.com")
	// Fail if no error or if a user is returned
	if err == nil || u != nil {
		t.Fail()
	}

	// Attempt to register with no password
	u, err = registerNewUser("invalidtest", "", "Space Woman", "space@pass.com")

	if err == nil || u != nil {
		t.Fail()
	}

	restoreLists()
}

func TestUserValidity(t *testing.T) {
	saveLists()

	// Valid login.
	if !isUserValid("test1", "pass1") {
		t.Fail()
	}

	// Invalid login.
	if isUserValid("test2", "pass1") {
		t.Fail()
	}

	// No spaces for password.
	if isUserValid("test1", "") {
		t.Fail()
	}

	// No spaces for username.
	if isUserValid("", "pass1") {
		t.Fail()
	}

	// No captials in username.
	if isUserValid("test1", "pass1") {
		t.Fail()
	}
}
