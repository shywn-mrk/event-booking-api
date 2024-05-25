package db

import (
	"fmt"

	"github.com/shywn-mrk/event-book-api/config"
	"github.com/shywn-mrk/event-book-api/models"
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
