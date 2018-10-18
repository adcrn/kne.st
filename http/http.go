package http

import (
	"github.com/adcrn/webknest"
	"github.com/gin-gonic/gin"
	"strconv"
)

// UserHandler implements the UserService interface along with a router and
// logger in order to work with requests from the frontend for user data
type UserHandler struct {
	*gin.Engine

	UserService webknest.UserService

	Logger gin.HandlerFunc
}

// FolderHandler implements the FolderService interface along with a router and
// logger in order to work with requests from the frontend for folder data
type FolderHandler struct {
	*gin.Engine

	FolderService webknest.FolderService

	Logger gin.HandlerFunc
}

// NewUserHandler instantiates a UserHandler along with some routes
func NewUserHandler() *UserHandler {
	h := &UserHandler{
		Engine: gin.Default(),
		Logger: gin.Logger(),
	}
	h.POST("/api/v1/register", h.register)

	h.Group("/api/v1/u")
	{
		h.GET("/:id/get", h.getUserInfo)
		h.POST("/:id/update", h.updateUserInfo)
	}

	return h
}

// NewFolderHandler instantiates a FolderHandler along with some routes
func NewFolderHandler() *FolderHandler {
	h := &FolderHandler{
		Engine: gin.Default(),
		Logger: gin.Logger(),
	}

	h.Group("/api/v1/f")
	{
		h.GET("/:id", h.fetchUserFolders)
		h.GET("/:id/:foldername", h.getFolderRecord)
		h.POST("/:id/:foldername/delete", h.deleteFolderRecord)
		h.POST("/:id/upload", h.createFolderRecord)
	}

	return h
}

// register is the function through which a user's desired credentials and
// details are taken and passed to the UserService interface.
func (h *UserHandler) register(c *gin.Context) {

	var u *webknest.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	if _, err := h.UserService.Create(u); err == nil {

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
				"error": err.Error(),
			},
		)
	}
}

func (h *UserHandler) getUserInfo(c *gin.Context) {

}

func (h *UserHandler) updateUserInfo(c *gin.Context) {

}

// showIndexPage will return the home page of the website when the user
// navigates to the root of the site.
func showIndexPage(c *gin.Context) {
	c.HTML(

		// HTTP 200 (OK)
		200,

		// Render home page
		"index.html",

		// Index template uses the following data
		gin.H{
			"title": "webknest",
		},
	)
}

// statusHandler returns the status of each component of the site. Right now,
// it just returns the status of the site.
func statusHandler(c *gin.Context) {

	c.JSON(

		200,

		gin.H{
			"status": "Good.",
		},
	)
}

// fetchUserFolders takes in the id parameter passed from the frontend
// and returns all folders of that particular user from the database
func (h *FolderHandler) fetchUserFolders(c *gin.Context) {
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
	if folders, err := h.FolderService.ListByUser(userID); err == nil {
		c.JSON(200, folders)
	}

}

func (h *FolderHandler) getFolderRecord(c *gin.Context) {

}

func (h *FolderHandler) createFolderRecord(c *gin.Context) {

}

func (h *FolderHandler) deleteFolderRecord(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id")[1:])
	foldername := c.Param("foldername")[1:]

	if err != nil {
		c.JSON(

			400,

			gin.H{
				"response": "malformed userID",
			},
		)
		return
	}

	err = h.FolderService.Delete(userID, foldername)

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
