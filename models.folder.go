package main

import (
	"net/url"
	"time"
)

type folder struct {
	Owner         string    `json:"owner"`
	FolderName    string    `json:"foldername"`
	FolderNameURL string    `json:"foldernameurl"`
	Created       time.Time `json:"created"`
	NumElements   int       `json:"numelements"`
	Completed     bool      `json:"completed"`
	Downloaded    bool      `json:"downloaded"`
}

// Creating fake folders until for testing
var folder1 = &folder{Owner: "nna24ga9zn2haa", FolderName: "terns", FolderNameURL: url.QueryEscape("terns"),
	Created: time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC), NumElements: 12,
	Completed: false, Downloaded: false}
var folder2 = &folder{Owner: "plabs42yabn6sd", FolderName: "2014 Milan", FolderNameURL: url.QueryEscape("2014 Milan"),
	Created: time.Date(2014, 9, 12, 8, 17, 42, 793654, time.UTC), NumElements: 27,
	Completed: true, Downloaded: false}
var folders = []*folder{}

// This will take a user ID as a parameter and query the
// database to return a list of folders associated with the user.
func getUsersFolders() []*folder {
	folders = append(folders, folder1)
	folders = append(folders, folder2)
	return folders
}
