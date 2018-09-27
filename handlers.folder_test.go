package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchUserFoldersValid(t *testing.T) {
	r := getRouter(true)

	r.GET("/folders/:id", fetchUserFolders)

	req, _ := http.NewRequest("GET", "/folders/:3", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// If the request was not completed properly
	if w.Code != http.StatusOK {
		t.Fail()
	}

	response := []folder{}
	json.Unmarshal([]byte(w.Body.String()), &response)

	value := response[0].OwnerID

	if value != 3 {
		t.Fail()
	}
}

func TestFetchUserFoldersInvalid(t *testing.T) {
	r := getRouter(true)

	r.GET("/folders/:id", fetchUserFolders)

	req, _ := http.NewRequest("GET", "/folders/:blah", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestDeleteUserFolderValid(t *testing.T) {
	r := getRouter(true)

	r.POST("/folders/:id/:foldername/delete", deleteUserFolder)

	req, _ := http.NewRequest("POST", "/folders/:3/terns/delete", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 204 {
		t.Fail()
	}

	// there needs to be a database call to the users table
	// so that proper access to storage is tested
}

func TestDeleteUserFolderInvalid(t *testing.T) {
	r := getRouter(true)

	r.POST("/folders/:id/:foldername/delete", deleteUserFolder)

	// Not sure how we'll get to this point through the app,
	// but it's always good to test stuff, right?
	req, _ := http.NewRequest("POST", "/folders/:3/blah/delete", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Fail()
	}

	// there should be a database call that looks for the
	// folder name under that user and folder and
	// should fail if neither is found
}
