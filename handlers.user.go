package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//func generateSessionToken() string {
// This will need to be in a safe, secure way in
// order to persist user sessions on the app.

//	return nil
//}

func showRegistrationPage(c *gin.Context) {

	render(c, gin.H{
		"title": "Register"}, "register.html")
}

func register(c *gin.Context) {

	var u user
	c.BindJSON(&u)

	if _, err := registerNewUser(u); err == nil {

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
