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
	registerNewUser("newone", "easypass")

	// Check to see that registration worked
	if isUsernameAvailable("definitelynotinlist") {
		t.Fail()
	}

	restoreLists()
}

func TestValidUserRegistration(t *testing.T) {
	saveLists()

	// Attempt to register new user
	u, err := registerNewUser("newone", "easypass")

	// Check for possible bad outcomes
	if err != nil || u.Username == "" {
		t.Fail()
	}

	restoreLists()

}

func TestInvalidUserRegistration(t *testing.T) {
	saveLists()

	// Attempt to register user already in database
	u, err := registerNewUser("test1", "pass1")

	// Fail if no error or if a user is returned
	if err == nil || u != nil {
		t.Fail()
	}

	// Attempt to register with no password
	u, err = registerNewUser("invalidtest", "")

	if err == nil || u != nil {
		t.Fail()
	}

	restoreLists()
}
