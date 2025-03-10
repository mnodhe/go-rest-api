package models

import (
	"go-rest-api/db"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

func (e Event) Save() error {
	//add it to a db
	query := "INSERT INTO event(name,description,location,dateTime,user_id) VALUES (?,?,?,?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, errr := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if errr != nil {
		return err
	}
	id, errrr := result.LastInsertId()
	if errrr != nil {
		return err
	}
	e.ID = id
	return errrr
}
func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM event"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func New(ID int64, name string, description string, location string, dateTime time.Time, userID int) *Event {
	return &Event{ID: ID, Name: name, Description: description, Location: location, DateTime: dateTime, UserID: userID}
}
