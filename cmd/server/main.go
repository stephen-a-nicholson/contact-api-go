package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stephen-a-nicholson/contact-api-go/pkg/routes"
)

func main() {
    router := gin.Default()

    // Initialize routes
    routes.InitRoutes(router)

    // Start the server
    router.Run(":8080")
}