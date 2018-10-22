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
	//"github.com/gin-gonic/gin"
)

//var r *gin.Engine

func main() {
	db, err := postgres.Open(os.Getenv("DB"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	us := &postgres.UserService{DB: db}
	fs := &postgres.FolderService{DB: db}

	var h http.Handler
	h.UserService = us
	h.FolderService = fs

	// Gin's default router uses radix trees, helpful for our use case.
	//r = gin.Default()

	// Serve assets through static middleware.
	h.Engine.Use(static.Serve("/assets", static.LocalFile("./assets", true)))

	// Process the templates.
	h.Engine.LoadHTMLGlob("templates/*")

	h.Engine.Run()
}
