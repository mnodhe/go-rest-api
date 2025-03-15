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
	UserID      int64
}

func (event *Event) Save() error {
	//add it to a db
	query := "INSERT INTO event(name,description,location,dateTime,user_id) VALUES (?,?,?,?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, errr := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID)
	if errr != nil {
		return err
	}
	id, errrr := result.LastInsertId()
	if errrr != nil {
		return err
	}
	event.ID = id
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
func GetEventById(id int64) (Event, error) {
	query := "SELECT * FROM event WHERE id=?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return Event{}, err
	}
	return event, nil
}
func (event Event) Update() error {
	query := "UPDATE event SET name=?, description=?, location=?,dateTime=?,user_id=? WHERE id=?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	exec, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserID, event.ID)
	if err != nil {
		return err
	}
	_, err = exec.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
func (event Event) Delete() error {
	query := "DELETE FROM event WHERE id=?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID)
	if err != nil {
		return err
	}
	return nil
}
func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id,user_id) VALUES (?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(e.ID, userId)
	return err

}
func (e Event) IsRegistrationExist(userId int64) (bool, error) {
	query := "SELECT COUNT(*) FROM registrations WHERE user_id=? AND event_id=?"
	var count int
	err := db.DB.QueryRow(query, userId, e.ID).Scan(&count)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registrations WHERE user_id=? AND event_id=?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(userId, e.ID)
	return err
}
