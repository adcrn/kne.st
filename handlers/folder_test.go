package handlers

import (
	"encoding/json"
	"kne.st/models"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// This function is used for setup before executing the test functions
func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}

func TestFetchUserFoldersValid(t *testing.T) {
	r := gin.Default()

	r.GET("/folders/:id", FetchUserFolders)

	req, _ := http.NewRequest("GET", "/folders/:3", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// If the request was not completed properly
	if w.Code != http.StatusOK {
		t.Fail()
	}

	response := []models.Folder{}
	json.Unmarshal([]byte(w.Body.String()), &response)

	value := response[0].OwnerID

	if value != 3 {
		t.Fail()
	}
}

func TestFetchUserFoldersInvalid(t *testing.T) {
	r := gin.Default()

	r.GET("/folders/:id", FetchUserFolders)

	req, _ := http.NewRequest("GET", "/folders/:blah", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 400 {
		t.Fail()
	}
}

func TestDeleteUserFolderValid(t *testing.T) {
	r := gin.Default()

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
	r := gin.Default()

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
