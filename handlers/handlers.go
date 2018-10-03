package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ShowIndexPage will return the home page of the website when the user
// navigates to the root of the site.
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

// StatusHandler returns the status of each component of the site. Right now,
// it just returns the status of the site.
func StatusHandler(c *gin.Context) {

	c.JSON(

		http.StatusOK,

		gin.H{
			"status": "Good.",
		},
	)
}
