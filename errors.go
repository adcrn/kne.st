package webknest

// General errors
const (
	ErrUnauthorized  = Error("Unauthorized")
	ErrBadRequest    = Error("Bad request")
	ErrDatabaseError = Error("Database error")
)

// User errors
const (
	ErrEmailExists      = Error("This email address is already in use")
	ErrEmailNotFound    = Error("An account associated with this email address could not be found")
	ErrPassDoesNotMatch = Error("Incorrect password, please try again")
	ErrUsernameExists   = Error("This username is already in use")
	ErrUsernameNotFound = Error("An account associated with this username could not be found")
)

// Folder errors
const (
	ErrFolderExists   = Error("A folder with this name already exists")
	ErrFolderNotFound = Error("A folder with this name could not be found")
)

// Error allows for custom error strings
type Error string

func (e Error) Error() string { return string(e) }
