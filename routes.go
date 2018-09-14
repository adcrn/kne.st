package main

func initRoutes() {

    // Index route
    r.GET("/", showIndexPage)

    r.GET("/folders", showFoldersPage)

}
