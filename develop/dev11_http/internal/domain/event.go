package domain

import (
	"time"
)

//Event represents caledar event to store
type Event struct {
	UserID		string 		`json:"user_id"`
	ID			string		`json:"event_id"`
	Date		time.Time	`json:"date"`	
	Description	string		`json:"description"`
}
