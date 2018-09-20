// Main driver for knest webapp.
// (c) 2018, knest.

package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

var r *gin.Engine

func main() {
	// Gin's default router uses radix trees, helpful for our use case.
	r = gin.Default()

	// Serve assets through static middleware.
	r.Use(static.Serve("/assets", static.LocalFile("./assets", true)))

	// Process the templates.
	r.LoadHTMLGlob("templates/*")

	// Initialize routes.
	initRoutes()

	r.Run()
}

func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}

}
