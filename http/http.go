package http

import (
	"encoding/json"
	"github.com/adcrn/webknest"
	"github.com/adcrn/webknest/errors"
	"github.com/adcrn/webknest/postgres"
	"github.com/auth0-community/go-auth0"
	"github.com/gin-gonic/gin"
	jose "gopkg.in/square/go-jose.v2"
	"os"
	"strconv"
)

// Handler holds the interfaces to services
type Handler struct {
	*gin.Engine
	UserService   *postgres.UserService
	FolderService *postgres.FolderService
}

// NewHandler returns a handler that allows for interfacing with services
func NewHandler() *Handler {

	h := &Handler{
		Engine: gin.Default(),
	}

	// Allow for versioning of API by making a group
	v1 := h.Group("/api/v1")
	{
		v1.POST("/register", h.register)

		// Separation of user- and folder-specific handler functions
		userRoutes := v1.Group("/u")
		userRoutes.Use(authRequired())
		{
			userRoutes.GET("/get/:id", h.getUserInfo)
			userRoutes.POST("/update/:id", h.updateUserInfo)
			userRoutes.POST("/changepass/:id", h.changePassword)
			userRoutes.POST("/changeemail/:id", h.changeEmail)
		}
		folderRoutes := v1.Group("/f")
		folderRoutes.Use(authRequired())
		{
			folderRoutes.GET("/fetch/:id", h.fetchUserFolders)
			folderRoutes.GET("/get/:id/:foldername", h.getFolderRecord)
			folderRoutes.POST("/delete/:id/:foldername", h.deleteFolderRecord)
			folderRoutes.POST("/update/:id/:foldername", h.updateFolderRecord)
			folderRoutes.POST("/upload/:id", h.createFolderRecord)
		}
	}

	return h
}

// UseMiddleware is just to access the internal gin Use function
func (h *Handler) UseMiddleware(middleware gin.HandlerFunc) {
	h.Use(middleware)
}

// RunServer is just to access the internal gin Run function
func (h *Handler) RunServer() {
	h.Run()
}

// register is the function through which a user's desired credentials and
// details are taken and passed to the UserService interface.
func (h *Handler) register(c *gin.Context) {

	// Instantiate user object and convert JSON response into object
	var u webknest.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(400, gin.H{"error": errors.ErrBadRequest})
		return
	}

	// Create the user record in storage using the details from frontend
	if _, err := h.UserService.Create(u); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, gin.H{"response": "success"})
}

// getUserInfo returns all non-confidential information about a user
func (h *Handler) getUserInfo(c *gin.Context) {
	var u webknest.User
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": errors.ErrBadRequest})
		return
	}

	u, err = h.UserService.GetByID(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, u)
}

// updateUserInfo allows for modification of non-sensitive information of a user
func (h *Handler) updateUserInfo(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": errors.ErrBadRequest})
		return
	}

	var u webknest.User
	var du webknest.DetailUpdate

	u, err = h.UserService.GetByID(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": errors.ErrDatabaseError})
		return
	}

	if err = c.BindJSON(&du); err != nil {
		c.JSON(400, gin.H{"error": errors.ErrBadRequest})
		return
	}

	err = h.UserService.UpdateDetails(u, du)
	if err != nil {
		c.JSON(400, gin.H{"error": errors.ErrDatabaseError})
		return
	}

	c.JSON(204, gin.H{"response": "Update successful"})

}

// changePassword allows for the singular action of changing a user's password
// in the case that the current password is known; this is in contrast to a
// password reset in which the psasword is not known
func (h *Handler) changePassword(c *gin.Context) {
	var pu webknest.PasswordUpdate
	if err := c.BindJSON(&pu); err != nil {
		c.JSON(400, gin.H{"error": errors.ErrBadRequest})
		return
	}

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": errors.ErrBadRequest})
		return
	}

	err = h.UserService.ChangePassword(userID, pu)
	if err != nil {
		c.JSON(400, gin.H{"error": errors.ErrDatabaseError})
		return
	}

	c.JSON(204, gin.H{"response": "success"})
}

// changeEmail allows for the singular action of changing email
func (h *Handler) changeEmail(c *gin.Context) {
	var email string
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": errors.ErrBadRequest})
		return
	}

	// No need to make a struct for one field
	var response map[string]string
	body, _ := c.GetRawData()
	err = json.Unmarshal(body, &response)
	if err != nil {
		c.JSON(400, gin.H{"error": errors.ErrBadRequest})
		return
	}

	email = response["email"]
	err = h.UserService.ChangeEmail(userID, email)
	if err != nil {
		c.JSON(400, gin.H{"error": errors.ErrDatabaseError})
		return
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
	var folders []webknest.Folder
	// Retrieve user ID from GET request and
	// convert it to an integer
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": errors.ErrBadRequest})
		return
	}

	// Otherwise, return all folders associated
	// with that user ID
	if folders, err = h.FolderService.ListByUser(userID); err != nil {
		c.JSON(400, gin.H{"error": errors.ErrDatabaseError})
	}

	c.JSON(200, folders)
}

// getFolderRecord will retrieve the corresponding record from storage
func (h *Handler) getFolderRecord(c *gin.Context) {
	var f webknest.Folder

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": errors.ErrBadRequest})
		return
	}

	folderName := c.Param("foldername")

	f, err = h.FolderService.GetByName(userID, folderName)
	if err != nil {
		c.JSON(400, gin.H{"error": errors.ErrDatabaseError})
		return
	}

	c.JSON(200, f)
}

// createFolderRecord will add the record of a folder into storage by calling
// the upload function (defined elsewhere) and storing its returned folder path
func (h *Handler) createFolderRecord(c *gin.Context) {
	// Retrieving the user object for now until sessions are implemented
	var u webknest.User
	var f webknest.Folder

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": errors.ErrBadRequest})
		return
	}

	u, err = h.UserService.GetByID(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": errors.ErrDatabaseError})
		return
	}

	// Pass the username to the upload function
	f, err = UploadHandler(c, u)
	if err != nil {
		c.JSON(400, gin.H{"error": errors.ErrUploadFailed})
		return
	}

	_, err = h.FolderService.Create(f)
	if err != nil {
		c.JSON(400, gin.H{"error": errors.ErrDatabaseError})
		return
	}

	c.JSON(200, gin.H{"response": "success"})
}

// deleteFolderRecord will remove a folder record from storage; this should be
// used after a person downloads their processed folders
func (h *Handler) deleteFolderRecord(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	foldername := c.Param("foldername")

	if err != nil {
		c.JSON(400, gin.H{"error": errors.ErrBadRequest})
		return
	}

	err = h.FolderService.Delete(userID, foldername)

	if err != nil {
		c.JSON(400, gin.H{"error": errors.ErrDatabaseError})
		return
	}

	c.JSON(204, gin.H{"response": "success"})
}

// updateFolderRecord will update details about the folder that weren't already
// present at time of upload, e.g. completed/downloaded.
func (h *Handler) updateFolderRecord(c *gin.Context) {
	var f webknest.Folder
	var fu webknest.FolderUpdate

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": errors.ErrBadRequest})
		return
	}

	foldername := c.Param("folder_name")

	f, err = h.FolderService.GetByName(userID, foldername)
	if err != nil {
		c.JSON(400, gin.H{"error": errors.ErrDatabaseError})
		return
	}

	if err := c.BindJSON(&fu); err != nil {
		c.JSON(400, gin.H{"error": errors.ErrBadRequest})
		return
	}

	if err = h.FolderService.Update(f, fu); err != nil {
		c.JSON(400, gin.H{"error": errors.ErrDatabaseError})
		return
	}

	c.JSON(204, gin.H{"response": "success"})
}

// This function was taken from Auth0's website in order to simplify user auth
func authRequired() gin.HandlerFunc {
	return func(c *gin.Context) {

		var auth0Domain = "https://" + os.Getenv("AUTH0_DOMAIN") + "/"
		client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: auth0Domain + ".well-known/jwks.json"}, nil)
		configuration := auth0.NewConfiguration(client, []string{os.Getenv("AUTH0_API_IDENTIFIER")}, auth0Domain, jose.RS256)
		validator := auth0.NewValidator(configuration, nil)

		_, err := validator.ValidateRequest(c.Request)

		if err != nil {
			c.JSON(401, gin.H{"error": "token is not valid"})
			c.Abort()
			return
		}
		c.Next()
	}
}
