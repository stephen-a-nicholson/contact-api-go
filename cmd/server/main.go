package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Define routes
    router.POST("/contacts", createContact)
    router.GET("/contacts/:id", getContact)
    router.PUT("/contacts/:id", updateContact)
    router.DELETE("/contacts/:id", deleteContact)

    // Start the server
    router.Run(":8080")
}
