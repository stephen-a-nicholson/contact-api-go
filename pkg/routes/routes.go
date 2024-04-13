package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stephen-a-nicholson/contact-api-go/pkg/handlers"
)

func InitRoutes(router *gin.Engine) {
    router.POST("/contacts", handlers.CreateContact)
    router.GET("/contacts/:id", handlers.GetContact)
    router.PUT("/contacts/:id", handlers.UpdateContact)
    router.DELETE("/contacts/:id", handlers.DeleteContact)
}