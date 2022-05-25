package domain

import (
	"time"
)

type Event struct {
	UserID		string 		`json:"user_id"`
	ID			string		`json:"event_id"`
	Date		time.Time	`json:"date"`	
	Description	string		`json:"description"`
}
