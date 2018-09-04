// Main driver for knest webapp.
// (c) 2018, knest.

package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/static"
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
