package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func showRegistrationPage(c *gin.Context) {
    c.HTML(

        http.StatusOK,

        "register.html",

        gin.H{
            "title": "Register - knest",
        },
    )
}

func register(c *gin.Context) {
    //
}