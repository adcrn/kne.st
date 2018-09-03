// Main driver for knest webapp.
// (c) 2018, knest.

package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
    // Gin's default router uses radix trees, helpful for our use case.
    router = gin.Default()

    // Process the templates
    router.LoadHTMLGlob("templates/*")

    // Route for index.
    router.GET("/", func(c *gin.Context) {
        c.HTML(
            // Return 200 (OK)
            http.StatusOK,

            // Use home page template
            "index.html",

            // Set the title used inside template
            gin.H{
                "title": "knest",
            },
        )
    })

    router.Run()
}
