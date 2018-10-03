package handlers

import (
	"github.com/gin-gonic/gin"
	"kne.st/models"
)

// Register is the handler through which a user's desired credentials and
// details are taken and passed to RegisterNewUser.
func Register(c *gin.Context) {

	var u models.User
	c.BindJSON(&u)

	if _, err := models.RegisterNewUser(u); err == nil {

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
