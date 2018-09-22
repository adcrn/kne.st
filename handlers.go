package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
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

func statusHandler(c *gin.Context) {

	c.JSON(

		http.StatusOK,

		gin.H{
			"status": "Good.",
		},
	)
}
