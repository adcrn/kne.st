package main

func TestUsernameAvailability(t *testing.T) {
    saveLists()

    // This user is definitely not in the list
    if !isUsernameAvailable("definitelynotinlist") {
        t.Fail()
    }

    // This user is definitely a part of the list
    if isUsernameAvailable("test1") {
        t.Fail()
    }

    // Add user to list of users
    registerNewUser("definitelynotinlist", "easypass")

    // Check to see that registration worked
    if isUsernameAvailable("definitelynotinlist") {
        t.Fail()
    }

    restoreLists()
}