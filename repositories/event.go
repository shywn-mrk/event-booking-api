package repositories

import (
	"fmt"

	"github.com/shywn-mrk/event-booking-api/models"
	"gorm.io/gorm"
)

type EventRepository interface {
	Create(event *models.Event) error
	Read(id int64) (*models.Event, error)
	Update(id int64, event *models.Event) error
	Delete(id int64) error
	List() ([]models.Event, error)
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) Create(event *models.Event) error {
	return r.db.Create(event).Error
}

func (r *eventRepository) Read(id int64) (*models.Event, error) {
	var event models.Event
	err := r.db.First(&event, id).Error
	if err != nil {
		return nil, fmt.Errorf("could not read event: %w", err)
	}

	return &event, nil
}

func (r *eventRepository) Update(id int64, event *models.Event) error {
	obj, _ := r.Read(id)

	event.ID = obj.ID
	err := r.db.Model(obj).Updates(event).Error
	if err != nil {
		return fmt.Errorf("could not update event: %w", err)
	}

	return nil
}

func (r *eventRepository) Delete(id int64) error {
	return r.db.Delete(&models.Event{}, id).Error
}

func (r *eventRepository) List() ([]models.Event, error) {
	var events []models.Event
	err := r.db.Find(&events).Error
	return events, err
}
