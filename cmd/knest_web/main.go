// Main driver for knest webapp.
// (c) 2018, knest.

package main

import (
	"github.com/adcrn/knest_web"
	"github.com/adcrn/knest_web/http"
	"github.com/adcrn/knest_web/postgres"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
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
