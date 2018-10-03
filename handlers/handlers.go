package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowIndexPage(c *gin.Context) {
	c.HTML(

		// HTTP 200 (OK)
		http.StatusOK,

		// Render home page
		"index.html",

		// Index template uses the following data
		gin.H{
			"title": "knest",
		},
	)
}

func StatusHandler(c *gin.Context) {

	c.JSON(

		http.StatusOK,

		gin.H{
			"status": "Good.",
		},
	)
}
