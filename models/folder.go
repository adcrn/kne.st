package models

import (
	"net/url"
	"time"
)

// Folder is the main construct that will associated with users.
type Folder struct {
	OwnerID       int       `json:"owner"`
	FolderName    string    `json:"foldername"`
	FolderNameURL string    `json:"foldernameurl"`
	Created       time.Time `json:"created"`
	NumElements   int       `json:"numelements"`
	Completed     bool      `json:"completed"`
	Downloaded    bool      `json:"downloaded"`
}

// FolderStorage is the interface through which methods will access the database
// in order to operate on folder objects.
type FolderStorage interface {
	List(...FolderFilter) ([]Folder, error)
	Get(...FolderFilter) (Folder, error)
	Create(Folder) error
	Update(Folder) error
	Delete(...FolderFilter) error
}

// FolderFilter is the set of critera that will be used to select certain
// folders
type FolderFilter func(*FolderFilterConfig) error

// FolderFilterConfig is the struct that will be edited and then called by the
// FolderFilter interface for searching.
type FolderFilterConfig struct {
	OwnerID    int
	FolderName string
	Created    time.Time
	Completed  bool
	Downloaded bool
}

// FolderOwnerIDFilter sets the ownerID field
func FolderOwnerIDFilter(ownerID int) FolderFilter {
	return func(fc *FolderFilterConfig) error {
		fc.OwnerID = ownerID
		return nil
	}
}

// FolderFolderNameFilter sets the folderName field
func FolderFolderNameFilter(folderName string) FolderFilter {
	return func(fc *FolderFilterConfig) error {
		fc.FolderName = folderName
		return nil
	}
}

// FolderCreatedFilter sets the created field
func FolderCreatedFilter(created time.Time) FolderFilter {
	return func(fc *FolderFilterConfig) error {
		fc.Created = created
		return nil
	}
}

// FolderCompletedFilter sets the completed field
func FolderCompletedFilter(completed bool) FolderFilter {
	return func(fc *FolderFilterConfig) error {
		fc.Completed = completed
		return nil
	}
}

// FolderDownloadedFilter sets the downloaded field
func FolderDownloadedFilter(downloaded bool) FolderFilter {
	return func(fc *FolderFilterConfig) error {
		fc.Downloaded = downloaded
		return nil
	}
}

// Creating fake folders until for testing
var folder1 = &Folder{OwnerID: 3, FolderName: "terns", FolderNameURL: url.QueryEscape("terns"),
	Created: time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC), NumElements: 12,
	Completed: false, Downloaded: false}
var folder2 = &Folder{OwnerID: 3, FolderName: "2014 Milan", FolderNameURL: url.QueryEscape("2014 Milan"),
	Created: time.Date(2014, 9, 12, 8, 17, 42, 793654, time.UTC), NumElements: 27,
	Completed: true, Downloaded: false}

// Folders is a list of fake folders
var Folders = []*Folder{}

// GetUsersFolders will take a user ID as a parameter and query the
// database to return a list of folders associated with the user.
func GetUsersFolders(id int) []*Folder {
	Folders = append(Folders, folder1)
	Folders = append(Folders, folder2)
	return Folders
}

// DeleteFolderDatabaseRecord takes in the ID of the user that owns the folder
// and will attempt to delete the record of the folder from the database.
func DeleteFolderDatabaseRecord(ownerid int, foldername string) (bool, error) {

	// Return false with an error if the user does not
	// exist in the database or if the folder name does not
	// match a folder that is associated with that particular user ID.

	return true, nil
}
