package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shywn-mrk/event-booking-api/models"
	"github.com/shywn-mrk/event-booking-api/services"
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

	if err := h.service.Create(&event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func (h *EventHandler) Read(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	event, err := h.service.Read(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func (h *EventHandler) Update(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var newEvent models.Event
	if err := ctx.ShouldBindJSON(&newEvent); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event, err := h.service.Read(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.Update(id, &newEvent); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func (h *EventHandler) Delete(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 0)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.service.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (h *EventHandler) List(ctx *gin.Context) {
	events, err := h.service.List()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, events)
}
