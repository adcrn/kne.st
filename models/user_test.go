package models

import "testing"

var tmpFolderList []*Folder
var tmpUserList []User

// This function is used to store the main lists into the temporary one
// for testing
func saveLists() {
	tmpFolderList = Folders
	tmpUserList = UserList
}

func restoreLists() {
	Folders = tmpFolderList
	UserList = tmpUserList
}

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

	u := User{"newone", "easypass", "New Person", "new@one.com"}

	// Add user to list of users
	RegisterNewUser(u)

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

	u := User{"newone", "easypass", "New Person", "new@one.com"}

	// Add user to list of users
	RegisterNewUser(u)

	// Email now in list.
	if isEmailAvailable("new@one.com") {
		t.Fail()
	}

	restoreLists()
}

func TestValidUserRegistration(t *testing.T) {
	saveLists()

	u := User{"newone", "easypass", "New Person", "new@one.com"}

	// Attempt to register new user
	result, err := RegisterNewUser(u)

	// Check for possible bad outcomes
	if err != nil || result != true {
		t.Fail()
	}

	restoreLists()
}

func TestInvalidUserRegistration(t *testing.T) {
	saveLists()

	uAlreadyTaken := User{"test1", "pass1", "Test OneGuy", "guy1@test.com"}

	// Attempt to register user already in database
	resultAlreadyTaken, errAlreadyTaken := RegisterNewUser(uAlreadyTaken)

	// Fail if no error or if a user is returned
	if errAlreadyTaken == nil || resultAlreadyTaken != false {
		t.Fail()
	}

	uNoPass := User{"invalidtest", "", "Space Woman", "space@pass.com"}

	// Attempt to register with no password
	resultNoPass, errNoPass := RegisterNewUser(uNoPass)

	if errNoPass == nil || resultNoPass != false {
		t.Fail()
	}

	restoreLists()
}

/*func TestUserValidity(t *testing.T) {
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

	// No capitals in username.
	if isUserValid("test1", "pass1") {
		t.Fail()
	}

	restoreLists()
}*/
