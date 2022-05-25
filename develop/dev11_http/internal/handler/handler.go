package handlers

import (
	"net/http"
)

type Handler interface {
	CreateEvent(w http.ResponseWriter, r *http.Request)
	UpdateEvent(w http.ResponseWriter, r *http.Request)
	DeleteEvent(w http.ResponseWriter, r *http.Request)
	GetEventForDay(w http.ResponseWriter, r *http.Request)
	GetEventForWeek(w http.ResponseWriter, r *http.Request)
	GetEventForMonth(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	Handler
}

func NewHandler() Handler {
	return &handler{
		Handler: NewCalendarHandler(),
	}
}
