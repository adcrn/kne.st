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

	var h http.Handler
	h.UserService = us
	h.FolderService = fs

	// Serve assets through static middleware.
	h.Engine.Use(static.Serve("/assets", static.LocalFile("./assets", true)))

	// Process the templates.
	h.Engine.LoadHTMLGlob("templates/*")

	h.Engine.Run()
}
