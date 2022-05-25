package repository

import (
	"fmt"
	"http_server/internal/domain"
	"reflect"
	"sync"
	"time"
)

type Calendar interface {
	CreateEvent(ev *domain.Event) error
	UpdateEvent(ev *domain.Event) error
	DeleteEvent(id string) error
	GetEventForDay(userID string, date *time.Time) []domain.Event
	GetEventForWeek(userID string, date *time.Time) []domain.Event
	GetEventForMonth(userID string, date *time.Time) []domain.Event
}

type calendar struct {
	Events map[string]domain.Event
	*sync.RWMutex
}

func NewCalendar() Calendar {
	return &calendar{
		Events: make(map[string]domain.Event),
		RWMutex: &sync.RWMutex{},
	}
}

func ValidEvent(ev *domain.Event) error {
	if ev.Date.IsZero() {
		return fmt.Errorf("empty date")
	}

	return nil
}

func (c *calendar) CreateEvent(ev *domain.Event) error {
	if err := ValidEvent(ev); err != nil {
		return fmt.Errorf("invalid event: %v", err)
	}

	c.RLock()
	if _, ok := c.Events[ev.ID]; ok {
		c.RUnlock()
		return fmt.Errorf("event with id=%s already exists", ev.ID)
	}
	c.RUnlock()
	c.Lock()
	c.Events[ev.ID] = *ev
	c.Unlock()
	return nil
}

func (c *calendar) UpdateEvent(ev *domain.Event) error {
	if err := ValidEvent(ev); err != nil {
		return fmt.Errorf("invalid event: %v", err)
	}

	c.RLock()
	if _, ok := c.Events[ev.ID]; ok {
		c.RUnlock()
		if reflect.DeepEqual(c.Events[ev.ID], ev) {
			return fmt.Errorf("no changes to update")
		}
		c.Lock()
		c.Events[ev.ID] = *ev
		c.Unlock()
		return nil
	}

	return fmt.Errorf("no event with id=%s exist, update declined", ev.ID)
}

func (c *calendar) DeleteEvent(id string) error {

	c.RLock()
	if _, ok := c.Events[id]; ok {
		c.RUnlock()

		delete(c.Events, id)
		return nil
	}
	c.RUnlock()

	return fmt.Errorf("no event with id=%s exist, delete declined", id)
}

func (c *calendar) GetEventForDay(userID string, date *time.Time) []domain.Event {
	res := make([]domain.Event, 0)
	for _, event := range c.Events {
		if event.UserID == userID {
			if date.Equal(event.Date) {
				res = append(res, event)
			}
		}
	}

	return res
}

func (c *calendar) GetEventForWeek(userID string, date *time.Time) []domain.Event {
	res := make([]domain.Event, 0)

	from := date.AddDate(0, 0, -int(date.Weekday()))
	to := date.AddDate(0, 0, 7-int(date.Weekday()))

	for _, event := range c.Events {
		if event.UserID == userID {
			if from.Before(event.Date) && to.After(event.Date) {
				res = append(res, event)
			}
		}
	}

	return res
}

func (c *calendar) GetEventForMonth(userID string, date *time.Time) []domain.Event {
	res := make([]domain.Event, 0)

	from := date.AddDate(0, 0, -int(date.Day()))
	to := from.AddDate(0, 1, 0)

	for _, event := range c.Events {
		if event.UserID == userID {
			if from.Before(event.Date) && to.After(event.Date) {
				res = append(res, event)
			}
		}
	}

	return res
}
