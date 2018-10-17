package http

import (
	"github.com/adcrn/knest_web"
	"github.com/gin-gonic/gin"
	"strconv"
)

func initRoutes() {

	// Index route
	r.GET("/", ShowIndexPage)

	// Status page
	r.GET("/status", StatusHandler)

	r.GET("/folders/:id", FetchUserFolders)

	userRoutes := r.Group("/u")
	{
		userRoutes.POST("/register", Register)
	}

}

// Register is the handler through which a user's desired credentials and
// details are taken and passed to RegisterNewUser.
func Register(c *gin.Context) {

	var u knest_web.User
	c.BindJSON(&u)

	if err := knest_web.RegisterNewUser(u); err == nil {

		c.JSON(

			200,

			gin.H{
				"response": "registration successful",
			},
		)
	} else {

		c.JSON(

			400,

			gin.H{
				"response": err.Error(),
			},
		)
	}
}

func showLoginPage(c *gin.Context) {
	//
}

func performLogin(c *gin.Context) {
	//
}

func performLogout(c *gin.Context) {
	//
}

// ShowIndexPage will return the home page of the website when the user
// navigates to the root of the site.
func ShowIndexPage(c *gin.Context) {
	c.HTML(

		// HTTP 200 (OK)
		200,

		// Render home page
		"index.html",

		// Index template uses the following data
		gin.H{
			"title": "knest_web",
		},
	)
}

// StatusHandler returns the status of each component of the site. Right now,
// it just returns the status of the site.
func StatusHandler(c *gin.Context) {

	c.JSON(

		200,

		gin.H{
			"status": "Good.",
		},
	)
}

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

	if errUser != nil {
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
