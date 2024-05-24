package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/shywn-mrk/event-book-api/config"
	"github.com/shywn-mrk/event-book-api/handlers"
	"github.com/shywn-mrk/event-book-api/models"
	"github.com/shywn-mrk/event-book-api/repositories"
	"github.com/shywn-mrk/event-book-api/routes"
	"github.com/shywn-mrk/event-book-api/services"
	"go.uber.org/fx"
	"golang.org/x/net/context"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.Event{})
	return db, nil
}

func NewRouter(lc fx.Lifecycle, eventHandler *handlers.EventHandler) *gin.Engine {
	router := routes.SetupEventsRouter(eventHandler)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go router.Run()
			return nil
		},
		OnStop: func(context.Context) error {
			return nil
		},
	})

	return router
}

func main() {
	fx.New(
		fx.Provide(
			config.LoadConfig,
			NewDatabase,
			repositories.NewEventRepository,
			services.NewEventService,
			handlers.NewEventHandler,
			NewRouter,
		),
		fx.Invoke(func (router *gin.Engine) {}),
	).Run()
}
