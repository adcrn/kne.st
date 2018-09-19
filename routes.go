package main

func initRoutes() {

	// Index route
	r.GET("/", showIndexPage)

	r.GET("/folders", showFoldersPage)

	userRoutes := r.Group("/u")
	{
		userRoutes.GET("/register", showRegistrationPage)

		userRoutes.POST("/register", register)
	}

}
