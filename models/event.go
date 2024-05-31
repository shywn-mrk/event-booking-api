package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name        string    `gorm:"not null" json:"title" binding:"required"`
	Description string    `gorm:"not null" json:"description" binding:"required"`
	Location    string    `gorm:"not null" json:"location" binding:"required"`
	DateTime    time.Time `gorm:"not null" json:"datetime" binding:"required"`
	UserID      int       `gorm:"not null" json:"user_id" binding:"required"`
}
