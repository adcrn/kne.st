// Main driver for knest webapp.
// (c) 2018, knest.

package main

import (
	//"github.com/adcrn/webknest"
	"github.com/adcrn/webknest/http"
	"github.com/adcrn/webknest/postgres"
	"log"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func main() {
	db, err := postgres.Open(os.Getenv("DB"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	us := &postgres.UserService{DB: db}
	fs := &postgres.FolderService{DB: db}

	var h http.Handler
	var uh http.UserHandler
	var fh http.FolderHandler
	uh.UserService = us
	fh.FolderService = fs
	h.UserHandler = uh
	h.FolderHandler = fh

	// Gin's default router uses radix trees, helpful for our use case.
	r = gin.Default()

	// Serve assets through static middleware.
	r.Use(static.Serve("/assets", static.LocalFile("./assets", true)))

	// Process the templates.
	r.LoadHTMLGlob("templates/*")

	r.Run()
}
