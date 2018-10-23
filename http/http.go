package http

import (
	"encoding/json"
	"github.com/adcrn/webknest"
	"github.com/adcrn/webknest/postgres"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Handler holds the interfaces to services
type Handler struct {
	*gin.Engine
	UserService   *postgres.UserService
	FolderService *postgres.FolderService
	Logger        gin.HandlerFunc
}

// NewHandler returns a handler that allows for interfacing with services
func NewHandler() *Handler {
	h := &Handler{
		Engine: gin.Default(),
		Logger: gin.Logger(),
	}

	// Allow for versioning of API by making a group
	h.Group("/api/v1")
	{
		h.POST("/register", h.register)

		// Separation of user- and folder-specific handler functions
		h.Group("/u")
		{
			h.GET("/:id/get", h.getUserInfo)
			h.POST("/:id/update", h.updateUserInfo)
			h.POST("/:id/changepass", h.changePassword)
			h.POST("/:id/changeemai", h.changeEmail)
		}
		h.Group("/f")
		{
			h.GET("/:id", h.fetchUserFolders)
			h.GET("/:id/:foldername", h.getFolderRecord)
			h.POST("/:id/:foldername/delete", h.deleteFolderRecord)
			h.POST("/:id/:foldername/update", h.updateFolderRecord)
			h.POST("/:id/upload", h.createFolderRecord)
		}
	}

	return h
}

// register is the function through which a user's desired credentials and
// details are taken and passed to the UserService interface.
func (h *Handler) register(c *gin.Context) {

	// Instantiate user object and convert JSON response into object
	var u webknest.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	// Create the user record in storage using the details from frontend
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

// getUserInfo returns all non-confidential information about a user
func (h *Handler) getUserInfo(c *gin.Context) {
	var u webknest.User
	userID, err := strconv.Atoi(c.Param("id")[1:])

	if err != nil {
		c.JSON(400, gin.H{"error": "Bad user ID"})
	}

	u, err = h.UserService.GetByID(userID)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	c.JSON(200, u)
}

// updateUserInfo allows for modification of non-sensitive information of a user
func (h *Handler) updateUserInfo(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	var u webknest.User
	var du webknest.DetailUpdate

	u, err := h.UserService.GetByID(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	if err = c.BindJSON(&du); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	err = h.UserService.UpdateDetails(u, du)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	c.JSON(204, gin.H{"response": "Update successful"})

}

// changePassword allows for the singular action of changing a user's password
// in the case that the current password is known; this is in contrast to a
// password reset in which the psasword is not known
func (h *Handler) changePassword(c *gin.Context) {
	var pu webknest.PasswordUpdate
	if err := c.BindJSON(&pu); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	userID, err := strconv.Atoi(c.Param("id")[1:])
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	err = h.UserService.ChangePassword(userID, pu)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	c.JSON(204, gin.H{"response": "success"})
}

// changeEmail allows for the singular action of changing email
func (h *Handler) changeEmail(c *gin.Context) {
	var email string
	userID, _ := strconv.Atoi(c.Param("id")[1:])

	// No need to make a struct for one field
	var response map[string]string
	body, _ := c.GetRawData()
	err := json.Unmarshal(body, &response)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	email = response["email"]
	err = h.UserService.ChangeEmail(userID, email)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	c.JSON(204, gin.H{"response": "success"})

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
			"status": "good",
		},
	)
}

// fetchUserFolders takes in the id parameter passed from the frontend
// and returns all folders of that particular user from the database
func (h *Handler) fetchUserFolders(c *gin.Context) {
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
	}

	// Otherwise, return all folders associated
	// with that user ID
	if folders, err := h.FolderService.ListByUser(userID); err == nil {
		c.JSON(200, folders)
	}

}

// getFolderRecord will retrieve the corresponding record from storage
func (h *Handler) getFolderRecord(c *gin.Context) {
	var f webknest.Folder

	userID, err := strconv.Atoi(c.Param("id")[1:])
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	folderName := c.Param("foldername")[1:]

	f, err = h.FolderService.GetByName(userID, folderName)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	c.JSON(200, f)
}

// createFolderRecord will add the record of a folder into storage by calling
// the upload function (defined elsewhere) and storing its returned folder path
func (h *Handler) createFolderRecord(c *gin.Context) {

}

// deleteFolderRecord will remove a folder record from storage; this should be
// used after a person downloads their processed folders
func (h *Handler) deleteFolderRecord(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id")[1:])
	foldername := c.Param("foldername")[1:]

	if err != nil {
		c.JSON(400, gin.H{"response": "malformed userID"})
	}

	err = h.FolderService.Delete(userID, foldername)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	c.JSON(204, gin.H{"response": "success"})
}

// updateFolderRecord will update details about the folder that weren't already
// present at time of upload, e.g. completed/downloaded.
func (h *Handler) updateFolderRecord(c *gin.Context) {
	var f webknest.Folder
	var fu webknest.FolderUpdate
	userID, err := strconv.Atoi(c.Param("id")[1:])
	foldername := c.Param("folder_name")

	f, err = h.FolderService.GetByName(userID, foldername)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	if err := c.BindJSON(&fu); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	if err = h.FolderService.Update(f, fu); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	c.JSON(204, gin.H{"response": "success"})
}
