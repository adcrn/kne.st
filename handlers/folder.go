package handlers

import (
	"github.com/gin-gonic/gin"
	"kne.st/models"
	"strconv"
)

// FetchUserFolders takes in the id parameter passed from the frontend
// and returns all folders of that particular user from the database
func FetchUserFolders(c *gin.Context) {
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
	folders := models.GetUsersFolders(userID)

	c.JSON(200, folders)
}

func deleteUserFolder(c *gin.Context) {
	userID, errUser := strconv.Atoi(c.Param("id")[1:])
	foldername := c.Param("foldername")[1:]

	if err_user != nil {
		c.JSON(

			400,

			gin.H{
				"response": "malformed userID",
			},
		)
		return
	}

	_, err := models.DeleteFolderDatabaseRecord(userID, foldername)

	if err != nil {
		c.JSON(

			400,

			gin.H{
				"response": err.Error(),
			},
		)
		return
	}

	c.JSON(
		204,
		gin.H{
			"response": "success",
		},
	)
}
