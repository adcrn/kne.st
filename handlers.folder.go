package main

import (
	//"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func showFoldersPage(c *gin.Context) {
	// Replace with getUserFoldersByID
	folders := getUsersFolders()

	c.HTML(
		http.StatusOK,
		"folder.html",
		gin.H{
			"title":   "Folders - knest",
			"payload": folders,
		},
	)
}

func fetchUserFolders(c *gin.Context) {
	// Replace with getUserFoldersByID
	folders := getUsersFolders()

	//payload, _ := json.Marshal(folders)

	c.JSON(200, folders)

}
