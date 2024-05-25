package server

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/shywn-mrk/event-book-api/handlers"
	"go.uber.org/fx"
)

func NewEngine(lc fx.Lifecycle, eventHandler *handlers.EventHandler) *gin.Engine {
	engine := SetupEngine(eventHandler)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go engine.Run()
			return nil
		},
		OnStop: func(context.Context) error {
			return nil
		},
	})

	return engine
}
