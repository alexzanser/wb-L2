package handlers

import (
	"net/http"
)

//Handler shows methods needed to hanle incoming requests
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

//NewHandler return new instance of type handler which implements Handler interface
func NewHandler() Handler {
	return &handler{
		Handler: NewCalendarHandler(),
	}
}
