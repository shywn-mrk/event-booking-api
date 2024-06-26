package services

import (
	"github.com/shywn-mrk/event-booking-api/models"
	"github.com/shywn-mrk/event-booking-api/repositories"
)

type EventService interface {
	Create(event *models.Event) error
	Read(id int64) (*models.Event, error)
	Update(id int64, event *models.Event) error
	Delete(id int64) error
	List() ([]models.Event, error)
}

type eventService struct {
	repo repositories.EventRepository
}

func NewEventService(repo repositories.EventRepository) EventService {
	return &eventService{repo}
}

func (s *eventService) Create(event *models.Event) error {
	return s.repo.Create(event)
}

func (s *eventService) Read(id int64) (*models.Event, error) {
	return s.repo.Read(id)
}

func (s *eventService) Update(id int64, event *models.Event) error {
	return s.repo.Update(id, event)
}

func (s *eventService) Delete(id int64) error {
	return s.repo.Delete(id)
}

func (s *eventService) List() ([]models.Event, error) {
	return s.repo.List()
}
