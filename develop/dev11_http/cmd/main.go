package main


import (
	"net/http"
	"http_server/internal/handler"
	"http_server/internal/middleware"
	"log"
	
)

func main() {
	mux := http.NewServeMux()
	server := handlers.NewCalendarHandler()
	mux.HandleFunc("/create_event/", server.CreateEvent)
	mux.HandleFunc("/update_event/", server.UpdateEvent)
	mux.HandleFunc("/delete_event/", server.DeleteEvent)
	mux.HandleFunc("/events_for_day/", server.GetEventForDay)
	mux.HandleFunc("/events_for_week/", server.GetEventForWeek)
	mux.HandleFunc("/events_for_month/", server.GetEventForMonth)

	handler := middleware.Logging(mux)
	log.Fatal(http.ListenAndServe("localhost:8080", handler))
}