// Main driver for knest webapp.
// (c) 2018, knest.

package main

import (
	"database/sql"
	"fmt"
	"github.com/adcrn/webknest/backend/http"
	"github.com/adcrn/webknest/backend/postgres"
	"log"

	"github.com/gin-contrib/static"
)

const (
	dbUser     = "postgres"
	dbPassword = "test"
	dbName     = "knest_test"
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

	h := http.NewHandler()
	h.UserService = us
	h.FolderService = fs

	// Serve assets through static middleware.
	h.UseMiddleware(static.Serve("/assets", static.LocalFile("./assets", true)))

	// Process the templates.
	//h.LoadHTMLGlob("/templates/*")

	h.RunServer()
}
