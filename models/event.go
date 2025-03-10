package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events []Event

func (e Event) Save() {
	//add it to a db
	events = append(events, e)
}
func GetAllEvents() []Event {
	return events
}

func New(ID int, name string, description string, location string, dateTime time.Time, userID int) *Event {
	return &Event{ID: ID, Name: name, Description: description, Location: location, DateTime: dateTime, UserID: userID}
}
