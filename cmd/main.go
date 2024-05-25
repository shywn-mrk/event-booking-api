package main

import (
	"github.com/shywn-mrk/event-book-api/config"
	"github.com/shywn-mrk/event-book-api/db"
	"github.com/shywn-mrk/event-book-api/handlers"
	"github.com/shywn-mrk/event-book-api/repositories"
	"github.com/shywn-mrk/event-book-api/server"
	"github.com/shywn-mrk/event-book-api/services"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.LoadConfig,
			db.NewDatabase,
			repositories.NewEventRepository,
			services.NewEventService,
			handlers.NewEventHandler,
		),
		fx.Invoke(server.NewEngine),
	).Run()
}
