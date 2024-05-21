package repositories

import (
	"github.com/shywn-mrk/event-book-api/models"
	"gorm.io/gorm"
)

type EventRepository interface {
  Create(event models.Event) error
  Read(id int) (*models.Event, error)
  Update(id int) error
  Delete(id int) error
  List() ([]models.Event, error)
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) Create(event models.Event) error {
	// Save the event to the database
	return nil
}

func (r *eventRepository) Read(id int) (*models.Event, error) {
	// Get an event by ID from the database
	return nil, nil
}

func (r *eventRepository) Update(id int) error {
	// Update the event in the database
	return nil
}

func (r *eventRepository) Delete(id int) error {
	// Delete the event from the database
	return nil
}

func (r *eventRepository) List() ([]models.Event, error) {
	// Get all events from the database
	return nil, nil
}
