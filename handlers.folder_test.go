package main

import (
	"encoding/json"
	//"fmt"
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
