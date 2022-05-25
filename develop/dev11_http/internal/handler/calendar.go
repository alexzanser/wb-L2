package handlers

import (
	"encoding/json"
	"fmt"
	"http_server/internal/domain"
	"http_server/internal/repository"
	"mime"
	"net/http"
)

//CalendarHandler is a concrete structure that implements Handler methods
type CalendarHandler struct {
	repo *repository.Repository
}

//NewCalendarHandler return new instance of calendarHandler
func NewCalendarHandler() *CalendarHandler {
	return &CalendarHandler{
		repo: repository.NewRepository(),
	}
}

func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

//CreateEvent creates new calendar event
func (c *CalendarHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var ev domain.Event

	if err := dec.Decode(&ev); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	if err := c.repo.Calendar.CreateEvent(&ev); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	type OkResponse struct {
		Result	string `json:"result"`
	}

	// type ErrResponse struct {
	// 	err	string `json:"err"`
	// }

	renderJSON(w, OkResponse{Result: fmt.Sprintf("new event with id %s created", ev.ID)})
}

//UpdateEvent updates existing event
func (c *CalendarHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var ev domain.Event

	if err := dec.Decode(&ev); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	if err := c.repo.Calendar.UpdateEvent(&ev); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	type OkResponse struct {
		Result	string `json:"result"`
	}

	// type ErrResponse struct {
	// 	err	string `json:"err"`
	// }

	renderJSON(w, OkResponse{Result: fmt.Sprintf("event with id %s updated", ev.ID)})
}

//DeleteEvent deletes event
func (c *CalendarHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var ev domain.Event

	if err := dec.Decode(&ev); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	if err := c.repo.Calendar.DeleteEvent(ev.ID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	type OkResponse struct {
		Result	string `json:"result"`
	}

	// type ErrResponse struct {
	// 	err	string `json:"err"`
	// }

	renderJSON(w, OkResponse{Result: fmt.Sprintf("event with id %s deleted", ev.ID)})
}

//GetEventForDay return list of events on a given day
func (c *CalendarHandler) GetEventForDay(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var ev domain.Event

	if err := dec.Decode(&ev); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	events := c.repo.Calendar.GetEventForDay(ev.UserID, &ev.Date)

	type OkResponse struct {
		Result	[]domain.Event `json:"result"`
	}

	// type ErrResponse struct {
	// 	err	string `json:"err"`
	// }

	renderJSON(w, OkResponse{Result: events})
}

//GetEventForWeek return list of events on a given week
func (c *CalendarHandler) GetEventForWeek(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var ev domain.Event

	if err := dec.Decode(&ev); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	events := c.repo.Calendar.GetEventForDay(ev.UserID, &ev.Date)

	type OkResponse struct {
		Result	[]domain.Event `json:"result"`
	}

	// type ErrResponse struct {
	// 	err	string `json:"err"`
	// }

	renderJSON(w, OkResponse{Result: events})
}

//GetEventForMonth return list of events on a given month
func (c *CalendarHandler) GetEventForMonth(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		http.Error(w, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var ev domain.Event

	if err := dec.Decode(&ev); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}


	events := c.repo.Calendar.GetEventForDay(ev.UserID, &ev.Date)

	type OkResponse struct {
		Result	[]domain.Event `json:"result"`
	}

	// type ErrResponse struct {
	// 	err	string `json:"err"`
	// }

	renderJSON(w, OkResponse{Result: events})
}
