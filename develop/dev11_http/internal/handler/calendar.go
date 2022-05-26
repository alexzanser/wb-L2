package handlers

import (
	"encoding/json"
	"fmt"
	"http_server/internal/domain"
	"http_server/internal/repository"
	"mime"
	"net/http"
	"time"
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

type OkResponse struct {
	Result string `json:"result"`
}

type EventResponse struct {
	Result []domain.Event `json:"result"`
}

func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func checkMediaType(w http.ResponseWriter, r *http.Request) bool {
	contentType := r.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)

	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%v"}`, err.Error()), http.StatusBadRequest)
		return false
	}

	if mediatype != "application/json" {
		http.Error(w, `{"error": "expect application/json Content-Type"}`, http.StatusUnsupportedMediaType)
		return false
	}

	return true
}

//CreateEvent creates new calendar event
func (c *CalendarHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {

	if checkMediaType(w, r) == false {
		return
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var ev domain.Event

	if err := dec.Decode(&ev); err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%v"}`, err.Error()), http.StatusBadRequest)
		return
	}

	if err := c.repo.Calendar.CreateEvent(&ev); err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%v"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	renderJSON(w, OkResponse{Result: fmt.Sprintf("new event with id %s created", ev.ID)})
}

//UpdateEvent updates existing event
func (c *CalendarHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
	if checkMediaType(w, r) == false {
		return
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var ev domain.Event

	if err := dec.Decode(&ev); err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%v"}`, err.Error()), http.StatusBadRequest)
		return
	}

	if err := c.repo.Calendar.UpdateEvent(&ev); err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%v"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	renderJSON(w, OkResponse{Result: fmt.Sprintf("event with id %s updated", ev.ID)})
}

//DeleteEvent deletes event
func (c *CalendarHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	if checkMediaType(w, r) == false {
		return
	}

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var ev domain.Event

	if err := dec.Decode(&ev); err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%v"}`, err.Error()), http.StatusBadRequest)
		return
	}

	if err := c.repo.Calendar.DeleteEvent(ev.ID); err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%v"}`, err.Error()), http.StatusInternalServerError)
		return
	}

	renderJSON(w, OkResponse{Result: fmt.Sprintf("event with id %s deleted", ev.ID)})
}

//GetEventForDay return list of events on a given day
func (c *CalendarHandler) GetEventForDay(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("user_id") || !r.URL.Query().Has("date") {
		http.Error(w, `{"error": "not enough parameter"}`, http.StatusBadRequest)
		return
	}

	UserID := r.URL.Query().Get("user_id")
	date, err  := time.Parse(time.RFC3339, r.URL.Query().Get("date"))

	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%v"}`, err.Error()), http.StatusBadRequest)
		return
	}

	events := c.repo.Calendar.GetEventForDay(UserID, &date)

	renderJSON(w, EventResponse{Result: events})
}

//GetEventForWeek return list of events on a given week
func (c *CalendarHandler) GetEventForWeek(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("user_id") || !r.URL.Query().Has("date") {
		http.Error(w, `{"error": "not enough parameter"}`, http.StatusBadRequest)
		return
	}

	UserID := r.URL.Query().Get("user_id")
	date, err  := time.Parse(time.RFC3339, r.URL.Query().Get("date"))

	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%v"}`, err.Error()), http.StatusBadRequest)
		return
	}

	events := c.repo.Calendar.GetEventForWeek(UserID, &date)

	renderJSON(w, EventResponse{Result: events})
}

//GetEventForMonth return list of events on a given month
func (c *CalendarHandler) GetEventForMonth(w http.ResponseWriter, r *http.Request) {
	if !r.URL.Query().Has("user_id") || !r.URL.Query().Has("date") {
		http.Error(w, `{"error": "not enough parameter"}`, http.StatusBadRequest)
		return
	}

	UserID := r.URL.Query().Get("user_id")
	date, err  := time.Parse(time.RFC3339, r.URL.Query().Get("date"))

	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%v"}`, err.Error()), http.StatusBadRequest)
		return
	}

	events := c.repo.Calendar.GetEventForMonth(UserID, &date)

	renderJSON(w, EventResponse{Result: events})
}
