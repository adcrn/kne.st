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

	username := c.PostForm("username")
	password := c.PostForm("password")
	fullname := c.PostForm("fullname")
	email := c.PostForm("email")

	if _, err := registerNewUser(username, password, fullname, email); err == nil {

		render(c, gin.H{
			"title": "Successful registration, logging in..."}, "login_success.html")
	} else {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"ErrorTitle":   "Registration Failed",
			"ErrorMessage": err.Error()})
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
