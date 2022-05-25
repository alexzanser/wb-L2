package handlers

import (
	"encoding/json"
	"fmt"
	"http_server/internal/domain"
	"http_server/internal/repository"
	"mime"
	"net/http"
)

type calendarHandler struct {
	repo *repository.Repository
}

func NewCalendarHandler() *calendarHandler {
	return &calendarHandler{
		repo: repository.New(),
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

func (c *calendarHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
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

func (c *calendarHandler) UpdateEvent(w http.ResponseWriter, r *http.Request) {
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

func (c *calendarHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
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

func (c *calendarHandler) GetEventForDay(w http.ResponseWriter, r *http.Request) {
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

func (c *calendarHandler) GetEventForWeek(w http.ResponseWriter, r *http.Request) {
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

func (c *calendarHandler) GetEventForMonth(w http.ResponseWriter, r *http.Request) {
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
