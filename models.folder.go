package main

import (
    "time"
    "net/url"
)

type folder struct {
    Owner           string     // Should be user's randomly-generated hash
    FolderName      string
    FolderNameURL   string
    Created         time.Time
    NumElements     int
    Completed       bool
    Downloaded      bool
}

// Creating fake folders until for testing
var folderList = []folder{
    folder{Owner: "nna24ga9zn2haa", FolderName: "terns", FolderNameURL: url.QueryEscape("terns"), 
           Created: time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC), NumElements: 12,
           Completed: false, Downloaded: false},
    folder{Owner: "plabs42yabn6sd", FolderName: "2014 Milan", FolderNameURL: url.QueryEscape("2014 Milan"), 
           Created: time.Date(2014, 9, 12, 8, 17, 42, 793654, time.UTC), NumElements: 27,
           Completed: true, Downloaded: false},
}

func getUsersFolders() []folder {
    return folderList
}