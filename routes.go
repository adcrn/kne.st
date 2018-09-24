package main

func initRoutes() {

	// Index route
	r.GET("/", showIndexPage)

	// Status page
	r.GET("/status", statusHandler)

	//r.GET("/folders", showFoldersPage)
	r.GET("/folders/:id", fetchUserFolders)

	userRoutes := r.Group("/u")
	{
		userRoutes.GET("/register", showRegistrationPage)

		userRoutes.POST("/register", register)
	}

}
