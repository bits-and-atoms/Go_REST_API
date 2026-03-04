package routes

import (
	"net/http"

	"github.com/bits-and-atoms/Go_REST_API/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "healthy"})
	})

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id",updateEvent)
	authenticated.DELETE("/events/:id",deleteEvent)
	authenticated.POST("/events/:id/register",registerForEvent)
	authenticated.DELETE("/events/:id/register",cancelRegistration)
	server.POST("/signup",signup)
	server.POST("/login",login)
}
