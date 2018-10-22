package http

import (
	"github.com/adcrn/webknest"
	"github.com/adcrn/webknest/postgres"
	"github.com/gin-gonic/gin"
	//"golang.org/x/crypto/bcrypt"
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

	h.Group("/api/v1")
	{
		h.POST("/register", h.register)
		h.Group("/u")
		{
			h.GET("/:id/get", h.getUserInfo)
			h.POST("/:id/update", h.updateUserInfo)
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

	var u webknest.User
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

func (h *Handler) updateUserInfo(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	var u webknest.User
	var cu webknest.CredentialUpdate

	u, err := h.UserService.GetByID(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	if err = c.BindJSON(&cu); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	// Should probably break out password changes into their own function, will
	// do soon
	if cu.Password == "" {
		cu.Password = u.Password
	}

	err = h.UserService.Update(u, cu)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}

	c.JSON(204, gin.H{"response": "Update successful"})

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
		return
	}

	// Otherwise, return all folders associated
	// with that user ID
	if folders, err := h.FolderService.ListByUser(userID); err == nil {
		c.JSON(200, folders)
	}

}

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

func (h *Handler) createFolderRecord(c *gin.Context) {

}

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
