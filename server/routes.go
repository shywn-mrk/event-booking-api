package server

import (
	"github.com/gin-gonic/gin"
	"github.com/shywn-mrk/event-booking-api/handlers"
)

func SetupEngine(eventHandler *handlers.EventHandler) *gin.Engine {
	engine := gin.Default()

	engine.POST("/events", eventHandler.Create)
	engine.GET("/events/:id", eventHandler.Read)
	engine.PUT("/events/:id", eventHandler.Update)
	engine.DELETE("/events/:id", eventHandler.Delete)
	engine.GET("/events", eventHandler.List)

	return engine
}
