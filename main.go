// Main driver for knest webapp.
// (c) 2018, knest.

package main

import (
    "github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
    // Gin's default router uses radix trees, helpful for our use case.
    router = gin.Default()

    // Process the templates
    router.LoadHTMLGlob("templates/*")

    // Initialize routes.
    initRoutes()

    router.Run()
}
