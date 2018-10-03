package main

import (
	"kne.st/handlers"
)

func initRoutes() {

	// Index route
	r.GET("/", handlers.ShowIndexPage)

	// Status page
	r.GET("/status", handlers.StatusHandler)

	//r.GET("/folders", showFoldersPage)
	r.GET("/folders/:id", handlers.FetchUserFolders)

	userRoutes := r.Group("/u")
	{
		//	userRoutes.GET("/register", showRegistrationPage)

		userRoutes.POST("/register", handlers.Register)
	}

}
