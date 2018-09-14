package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func showFoldersPage(c *gin.Context) {
    folders := getUsersFolders()

    c.HTML(
        http.StatusOK,
        "folder.html",
        gin.H{
            "title": "Folders - knest",
            "payload": folders,
        },
    )
}