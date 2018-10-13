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

	u := User{4, "newone", "easypass", "New Person", "new@one.com", 1}

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

	u := User{4, "newone", "easypass", "New Person", "new@one.com", 1}

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

	u := User{4, "newone", "easypass", "New Person", "new@one.com", 1}

	// Attempt to register new user
	err := RegisterNewUser(u)

	// Check for possible bad outcomes
	if err != nil {
		t.Fail()
	}

	restoreLists()
}

func TestInvalidUserRegistration(t *testing.T) {
	saveLists()

	uAlreadyTaken := User{1, "test1", "pass1", "Test OneGuy", "guy1@test.com", 1}

	// Attempt to register user already in database
	errAlreadyTaken := RegisterNewUser(uAlreadyTaken)

	// Fail if no error or if a user is returned
	if errAlreadyTaken == nil {
		t.Fail()
	}

	uNoPass := User{5, "invalidtest", "", "Space Woman", "space@pass.com", 3}

	// Attempt to register with no password
	errNoPass := RegisterNewUser(uNoPass)

	if errNoPass == nil {
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
