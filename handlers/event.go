package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shywn-mrk/event-book-api/services"
)


type EventHandler struct {
	service services.EventService
}

func NewEventHandler(service services.EventService) *EventHandler {
	return &EventHandler{service}
}

func (h *EventHandler) Create(*gin.Context) {

}

func (h *EventHandler) Read(*gin.Context) {
	
}

func (h *EventHandler) Update(*gin.Context) {

}

func (h *EventHandler) Delete(*gin.Context) {
	
}

func (h *EventHandler) List(*gin.Context) {
	
}
