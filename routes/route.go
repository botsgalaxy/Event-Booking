package routes

import (
	"github.com/botsgalaxy/Event-Booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	authenticate := r.Group("/")
	authenticate.Use(middlewares.Authenticate)
	authenticate.POST("/events", middlewares.Authenticate, createEvent)
	authenticate.PUT("/events/:id", updateEvent)
	authenticate.DELETE("/events/:id", deleteEvent)
	authenticate.POST("/events/:id/register", registerForEvent)
	authenticate.DELETE("/events/:id/register", cancelRegistration)

	r.GET("/events", getEvents)
	r.GET("/events/:id", getEvent)
	r.POST("/signup", signUp)
	r.POST("/login", login)
}
