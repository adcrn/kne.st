package models

import (
    "time"
)

type folder struct {
    modelImpl
    Owner        string     // Should be user's randomly-generated hash
    FolderName   string
    Created      time.Date
    NumElements  int
}

// Creating fake folders until for testing
var folderList = []folder{
    folder{Owner: "nna24ga9zn2haa", FolderName: "terns", 
           Created: time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC), NumElements: 12},
    folder{Owner: "plabs42yabn6sd", FolderName: "2014 Milan", 
           Created: time.Date(2014, 09, 12, 08, 17, 42, 793654, time.UTC), NumElements: 27},
}

func getUsersFolders() []folder {
    return folderList
}