package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

/*func showFoldersPage(c *gin.Context) {
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
}*/

func fetchUserFolders(c *gin.Context) {
	// Retrieve user ID from GET request and
	// convert it to an integer
	userID, err := strconv.Atoi(c.Param("id")[1:])

	// If ID is not a parsable number,
	// return an error string
	if err != nil {
		c.JSON(

			400,

			gin.H{
				"response": "malformed userID",
			},
		)
		return
	}

	// Otherwise, return all folders associated
	// with that user ID
	folders := getUsersFolders(userID)

	c.JSON(200, folders)
}
