package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shywn-mrk/event-book-api/handlers"
)

func SetupRouter(eventHandler *handlers.EventHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/events", eventHandler.Create)
	r.GET("/events/:id", eventHandler.Read)
	r.PUT("/events/:id", eventHandler.Update)
	r.DELETE("/events/:id", eventHandler.Delete)
	r.GET("/events", eventHandler.List)

	return r
}
