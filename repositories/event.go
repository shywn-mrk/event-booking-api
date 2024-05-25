package repositories

import (
	"fmt"

	"github.com/shywn-mrk/event-book-api/models"
	"gorm.io/gorm"
)

type EventRepository interface {
  Create(event models.Event) error
  Read(id int64) (*models.Event, error)
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
	err := r.db.Create(&event).Error
	if err != nil {
		return fmt.Errorf("could not create event: %w", err)
	}

	return nil
}

func (r *eventRepository) Read(id int64) (*models.Event, error) {
	var event models.Event
	err := r.db.First(&event, id).Error
	if err != nil {
		return nil, fmt.Errorf("could not read event: %w", err)
	}

	return &event, nil
}

func (r *eventRepository) Update(id int) error {
	// Update the event in the database
	return nil
}

func (r *eventRepository) Delete(id int) error {
	err := r.db.Delete(&models.Event{}, id).Error
	if err != nil {
		return fmt.Errorf("could not delete event: %w", err)
	}

	return nil
}

func (r *eventRepository) List() ([]models.Event, error) {
	var events []models.Event
	err := r.db.Find(&events).Error
	return events, err
}
