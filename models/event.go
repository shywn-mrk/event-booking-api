package models

import (
	"time"
)

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"datetime"`
	UserID      int       `json:"user_id"`
}
