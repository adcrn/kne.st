package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func showRegistrationPage(c *gin.Context) {
	c.HTML(

		http.StatusOK,

		"register.html",

		gin.H{
			"title": "Register - knest",
		},
	)
}

func register(c *gin.Context) {
	//
}
