package main

import (
	"github.com/shywn-mrk/event-booking-api/config"
	"github.com/shywn-mrk/event-booking-api/db"
	"github.com/shywn-mrk/event-booking-api/handlers"
	"github.com/shywn-mrk/event-booking-api/repositories"
	"github.com/shywn-mrk/event-booking-api/server"
	"github.com/shywn-mrk/event-booking-api/services"
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
