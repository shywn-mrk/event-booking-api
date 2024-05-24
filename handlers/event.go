package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shywn-mrk/event-book-api/models"
	"github.com/shywn-mrk/event-book-api/services"
)


type EventHandler struct {
	service services.EventService
}

func NewEventHandler(service services.EventService) *EventHandler {
	return &EventHandler{service}
}

func (h *EventHandler) Create(ctx *gin.Context) {
	var event models.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.Create(event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, event)
}

func (h *EventHandler) Read(ctx *gin.Context) {

}

func (h *EventHandler) Update(ctx *gin.Context) {

}

func (h *EventHandler) Delete(ctx *gin.Context) {
	
}

func (h *EventHandler) List(ctx *gin.Context) {
	events, err := h.service.List()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if events == nil {
		ctx.JSON(http.StatusNoContent, gin.H{})
		return
	}

	ctx.JSON(http.StatusOK, events)
}
