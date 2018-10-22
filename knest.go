package webknest

import (
	"time"
)

// User contains login credentials and details about their profile including
// subscription type, which will dictate certain capabilities
type User struct {
	ID               int    `json:"id"`
	Username         string `json:"username"`
	Password         string `json:"-"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	SubscriptionType int    `json:"sub_type"`
}

// CredentialUpdate allows for easy updating of user details
type CredentialUpdate struct {
	Password         string `json:"password"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Email            string `json:"email"`
	SubscriptionType int    `json:"sub_type"`
}

// UserService is the interface through which the handlers will interact with
// the user data model
type UserService interface {
	ListBySubscriptionType(int) ([]*User, error)
	GetByID(int) (*User, error)
	GetByUsername(string) (*User, error)
	Create(*User) (int, error)
	Update(*User, *CredentialUpdate) error
	Delete(int) error
}

// Folder is the main construct that will associated with users.
type Folder struct {
	OwnerID     int       `json:"owner"`
	FolderName  string    `json:"foldername"`
	S3Path      string    `json:"s3_path"`
	UploadTime  time.Time `json:"upload_time"`
	NumElements int       `json:"num_elements"`
	Completed   bool      `json:"completed"`
	Downloaded  bool      `json:"downloaded"`
}

// FolderUpdate allows for easy modification of the two important flags.
type FolderUpdate struct {
	Completed  bool `json:"completed"`
	Downloaded bool `json:"downloaded"`
}

// FolderService is the interface through which the handlers will interact with
// the folder data model
type FolderService interface {
	ListByUser(int) (*Folder, error)
	GetByName(int, string) (*Folder, error)
	Create(*Folder) (int, error)
	Update(*Folder, *FolderUpdate) error
	Delete(int, string) error
}