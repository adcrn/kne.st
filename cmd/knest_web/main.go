// Main driver for knest webapp.
// (c) 2018, knest.

package main

import (
	"github.com/adcrn/webknest"
	"github.com/adcrn/webknest/http"
	"github.com/adcrn/webknest/postgres"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

//var r *gin.Engine

func main() {
	db, err := postgres.Open(os.getEnv("DB"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	us := &postgres.UserService{DB: db}
	fs := &postgres.FolderService{DB: db}

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
