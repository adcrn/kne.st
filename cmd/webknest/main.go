// Main driver for knest webapp.
// (c) 2018, knest.

package main

import (
	"database/sql"
	"fmt"
	"github.com/adcrn/webknest/http"
	"github.com/adcrn/webknest/postgres"
	"log"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

const (
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "test"
)

func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	us := &postgres.UserService{DB: db}
	fs := &postgres.FolderService{DB: db}
	r := gin.Default()

	var h http.Handler
	h.Engine = r
	h.UserService = us
	h.FolderService = fs

	// Serve assets through static middleware.
	h.Use(static.Serve("/assets", static.LocalFile("./assets", true)))

	// Process the templates.
	//h.LoadHTMLGlob("/templates/*")

	h.Run()
}
